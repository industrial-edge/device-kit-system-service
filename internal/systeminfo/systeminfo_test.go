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
	systemmocks "systemservice/internal/systeminfo/mocks"
	"testing"

	"github.com/shirou/gopsutil/v3/cpu"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

var TestDeviceDetailsInfo = `MANUFACTURER_NAME="Siemens AG"
SERIAL_NUMBER="1234567890"
PRODUCT_ID="1234567890"
VARIANT="IEMS_0.0.9-dev_x86"
PRODUCT_NAME="IEMS"
BIOS_VERSION="1.0.0"
`

var TestDeviceDetailsWithoutProperInfo = `MANUFACTURER_NAME="Siemens AG"
PRODUCT_ID="1234567890"
VARIANT="IEMS_0.0.9-dev_x86"
BIOS_VERSION="1.0.0"
`

var TestDeviceDetailsWithoutAnyInfo = ``

var TestDeviceDetailsWIthMalformedInfo = `MANUFACTURER_NAME=
SERIAL_NUMBER="1234567890
PRODUCT_ID="1234567890"
VARIANT = 
PRODUCT_NAME="IEMS"
BIOS_VERSION="1.0.0
`

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

var DefaultModelInfo = "NOTDEFINED"

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

func Test_GetDeviceInfo_Details(t *testing.T) {
	//Prepare Content
	t.Parallel()
	tFs, _, tsysInfo := initialize()

	tFs.ReadFileList = make([]mocks.ReadFileContainer, 0)
	r1 := mocks.ReadFileContainer{ReadFileVal: []byte(TestDeviceDetailsInfo), ReadFileErr: nil}
	tFs.ReadFileList = append(tFs.ReadFileList, r1)

	deviceDetails, deviceDetailsError := tsysInfo.GetDeviceInfo()

	assert.NoError(t, deviceDetailsError)
	assert.Equal(t, "Siemens AG", deviceDetails.ManufacturerName)
	assert.Equal(t, "1234567890", deviceDetails.SerialNumber)
	assert.Equal(t, "1234567890", deviceDetails.ProductId)
	assert.Equal(t, "IEMS_0.0.9-dev_x86", deviceDetails.FirmwareVersion)
	assert.Equal(t, "IEMS", deviceDetails.ProductName)
	assert.Equal(t, "1.0.0", deviceDetails.BiosVersion)
}

func Test_GetDeviceInfo_WithoutProperContent(t *testing.T) {
	//Prepare Content
	t.Parallel()
	tFs, _, tsysInfo := initialize()

	tFs.ReadFileList = make([]mocks.ReadFileContainer, 0)
	r1 := mocks.ReadFileContainer{ReadFileVal: []byte(TestDeviceDetailsWithoutProperInfo), ReadFileErr: nil}
	tFs.ReadFileList = append(tFs.ReadFileList, r1)

	deviceDetails, deviceDetailsError := tsysInfo.GetDeviceInfo()

	assert.NoError(t, deviceDetailsError)
	assert.Equal(t, "Siemens AG", deviceDetails.ManufacturerName)
	assert.Empty(t, deviceDetails.SerialNumber)
	assert.Equal(t, "1234567890", deviceDetails.ProductId)
	assert.Equal(t, "IEMS_0.0.9-dev_x86", deviceDetails.FirmwareVersion)
	assert.Empty(t, deviceDetails.ProductName)
	assert.Equal(t, "1.0.0", deviceDetails.BiosVersion)
}

func Test_GetDeviceInfo_MalformedContent(t *testing.T) {
	//Prepare Content
	t.Parallel()
	tFs, _, tsysInfo := initialize()
	tFs.ReadFileList = make([]mocks.ReadFileContainer, 0)
	r1 := mocks.ReadFileContainer{ReadFileVal: []byte(TestDeviceDetailsWIthMalformedInfo), ReadFileErr: nil}
	tFs.ReadFileList = append(tFs.ReadFileList, r1)
	deviceDetails, deviceDetailsError := tsysInfo.GetDeviceInfo()

	assert.NoError(t, deviceDetailsError)
	assert.Empty(t, deviceDetails.ManufacturerName)
	assert.Equal(t, "1234567890", deviceDetails.SerialNumber)
	assert.Equal(t, "1234567890", deviceDetails.ProductId)
	assert.Empty(t, deviceDetails.FirmwareVersion)
	assert.Equal(t, "IEMS", deviceDetails.ProductName)
	assert.Equal(t, "1.0.0", deviceDetails.BiosVersion)
}

func Test_GetDeviceInfo_WithEmptyData(t *testing.T) {
	//Prepare Content
	t.Parallel()
	tFs, _, tsysInfo := initialize()

	tFs.ReadFileList = make([]mocks.ReadFileContainer, 0)
	r1 := mocks.ReadFileContainer{ReadFileVal: []byte(TestDeviceDetailsWithoutAnyInfo), ReadFileErr: nil}
	tFs.ReadFileList = append(tFs.ReadFileList, r1)
	deviceDetails, deviceDetailsError := tsysInfo.GetDeviceInfo()

	assert.NoError(t, deviceDetailsError)
	assert.Empty(t, deviceDetails.ManufacturerName)
	assert.Empty(t, deviceDetails.SerialNumber)
	assert.Empty(t, deviceDetails.ProductId)
	assert.Empty(t, deviceDetails.FirmwareVersion)
	assert.Empty(t, deviceDetails.ProductName)
	assert.Empty(t, deviceDetails.BiosVersion)
}

