/*
 * Copyright (c) 2022 Siemens AG
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package systeminfo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	systemapi "systemservice/api/siemens_iedge_dmapi_v1"
	"systemservice/internal/common/utils"
	"systemservice/internal/limitprovider"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

const (
	firmwareFile = "/etc/os-release"
	versionKEY   = "VARIANT="
	//in order to get a proper cpu usage at least we need to run this command twice (!NOT IN USE Right Now)
	cpuUsage         = "top -b -n2 | grep Cpu"
	dmidecodeVersion = "dmidecode --string system-product-name"
	templogfileDir   = "/tmp"
)

// SystemInfo is a struct that provides System Info
type SystemInfo struct {
	fs   utils.FileSystem
	util utils.Utils
	cfg  *utils.DefaultConfig
}

// NewSystemInfo to create a new instance of SystemInfo
func NewSystemInfo(fsVal utils.FileSystem, utVal utils.Utils) *SystemInfo {
	var systeminfo = SystemInfo{fs: fsVal, util: utVal}
	return &systeminfo
}

// GetResourceStats method provides Cpu, Storage, Memory and Uptime Stats
func (s SystemInfo) GetResourceStats() (*systemapi.Stats, error) {
	var err error = nil

	cpu, err1 := s.getCPUStats()
	storage := s.getStorageStats()
	memory, err2 := s.getMemoryStat()
	upTime, err3 := s.getUpTime()

	if err1 != nil || err2 != nil || err3 != nil {
		err = errors.New("systeminfo:getResourcestats(), failure on a get call" + err1.Error() + err2.Error() + err3.Error())
	}

	return &systemapi.Stats{Cpu: cpu,
		StorageDevices: storage,
		Memory:         memory,
		UpTime:         upTime}, err
}

func (s SystemInfo) getFileContent(path string) (string, error) {
	content, err := s.fs.ReadFile(path)
	if err != nil {
		log.Println("systeminfo:getFileContent(), Cannot Read File! err:", err)
		return "", errors.New("cannot read file")
	}
	return string(content), nil
}

func (s SystemInfo) getMatchingField(content string, key string, space uint16) (string, error) {
	//Search for key value
	if index := strings.Index(content, key); index >= 0 {
		log.Printf("systeminfo:getMatchingField(), key value:%q", key)
		var val string
		if indexEnd := strings.Index(string(content[index:len(content)]), string('\n')); indexEnd >= 0 {
			val = string(content[(index + len(key) + int(space)):(index + indexEnd)])
			val = strings.Replace(val, "\"", "", -1)
			log.Println("systeminfo:getMatchingField(), value:", val)
		} else {
			//else means there is no new line after key value..Therefore we need read up to end of string
			val = string(content[(index + len(key) + int(space)):])
			val = strings.Replace(val, "\"", "", -1)
			log.Println("systeminfo:getMatchingField(), -no new line- value: ", val)
		}
		return val, nil
	}

	return content, errors.New("key value cannot be found")
}

// GetFirmwareInfo provides firmwave Version Info
func (s SystemInfo) GetFirmwareInfo() (*systemapi.FirmwareInfo, error) {
	//Read Firmware Info from file
	content, err1 := s.getFileContent(firmwareFile)
	if err1 != nil {
		return &systemapi.FirmwareInfo{Version: "0"}, err1
	}

	value, err2 := s.getMatchingField(content, versionKEY, 0)
	if err2 != nil {
		return &systemapi.FirmwareInfo{Version: "0"}, err2
	}

	return &systemapi.FirmwareInfo{Version: value}, nil
}

// GetLimits method provides NFR Limit values via limitprovider
func (s SystemInfo) GetLimits() (*systemapi.Limits, error) {
	limitInstance := limitprovider.CreateLimitProvider(s.fs, s.util)
	return limitInstance.GetLimitContent()
}

// GetModelNumber provides device model number
func (s SystemInfo) GetModelNumber() (*systemapi.ModelNumber, error) {
	var value string

	//Read MLFB number from SMBIOS
	command := dmidecodeVersion

	out, err := s.util.Commander(command)
	if err != nil {
		log.Println("systeminfo:getModelNumber(), cannot read ModelNumber err:", err)
	} else {
		value = string(out)
	}

	return &systemapi.ModelNumber{ModelNumber: value}, err
}

func (s SystemInfo) getUpTime() (string, error) {
	//Get UpTime value
	uptime, err := s.util.Uptime()
	if err != nil {
		log.Println("systeminfo:getUpTime(), Cannot Read host.uptime! err: ", err)
		return "0 days, 0 hours, 0 minutes", err
	}

	//host.Uptime Returns the uptime as Seconds
	days := uptime / (60 * 60 * 24)
	hours := (uptime - (days * 60 * 60 * 24)) / (60 * 60)
	minutes := ((uptime - (days * 60 * 60 * 24)) - (hours * 60 * 60)) / 60

	log.Printf("systeminfo:getUpTime(), %d days, %d hours, %d minutes", days, hours, minutes)

	return (strconv.FormatUint(days, 10) + " days, " + strconv.FormatUint(hours, 10) + " hours, " + strconv.FormatUint(minutes, 10) + " minutes"), nil
}

func convertToGb(val uint64) float32 {
	return float32((float64(val) / (1024 * 1024 * 1024)))
}

func (s SystemInfo) getMemoryStat() (*systemapi.Resource, error) {

	var totalSpaceInGb float32
	var freeSpaceInGb float32
	var usedSpaceInGb float32
	var percentageFreeSpace float32
	var percentageUsedSpace float32

	//Get Memory Stats
	retval, err := s.util.VirtualMemory()
	if err != nil {
		log.Println("systeminfo:getMemoryStat(), Cannot Read memSwapMemory! err: ", err)
	} else {
		log.Println("systeminfo:getMemoryStat(), Memory:", retval)
		totalSpaceInGb = convertToGb(retval.Total)
		freeSpaceInGb = convertToGb(retval.Available)
		usedSpaceInGb = convertToGb(retval.Used)
		percentageUsedSpace = float32(retval.UsedPercent)
		percentageFreeSpace = float32(float64(retval.Available*100) / float64(retval.Total))
	}

	return &systemapi.Resource{
		TotalSpaceInGB:      totalSpaceInGb,
		FreeSpaceInGB:       freeSpaceInGb,
		UsedSpaceInGB:       usedSpaceInGb,
		PercentageFreeSpace: percentageFreeSpace,
		PercentageUsedSpace: percentageUsedSpace,
	}, err
}

func (s SystemInfo) getStorageStats() []*systemapi.Resource {

	// Load /opt/limits/default.json file into memory
	content, err := s.getFileContent(utils.DefaultConfigPath)
	if err != nil {
		return []*systemapi.Resource{}
	}

	log.Println("systeminfo:getStorageStats(), jsonContent", content)

	err = json.Unmarshal([]byte(content), &s.cfg)
	if err != nil {
		log.Println("systeminfo:getStorageStats(), Unmarshal() Fail, err:", err)
		return []*systemapi.Resource{}
	}

	//log.Println("systeminfo:getStorageStats(), path to be monitored: ", s.cfg))
	//log.Printf("systeminfo:getStorageStats(), path to be monitored: %s", *s.cfg)
	log.Println("systeminfo:getStorageStats(), path to be monitored: ", s.cfg.MonitoredStorage.Path)

	diskSt := utils.DiskUsage(s.cfg.MonitoredStorage.Path)
	totalSpaceInGb := float32(float64(diskSt.All) / float64(utils.GB))
	freeSpaceInGb := float32(float64(diskSt.Avail) / float64(utils.GB))
	usedSpaceInGb := float32(float64(diskSt.Used) / float64(utils.GB))
	percentageFreeSpace := (freeSpaceInGb * 100) / totalSpaceInGb
	percentageUsedSpace := (usedSpaceInGb * 100) / totalSpaceInGb

	return []*systemapi.Resource{{
		TotalSpaceInGB:      totalSpaceInGb,
		FreeSpaceInGB:       freeSpaceInGb,
		UsedSpaceInGB:       usedSpaceInGb,
		PercentageFreeSpace: percentageFreeSpace,
		PercentageUsedSpace: percentageUsedSpace,
	}}
}

func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

/*
//****** DO NOT DELETE, THIS FUNCTION CAN BE USED IF cpu.Percent is not stable

func (s SystemInfo) findCpuUsage(content string) (float32, float32) {
	//Get Cpu stats
	if indexCall1 := strings.Index(content, "Cpu"); indexCall1 >= 0 {
		//Content belongs to top command
		//First result is not what we are looking for
		content = content[indexCall1+len("Cpu"):]

		if indexCall2 := strings.Index(content, "Cpu"); indexCall2 >= 0 {
			content = content[indexCall2+len("Cpu"):]
			//Search for idle time
			if indexId := strings.Index(content, "id"); indexId >= 0 {

				//Get from idle time
				strFreeCpu := content[indexId-6 : indexId]
				strFreeCpu = strings.TrimSpace(strFreeCpu)
				strFreeCpu = strings.ReplaceAll(strFreeCpu, ",", ".")

				//Calculate the usage value
				freeCpu, err := strconv.ParseFloat(strFreeCpu, 32)
				if err != nil {
					log.Println("findCpuUsage(): Error Occured during Parse! err:", err)
					return 0, 0
				}
				usedCpu := 100 - float64(freeCpu)
				return float32(round(usedCpu, 0.05)), float32(round(freeCpu, 0.05))
			}
		}
	}

	log.Println("findCpuUsage(): Cannot find cpuUsage!")
	return 0, 0
}
*/

