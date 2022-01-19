/*
 * Copyright (c) 2021 Siemens AG
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package mocks

import (
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

//MUtil represents Mocked utils
type MUtil struct {
	CommandCnt     int
	CommandList    []CmdContainer
	UptimeVal      uint64
	UptitmeErr     error
	VirtualmemStat mem.VirtualMemoryStat
	VirtualMemErr  error
	CPUPercErr     error
	CPUCountErr    error
	CPUInfoErr     error
	CPUPercVal     []float64
	CPUCountVal    int
	CPUInfoVal     []cpu.InfoStat
}

//CmdContainer holds the Command and error list for the mock command method
type CmdContainer struct {
	CommandErr error
	CommandVal []byte
}

//Commander method is mock of exec.Command
func (util *MUtil) Commander(command string) ([]byte, error) {
	util.CommandCnt++
	return util.CommandList[util.CommandCnt-1].CommandVal, util.CommandList[util.CommandCnt-1].CommandErr
}

//Uptime method is mock of host.Uptime
func (util *MUtil) Uptime() (uint64, error) {
	return util.UptimeVal, util.UptitmeErr
}

//VirtualMemory method is mock of mem.VirtualMemory
func (util *MUtil) VirtualMemory() (*mem.VirtualMemoryStat, error) {
	return &util.VirtualmemStat, util.VirtualMemErr
}

//CPUPercent method is mock of cpu.Percent
func (util *MUtil) CPUPercent(interval time.Duration, percpu bool) ([]float64, error) {
	return util.CPUPercVal, util.CPUPercErr
}

//CPUCounts method is mock of cpu.Counts
func (util *MUtil) CPUCounts(logical bool) (int, error) {
	return util.CPUCountVal, util.CPUCountErr
}

//CPUInfo method is mock of cpu.CpuInfo
func (util *MUtil) CPUInfo() ([]cpu.InfoStat, error) {
	return util.CPUInfoVal, util.CPUInfoErr
}

//MFS represents Mocked File System
type MFS struct {
	ReadFileCnt    int
	ReadFileList   []ReadFileContainer
	OpenFileErr    error
	OpenFileHandle *os.File
	WrtieFileErr   error
	StatVal        os.FileInfo
	StatErr        error
}

//ReadFileContainer holds the file content and error list for the mock readfile method
type ReadFileContainer struct {
	ReadFileErr error
	ReadFileVal []byte
}

//OpenFile method is mock of os.OpenFile
func (fs *MFS) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return fs.OpenFileHandle, fs.OpenFileErr
}

//ReadFile method is mock of ioutil.ReadFile
func (fs *MFS) ReadFile(filename string) ([]byte, error) {
	fs.ReadFileCnt++
	return fs.ReadFileList[fs.ReadFileCnt-1].ReadFileVal, fs.ReadFileList[fs.ReadFileCnt-1].ReadFileErr
}

//WriteFile method is mock of ioutil.WriteFile
func (fs *MFS) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return fs.WrtieFileErr
}

//Stat method is mock of os.Stat
func (fs *MFS) Stat(name string) (os.FileInfo, error) {
	return fs.StatVal, fs.StatErr
}