func Test_GetDeviceInfo_FileReadError(t *testing.T) {
	t.Parallel()
	tFs, _, tsysInfo := initialize()

	tFs.ReadFileList = make([]mocks.ReadFileContainer, 0)
	r1 := mocks.ReadFileContainer{ReadFileVal: nil, ReadFileErr: errors.New("failed to read device info: cannot read file")}
	tFs.ReadFileList = append(tFs.ReadFileList, r1)

	_, err := tsysInfo.GetDeviceInfo()

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to read device info: cannot read file")
}

func Test_GetDeviceInfo_NonUTF8Content(t *testing.T) {
	//Prepare Content
	t.Parallel()
	tFs, _, tsysInfo := initialize()
	tFs.ReadFileList = make([]mocks.ReadFileContainer, 0)
	r1 := mocks.ReadFileContainer{ReadFileVal: []byte{0xff, 0xfe, 0xfd, 0xfc}, ReadFileErr: nil}
	tFs.ReadFileList = append(tFs.ReadFileList, r1)
	deviceDetails, deviceDetailsError := tsysInfo.GetDeviceInfo()

	assert.NoError(t, deviceDetailsError)
	assert.Empty(t, deviceDetails.ManufacturerName)
	assert.Empty(t, deviceDetails.SerialNumber)
	assert.Empty(t, deviceDetails.ProductId)
	assert.Empty(t, deviceDetails.FirmwareVersion)
	assert.Empty(t, deviceDetails.ProductName)
	assert.Empty(t, deviceDetails.BiosVersion)
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

func Test_GetCpuStatsFailure(t *testing.T) {
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
	assert.Equal(t, frequency, float64(-1), "Did not get expected result. Wanted: %q, got: %q", frequency, float64(0))
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

func assertDefaultCPUValues(t *testing.T, cpu *systemapi.Cpu) {
    require.NotNil(t, cpu)
    assert.Equal(t, float32(0), cpu.UsedCpuPercentage)
    assert.Equal(t, float32(0), cpu.FreeCpuPercentage)
    assert.Equal(t, int32(0), cpu.CoreCount)
    assert.Equal(t, DefaultModelInfo, cpu.ModelInfo)
    assert.Equal(t, float64(0), cpu.IdleTime)
    assert.Equal(t, float64(-1), cpu.Frequency)
}

func assertDefaultMemoryValues(t *testing.T, memory *systemapi.Resource) {
    require.NotNil(t, memory)
    assert.Equal(t, float32(0), memory.TotalSpaceInGB)
    assert.Equal(t, float32(0), memory.FreeSpaceInGB)
    assert.Equal(t, float32(0), memory.UsedSpaceInGB)
    assert.Equal(t, float32(0), memory.PercentageFreeSpace)
    assert.Equal(t, float32(0), memory.PercentageUsedSpace)
}

func assertDefaultStorageValues(t *testing.T, storageDevices []*systemapi.Resource) {
    require.NotNil(t, storageDevices)
    assert.Len(t, storageDevices, 0)
}

func Test_GetResourceStats_ErrorAggregationAndDefaults(t *testing.T) {
    t.Parallel()

    // Initialize real SystemInfo with mocks
    tFs, tUtil, sysInfo := initialize()

    // Mock command errors
    tUtil.CommandList = []mocks.CmdContainer{
        {CommandVal: []byte("Max Speed: 2400 MHz\nMax Speed: 2000 MHz"), CommandErr: errors.New("mocked command fail")},
    }

    // Mock FS read failures
    tFs.ReadFileList = []mocks.ReadFileContainer{
        {ReadFileVal: nil, ReadFileErr: errors.New("read fail")},
    }

    // Mock util-level errors
    tUtil.CPUPercErr = errors.New("cpu fail")
    tUtil.CPUCountErr = errors.New("cpu count fail")
    tUtil.CPUInfoErr = errors.New("cpu info fail")
    tUtil.VirtualMemErr = errors.New("mem fail")
    tUtil.UptimeErr = errors.New("uptime fail")

    expectedCPUErr := "CPUPercent: cpu fail"
    expectedCPUCountErr := "CPUCounts: cpu count fail"
    expectedCPUInfoErr := "CPUInfo: cpu info fail"
    expectedMemErr := "Memory: mem fail"
    expectedUptimeErr := "Uptime: uptime fail"

    stats, err := sysInfo.GetResourceStats()

    require.NotNil(t, stats, "Stats result should never be nil")
    require.Error(t, err)

    errStr := err.Error()

    assert.Contains(t, errStr, expectedCPUErr)
    assert.Contains(t, errStr, expectedCPUCountErr)
    assert.Contains(t, errStr, expectedCPUInfoErr)
    assert.Contains(t, errStr, expectedMemErr)
    assert.Contains(t, errStr, expectedUptimeErr)

    assertDefaultCPUValues(t, stats.Cpu)

    assertDefaultMemoryValues(t, stats.Memory)

    assertDefaultStorageValues(t, stats.StorageDevices)

    // Default Uptime
    assert.Equal(t, "0 days, 0 hours, 0 minutes", stats.UpTime)
}

func Test_getCpuFrequency_OutputParsing(t *testing.T) {
    testCases := []struct {
        name           string
        commandOutput  string
        commandError   error
        expectedSpeed  float64
        expectedError  error
    }{
        {
            name:          "Multiple speeds with parse error - returns max valid speed",
            commandOutput: "Max Speed: 2000 MHz\nMax Speed: 3000 MHz\nMax Speed: notanumber MHz\nMax Speed: 2500 MHz",
            commandError:  nil,
            expectedSpeed: 3000,
            expectedError: nil,
        },
        {
            name:          "Single valid speed",
            commandOutput: "Max Speed: 2400 MHz",
            commandError:  nil,
            expectedSpeed: 2400,
            expectedError: nil,
        },
        {
            name:          "Command error - virtualized environment",
            commandOutput: "",
            commandError:  errors.New("command failed"),
            expectedSpeed: -1,
            expectedError: nil,
        },
        {
            name:          "No speed information found",
            commandOutput: "No speed information available",
            commandError:  nil,
            expectedSpeed: -1,
            expectedError: nil,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            t.Parallel()
            _, tUtil, tsysInfo := initialize()

            tUtil.CommandList = []mocks.CmdContainer{
                {CommandVal: []byte(tc.commandOutput), CommandErr: tc.commandError},
            }

            speed, err := tsysInfo.getCpuFrequency()

            if tc.expectedError != nil {
                assert.Error(t, err)
                assert.Contains(t, err.Error(), tc.expectedError.Error())
            } else {
                assert.NoError(t, err)
            }
            assert.Equal(t, tc.expectedSpeed, speed)
        })
    }
}

func Test_GetResourceStats_CpuNil(t *testing.T) {
    t.Parallel()

    tFs, tUtil, sysInfo := initialize()

    // Prevent panic in Commander
    tUtil.CommandList = []mocks.CmdContainer{
        {CommandVal: []byte("Max Speed: 2400 MHz"), CommandErr: nil},
    }

    // Mock CPU to make getCpuStats() return nil internally
    tUtil.CPUPercVal = nil
    tUtil.CPUCountVal = 0
    tUtil.CPUInfoVal = nil
    tUtil.CPUPercErr = nil
    tUtil.CPUCountErr = nil
    tUtil.CPUInfoErr = nil

    // Safe FS mock — ReadFile returns error so storage is skipped
    tFs.ReadFileList = []mocks.ReadFileContainer{
        {ReadFileVal: nil, ReadFileErr: errors.New("read fail")},
    }

    sysInfo.util = tUtil
    sysInfo.fs = tFs

    // Execute the method
    stats, err := sysInfo.GetResourceStats()

    require.NotNil(t, stats)
    require.NoError(t, err)

    // Validate that default CPU struct was created
    require.NotNil(t, stats.Cpu)
    assert.Equal(t, float32(0), stats.Cpu.UsedCpuPercentage)
    assert.Equal(t, "NOTDEFINED", stats.Cpu.ModelInfo)
}

func Test_GetResourceStats_CpuNil_NoError(t *testing.T) {
    t.Parallel()

    tFs, tUtil, _ := initialize()

    // Safe Commander setup
    tUtil.CommandList = []mocks.CmdContainer{
        {CommandVal: []byte("Max Speed: 2400 MHz"), CommandErr: nil},
    }

    // Mock FS — prevents panic in storage
    tFs.ReadFileList = []mocks.ReadFileContainer{
        {ReadFileVal: nil, ReadFileErr: errors.New("read fail")},
    }

    // Replace with a mock that returns nil CPU but no error
    mockSys := &systemmocks.MockSystemInfo{
        MockCPUStatsFunc: func() (*systemapi.Cpu, error) {
            // Return a default CPU object with NOTDEFINED (matching real behavior when CPU info fails)
            return &systemapi.Cpu{
                UsedCpuPercentage: 0,
                FreeCpuPercentage: 0,
                CoreCount:         0,
                ModelInfo:         "NOTDEFINED",
                IdleTime:          0,
                Frequency:         0,
            }, nil
        },
        MockMemoryStatFunc: func() (*systemapi.Resource, error) {
            return &systemapi.Resource{}, nil // normal memory path
        },
        MockStorageStatsFunc: func() []*systemapi.Resource {
            return []*systemapi.Resource{}
        },
        MockUpTimeFunc: func() (string, error) {
            return "0 days, 0 hours, 0 minutes", nil
        },
    }

    stats, err := mockSys.GetResourceStats()
    require.NoError(t, err)
    require.NotNil(t, stats.Cpu)
    assert.Equal(t, "NOTDEFINED", stats.Cpu.ModelInfo)
}