func (s SystemInfo) getCPUStats() (*systemapi.Cpu, error) {

	var usedCPUPercentage float32
	var freeCPUPercentage float32
	var modelInfo string
	var coreCount int32

	var errFlag error = nil //errFlag will return only the latest error info
	//Get cpu usage
	if retval, err := s.util.CPUPercent((time.Second * 1), false); err != nil {
		log.Println("systeminfo:getCpuStats(), Cannot Read cpu.Percent! err:", err)
		errFlag = err
	} else {
		if len(retval) > 0 {
			usedCPUPercentage = float32(round(retval[0], 0.05))
			freeCPUPercentage = float32(round((100 - float64(usedCPUPercentage)), 0.05))
		} else {
			log.Println("systeminfo:getCpuStats(), No percentage value!")
		}
	}

	log.Printf("systeminfo:getCpuStats(), Used Cpu: %6.2f, Free Cpu: %6.2f", usedCPUPercentage, freeCPUPercentage)

	//Get Cpu Core Count
	if retval, err := s.util.CPUCounts(false); err != nil {
		log.Println("systeminfo:getCpuStats(), Cannot Read cpu.Counts!, err: ", err)
		errFlag = err
	} else {
		coreCount = int32(retval)
	}

	log.Println("systeminfo:getCpuStats(), Core Count: ", coreCount)

	//Get Model info
	if retval, err := s.util.CPUInfo(); err != nil {
		log.Println("systeminfo:getCpuStats(), Cannot Read cpu.Info!, err: ", err)
		modelInfo = "NOTDEFINED"
		errFlag = err
	} else {
		if len(retval) > 0 {
			modelInfo = s.getModelInfo(retval, modelInfo)
		} else {
			log.Println("systeminfo:getCpuStats(), No Info value!")
			modelInfo = "NOTDEFINED"
		}
	}

	log.Println("systeminfo:getCpuStats(), Model Info: ", modelInfo)

	return &systemapi.Cpu{UsedCpuPercentage: usedCPUPercentage,
		FreeCpuPercentage: freeCPUPercentage,
		CoreCount:         int32(coreCount),
		ModelInfo:         modelInfo}, errFlag

}

