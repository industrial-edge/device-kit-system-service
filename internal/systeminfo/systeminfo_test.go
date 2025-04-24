/*
 * Copyright © Siemens 2020 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package systeminfo

import (
	"errors"
	systemapi "systemservice/api/siemens_iedge_dmapi_v1"
	"systemservice/internal/common/mocks"
	"testing"

	"github.com/shirou/gopsutil/v3/cpu"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/stretchr/testify/assert"
)

var TestContentProperFirmwareInfo = `VARIANT="IEMS_0.0.9-dev_x86"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME=Ubuntu 18.04.4 LTS`

var TestContentNotProperFirmwareInfo = `ID=ubuntu
ID_LIKE=debian
PRETTY_NAME=Ubuntu 18.04.4 LTS`

var TestContentNoNewLineFirmwareInfo = `VARIANT="IEMS_0.0.9-dev_x86"`

var SystemProductName = "VMware Virtual Platform"
var DmidecodeCpuMaxSpeed = "dmidecode -t processor | grep 'Max Speed'"


const testContentProperJSON = `
{
"Limits":
{
	"MaxInstalledApplications":100,
	"MaxRunningApplications":100,
	"MaxMemoryUsageInGB":4.0,
	"MaxStorageUsageInGB":20.0,
	"MaxCpuUsagePerecentage":60
},
"MonitoredStorage":
{
	"Path":"/tmp"
}
}
`

const testContentNonProperJSON = `
{
"Limits":
{
	"MaxInstalledApplications":100,
	"MaxRunningApplications":100,
	"MaxMemoryUsageInGB":4.0,
	"MaxStorageUsageInGB":20.0,
	"MaxCpuUsagePerecentage":60
},
"MonitoredStorage":
{
	"Path":"/dummymount"
`

func initialize() (*mocks.MFS, *mocks.MUtil, *SystemInfo) {
	tFs := new(mocks.MFS)
	tUtil := new(mocks.MUtil)
	tSystemInfo := NewSystemInfo(tFs, tUtil)

	return tFs, tUtil, tSystemInfo
}

func Test_GetFirmwareInfo_WithProperContent(t *testing.T) {
	//Prepare Content
	t.Parallel()
	tFs, _, tsysInfo := initialize()

	tFs.ReadFileList = make([]mocks.ReadFileContainer, 0)
	r1 := mocks.ReadFileContainer{ReadFileVal: []byte(TestContentProperFirmwareInfo), ReadFileErr: nil}
	tFs.ReadFileList = append(tFs.ReadFileList, r1)

	firmwareInfo, err := tsysInfo.GetFirmwareInfo()

	assert.Nil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)
	assert.Equal(t, "IEMS_0.0.9-dev_x86", firmwareInfo.Version, "Did not get expected result. Wanted: %q, got: %q", "IEMS_0.0.9-dev_x86", firmwareInfo.Version)
}

func Test_GetMatchingField_NoNewLineExist(t *testing.T) {
	//Prepare Content
	t.Parallel()
	_, _, tsysInfo := initialize()
	str, err := tsysInfo.getMatchingField(TestContentNoNewLineFirmwareInfo, versionKEY, 0)

	assert.Nil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)
	assert.Equal(t, "IEMS_0.0.9-dev_x86", str, "Did not get expected result. Wanted: %q, got: %q", "IEMS_0.0.9-dev_x86", str)
}

func Test_GetValidModelNumber(t *testing.T) {
	//Prepare Content
	t.Parallel()
	_, tUtil, tsysInfo := initialize()

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte(SystemProductName), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	modelNumber, err := tsysInfo.GetModelNumber()

	assert.Nil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)
	assert.Equal(t, modelNumber.ModelNumber, SystemProductName, "Did not get expected result. Wanted: %q, got: %q", SystemProductName, modelNumber.ModelNumber)
}

func Test_GetUptimeVerify(t *testing.T) {
	//Prepare Content
	t.Parallel()
	_, tUtil, tsysInfo := initialize()

	tUtil.UptimeErr = nil
	tUtil.UptimeVal = 60
	uptime, _ := tsysInfo.getUpTime()
	assert.Equal(t, uptime, "0 days, 0 hours, 1 minutes", "Did not get expected result. Wanted: %q, got: %q", "0 days, 0 hours, 1 minutes", uptime)

	tUtil.UptimeVal = 100000
	uptime, _ = tsysInfo.getUpTime()
	assert.Equal(t, uptime, "1 days, 3 hours, 46 minutes", "Did not get expected result. Wanted: %q, got: %q", "1 days, 3 hours, 46 minutes", uptime)

	tUtil.UptimeVal = 18446744073709551615
	uptime, _ = tsysInfo.getUpTime()
	assert.Equal(t, uptime, "213503982334601 days, 7 hours, 0 minutes", "Did not get expected result. Wanted: %q, got: %q", "213503982334601 days, 7 hours, 0 minutes", uptime)

}

func Test_GetUptimeFailure(t *testing.T) {
	//Prepare Content
	t.Parallel()
	_, tUtil, tsysInfo := initialize()

	tUtil.UptimeErr = errors.New("Failed to GetUpTime")
	uptime, err := tsysInfo.getUpTime()

	assert.Equal(t, uptime, "0 days, 0 hours, 0 minutes", "Did not get expected result. Wanted: %q, got: %q", "0 days, 0 hours, 0 minutes", uptime)
	assert.NotNil(t, err, "Did not get expected result. got: %q", err)
}

func Test_GetMemoryStatFailure(t *testing.T) {
	//Prepare Content
	t.Parallel()
	_, tUtil, tsysInfo := initialize()
	tUtil.VirtualMemErr = errors.New("Failed to GetMemoryStat")

	_, err := tsysInfo.getMemoryStat()

	assert.NotNil(t, err, "Did not get expected result.  got: %q", err)
}

func Test_GetMemoryStatVerifyWithProperStats(t *testing.T) {
	//Prepare Content
	t.Parallel()
	_, tUtil, tsysInfo := initialize()
	tUtil.VirtualMemErr = nil
	tUtil.VirtualMemStat = mem.VirtualMemoryStat{
		Total:       100000000,
		Available:   10000000,
		Used:        80000000,
		Free:        10000000,
		UsedPercent: 80,
	}

	memoryStat, err := tsysInfo.getMemoryStat()
	assert.Nil(t, err, "Did not get expected result.  got: %q", err)
	assert.Equal(t, memoryStat.TotalSpaceInGB, float32(0.09313226), "Did not get expected result. Wanted: %q, got: %q", float32(0.09313226), memoryStat.TotalSpaceInGB)
	assert.Equal(t, memoryStat.FreeSpaceInGB, float32(0.009313226), "Did not get expected result. Wanted: %q, got: %q", float32(0.009313226), memoryStat.TotalSpaceInGB)
	assert.Equal(t, memoryStat.UsedSpaceInGB, float32(0.074505806), "Did not get expected result. Wanted: %q, got: %q", float32(0.074505806), memoryStat.TotalSpaceInGB)
	assert.Equal(t, memoryStat.PercentageUsedSpace, float32(80), "Did not get expected result. Wanted: %q, got: %q", float32(80), memoryStat.TotalSpaceInGB)
	assert.Equal(t, memoryStat.PercentageFreeSpace, float32(10), "Did not get expected result. Wanted: %q, got: %q", float32(10), memoryStat.TotalSpaceInGB)
}

func Test_GetCpuStatsFailures(t *testing.T) {
	//Prepare Content
	t.Parallel()
	_, tUtil, tsysInfo := initialize()
	tUtil.CPUPercErr = errors.New("Failed to get CpuPercentage")
	tUtil.CPUCountErr = errors.New("Failed to get CpuCount")
	tUtil.CPUInfoErr = errors.New("Failed to get CpuInfo")
	tUtil.IdleTimeErr = errors.New("Failed to get CpuIdleTime")
	tUtil.FrequencyErr = errors.New("Failed to get CpuFrequency")

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte("Max Speed: 30000 MHz"), CommandErr: tUtil.FrequencyErr}
	tUtil.CommandList = append(tUtil.CommandList, s1)
	
	_, err := tsysInfo.getCPUStats()
	assert.NotNil(t, err, "Did not get expected result.  got: %q", err)
}

func Test_GetCpuStatsFailure(t *testing.T){
	t.Parallel()
	_, tUtil, tsysInfo := initialize()
	
	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte(SystemProductName), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	//Set only one Failure
	tUtil.CPUPercErr = errors.New("Failed to get CpuPercentage")
	_, err := tsysInfo.getCPUStats()
	assert.NotNil(t, err, "Did not get expected result.  got: %q", err)
}

func Test_GetCpuStatsSuccess(t *testing.T) {
	//Prepare Content
	t.Parallel()
	_, tUtil, tsysInfo := initialize()
	res, _ := tUtil.CPUInfo()
	newRes := append(res, cpu.InfoStat{ModelName: "test"})
	tUtil.CPUInfoVal = newRes

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte(SystemProductName), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	_, err := tsysInfo.getCPUStats()
	assert.Nil(t, err, "Did not get expected result.  got: %q", err)
}

func Test_getCpuFrequency(t *testing.T) {
	//Prepare Content
	t.Parallel()
	_, tUtil, tsysInfo := initialize()

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte(DmidecodeCpuMaxSpeed), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	frequency, err := tsysInfo.getCpuFrequency()

	assert.Nil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)
	assert.Equal(t, frequency, float64(0), "Did not get expected result. Wanted: %q, got: %q", frequency, float64(0))
}

func Test_getCpuIdleTime(t *testing.T) {
	//Prepare Content
	t.Parallel()
	_, tUtil, tsysInfo := initialize()
	idle := 123.99
	res, _ := tUtil.CPUIdleTime()
	newRes := append(res, cpu.TimesStat{Idle: idle})
	tUtil.IdleTime = newRes

	result, err := tsysInfo.getCpuIdleTime()
	assert.Equal(t, idle, result)
	assert.Nil(t, err, "Did not get expected result.  got: %q", err)
}

func Test_GetStorageStats_Fails_Due_To_Read_Error(t *testing.T) {
	//Prepare Content
	t.Parallel()
	tFs, _, tsysInfo := initialize()

	tFs.ReadFileList = make([]mocks.ReadFileContainer, 0)
	r1 := mocks.ReadFileContainer{ReadFileVal: []byte(testContentProperJSON), ReadFileErr: errors.New("Failed to Read")}
	tFs.ReadFileList = append(tFs.ReadFileList, r1)

	storageStats := tsysInfo.getStorageStats()
	t.Log("StorageStats: ", storageStats)

	assert.Equal(t, storageStats, []*systemapi.Resource{}, "Content: %q Expected: %q", storageStats, []*systemapi.Resource{})
}

func Test_GetStorageStats_Fails_Due_To_Unmarshall_Error(t *testing.T) {
	//Prepare Content
	t.Parallel()
	tFs, _, tsysInfo := initialize()

	tFs.ReadFileList = make([]mocks.ReadFileContainer, 0)
	r1 := mocks.ReadFileContainer{ReadFileVal: []byte(testContentNonProperJSON), ReadFileErr: nil}
	tFs.ReadFileList = append(tFs.ReadFileList, r1)

	storageStats := tsysInfo.getStorageStats()
	t.Log("StorageStats: ", storageStats)

	assert.Equal(t, storageStats, []*systemapi.Resource{}, "Content: %q Expected: %q", storageStats, []*systemapi.Resource{})
}

func Test_GetStorageStats_Success(t *testing.T) {
	//Prepare Content
	t.Parallel()
	tFs, _, tsysInfo := initialize()

	tFs.ReadFileList = make([]mocks.ReadFileContainer, 0)
	r1 := mocks.ReadFileContainer{ReadFileVal: []byte(testContentProperJSON), ReadFileErr: nil}
	tFs.ReadFileList = append(tFs.ReadFileList, r1)

	storageStats := tsysInfo.getStorageStats()
	t.Log("StorageStats: ", storageStats)

	assert.NotNil(t, storageStats[0].FreeSpaceInGB, "Did not get expected result for FreeSpaceInGB.  got: %q", storageStats[0].FreeSpaceInGB)
	assert.NotNil(t, storageStats[0].TotalSpaceInGB, "Did not get expected result for TotalSpaceInGB.  got: %q", storageStats[0].TotalSpaceInGB)
	assert.NotNil(t, storageStats[0].UsedSpaceInGB, "Did not get expected result for UsedSpaceInGB.  got: %q", storageStats[0].UsedSpaceInGB)
}

func Test_SystemInfo_GetLogFile_HappyPath(t *testing.T) {
	t.Parallel()
	_, tUtil, tsysInfo := initialize()

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte("[-p /logpath ] check directory"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	s2 := mocks.CmdContainer{CommandVal: []byte("journalctl > logs"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s2)

	s3 := mocks.CmdContainer{CommandVal: []byte("[-f /var/device.name ]"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s3)

	s4 := mocks.CmdContainer{CommandVal: []byte("hostname"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s4)

	s5 := mocks.CmdContainer{CommandVal: []byte("tar -czvf log ..."), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s5)

	_, err := tsysInfo.GetLogFile(&systemapi.LogRequest{SaveFolderPath: "/tmp/tmp"})
	if err != nil {
		t.Log(err.Error())
	}
	assert.NoError(t, err)
}

func Test_SystemInfo_GetLogFile_ErrorOnZip(t *testing.T) {
	t.Parallel()
	_, tUtil, tsysInfo := initialize()

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte("[-p /logpath ] check directory"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	s2 := mocks.CmdContainer{CommandVal: []byte("journalctl > logs"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s2)

	s3 := mocks.CmdContainer{CommandVal: []byte("[-f /var/device.name ]"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s3)

	s4 := mocks.CmdContainer{CommandVal: []byte("hostname"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s4)

	s5 := mocks.CmdContainer{CommandVal: []byte("tar -czvf log ..."), CommandErr: errors.New("An error was encountered during the compression process.")}
	tUtil.CommandList = append(tUtil.CommandList, s5)

	_, err := tsysInfo.GetLogFile(&systemapi.LogRequest{SaveFolderPath: "/tmp/tmp"})
	if err != nil {
		t.Log(err.Error())
	}
	assert.Contains(t, err.Error(), "An error was encountered during the compression process.")
}

func Test_SystemInfo_GetLogFile_ErrorOnCatDeviceName(t *testing.T) {
	t.Parallel()
	_, tUtil, tsysInfo := initialize()

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte("[-p /logpath ] check directory"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	s2 := mocks.CmdContainer{CommandVal: []byte("journalctl > logs"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s2)

	s3 := mocks.CmdContainer{CommandVal: []byte("[-f /var/device.name ]"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s3)

	s4 := mocks.CmdContainer{CommandVal: []byte("cat device.name"), CommandErr: errors.New("error open")}
	tUtil.CommandList = append(tUtil.CommandList, s4)

	s5 := mocks.CmdContainer{CommandVal: []byte("tar -czvf log ..."), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s5)

	_, err := tsysInfo.GetLogFile(&systemapi.LogRequest{SaveFolderPath: "/tmp/tmp"})
	if err != nil {
		t.Log(err.Error())
	}
	assert.NoError(t, err)
}

func Test_SystemInfo_GetLogFile_ErrorOnPathCheck(t *testing.T) {
	t.Parallel()
	_, tUtil, tsysInfo := initialize()

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte("[-p /logpath ] check directory"), CommandErr: errors.New("The directory does not exist in the system.")}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	_, err := tsysInfo.GetLogFile(&systemapi.LogRequest{SaveFolderPath: "/tmp/tmp"})
	if err != nil {
		t.Log(err.Error())
	}
	assert.Contains(t, err.Error(), "The directory does not exist in the system.")
}

func Test_SystemInfo_GetLogFile_ErrorOnJournalctl(t *testing.T) {
	t.Parallel()
	_, tUtil, tsysInfo := initialize()

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte("[-p /logpath ] check directory"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	s2 := mocks.CmdContainer{CommandVal: []byte("journalctl > logs"), CommandErr: errors.New("Something went wrong to execute journalctl.")}
	tUtil.CommandList = append(tUtil.CommandList, s2)

	_, err := tsysInfo.GetLogFile(&systemapi.LogRequest{SaveFolderPath: "/tmp/tmp"})
	if err != nil {
		t.Log(err.Error())
	}
	assert.Contains(t, err.Error(), "Something went wrong to execute journalctl.")
}

func Test_SystemInfo_GetLogFile_ErrorOnDeviceNameFileCheck(t *testing.T) {
	t.Parallel()
	_, tUtil, tsysInfo := initialize()

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte("[-p /logpath ] check directory"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	s2 := mocks.CmdContainer{CommandVal: []byte("journalctl > logs"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s2)

	s3 := mocks.CmdContainer{CommandVal: []byte("[-f /var/device.name ]"), CommandErr: errors.New("file doesn't exist")}
	tUtil.CommandList = append(tUtil.CommandList, s3)

	s4 := mocks.CmdContainer{CommandVal: []byte("hostname"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s4)

	s5 := mocks.CmdContainer{CommandVal: []byte("tar -czvf log ..."), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s5)

	_, err := tsysInfo.GetLogFile(&systemapi.LogRequest{SaveFolderPath: "/tmp/tmp"})
	if err != nil {
		t.Log(err.Error())
	}
	assert.NoError(t, err)
}

func Test_SystemInfo_GetLogFile_ErrorOnDeviceNameFileCheckAndHostname(t *testing.T) {
	t.Parallel()
	_, tUtil, tsysInfo := initialize()

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte("[-p /logpath ] check directory"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	s2 := mocks.CmdContainer{CommandVal: []byte("journalctl > logs"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s2)

	s3 := mocks.CmdContainer{CommandVal: []byte("[-f /var/device.name ]"), CommandErr: errors.New("file doesn't exist")}
	tUtil.CommandList = append(tUtil.CommandList, s3)

	s4 := mocks.CmdContainer{CommandVal: []byte("hostname"), CommandErr: errors.New("hostname error")}
	tUtil.CommandList = append(tUtil.CommandList, s4)

	s5 := mocks.CmdContainer{CommandVal: []byte("tar -czvf log ..."), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s5)

	_, err := tsysInfo.GetLogFile(&systemapi.LogRequest{SaveFolderPath: "/tmp/tmp"})
	if err != nil {
		t.Log(err.Error())
	}
	assert.NoError(t, err)
}