// getModelInfo, Added for arm64. The gopsutil library sets this value for the latest cpu core in /proc/cpuinfo on arm devices.
func (s SystemInfo) getModelInfo(retval []cpu.InfoStat, modelInfo string) string {
	for i := len(retval) - 1; i >= 0; i-- {
		if retval[i].ModelName != "" {
			log.Printf("Model information found for the %d. CPU core.", i+1)
			modelInfo = retval[i].ModelName
			break
		}
	}
	return modelInfo
}

// GetLogFile method returns the path of compressed log file with LogResponse struct
func (s SystemInfo) GetLogFile(request *systemapi.LogRequest) (*systemapi.LogResponse, error) {

	// Check whether given path exist or not if it doesn't exist then return.
	pathCheckCommand := "[ -d " + strings.TrimSuffix(request.SaveFolderPath, "/") + " ]"
	_, errCheck := s.util.Commander(pathCheckCommand)
	if errCheck != nil {
		log.Printf("Error %s directory doesn't exist", request.SaveFolderPath)
		return nil, errCheck
	}

	templogfileName := "logs" + "_" + time.Now().Format("20060102150405")
	templogfilePath := templogfileDir + "/" + templogfileName

	// Get logs with the help of "journalctl" command
	saveJournalCommand := "journalctl > " + templogfilePath
	out, err := s.util.Commander(saveJournalCommand)
	if err != nil {
		log.Printf("Error  %s output : %s, error : %s",
			saveJournalCommand,
			strings.Trim(strings.TrimSuffix(string(out), "\n"), "\""), err.Error())
		return nil, err
	}

	// Check,  if device.name file created  get deviceName, otherwise get device hostname
	uniqIdentifierForLogFile := ""
	deviceNameFile := "/var/device.name"
	deviceNameCommand := "[ -f " + deviceNameFile + " ]"

	_, errCheckDevName := s.util.Commander(deviceNameCommand)
	if errCheckDevName != nil {
		log.Println("Error  directory doesn't exist, device is not onboarded...")
		// Device is not onboarded yet, get hostname from system.
		hostnameCommand := "hostname"
		outHost, errHost := s.util.Commander(hostnameCommand)
		if errHost != nil {
			log.Println("Hostname command run error, ", errHost.Error())
		} else {
			uniqIdentifierForLogFile = strings.TrimSuffix(string(outHost), "\n") + "_"
		}

	} else {
		deviceNameCommand := "cat " + deviceNameFile
		outDev, errDev := s.util.Commander(deviceNameCommand)
		if errDev != nil {
			log.Printf("Error while reading file : %s Error : %s \n", outDev, errDev.Error())
		} else {
			log.Println(" device.name file content :", string(outDev))
			uniqIdentifierForLogFile = strings.TrimSpace(string(outDev)) + "_"
		}
	}

	// If device name has `/`, it is not supported for file names, it will be replaced with `?`
	uniqIdentifierForLogFile = strings.ReplaceAll(uniqIdentifierForLogFile, "/", "?")
	// Single quotes also needs to be escaped since file name in zip creation command will be in between single quotes.
	uniqIdentifierForLogFile = strings.ReplaceAll(uniqIdentifierForLogFile, "'", `\'`)

	// create timestamp as "YYYYMMDDhhmmss" format
	logIdentifier := strings.ReplaceAll(uniqIdentifierForLogFile, " ", "") + time.Now().Format("20060102150405")
	logFileName := fmt.Sprintf("devicelogs_%s.tar.gz", logIdentifier)
	log.Println("logName  :", logFileName)

	logFilePath := strings.TrimSuffix(request.SaveFolderPath, "/") + "/" + logFileName
	//compress log file and create zip with the help of "tar"  command. Rename the file to "logs" in the archive.
	zipCommand := fmt.Sprintf("tar -czvf $'%s' --transform='flags=r;s|%s|logs|' -C %s %s --remove-files",
		logFilePath,
		templogfileName,
		templogfileDir,
		templogfileName)

	outZip, errZip := s.util.Commander(zipCommand)

	if errZip != nil {
		log.Printf("Error %s output : %s, error : %s",
			zipCommand,
			strings.Trim(strings.TrimSuffix(string(outZip), "\n"), "\""), errZip.Error())
		return nil, errZip
	}

	return &systemapi.LogResponse{LogPath: logFilePath}, nil
}
