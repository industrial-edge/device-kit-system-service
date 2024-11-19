/*
 * Copyright (c) 2021 Siemens AG
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package utils

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/shirou/gopsutil/v3/disk"
	syscall "golang.org/x/sys/unix"
)

const (
	// B is Byte
	B = 1
	// KB is KiloByte
	KB = 1024 * B
	// MB is MegaByte
	MB = 1024 * KB
	// GB is GigaByte
	GB = 1024 * MB
	// storageDiskTypeCommand is cmd to get block device information
	storageDiskTypeCommand = "lsblk -Jd -o 'NAME,TYPE,ROTA,RM'"
	// blockType is type of block device
	blockType = "disk"
	// ssdDisk is ssd disk type
	ssdDisk = "SSD"
	// hddDisk is hdd disk type
	hddDisk = "HDD"
)

// DiskStatus struct for Status Infos
type DiskStatus struct {
	All                   uint64 `json:"all"`
	Used                  uint64 `json:"used"`
	Free                  uint64 `json:"free"`
	Avail                 uint64 `json:"avail"`
	DiskType              string `json:"disktype"`
	DiskTotalReadSectors  uint64 `json:"disktotalreadsectors"`
	DiskTotalWriteSectors uint64 `json:"disktotalwritesectors"`
}

type BlockDevice struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Rotational bool   `json:"rota"`
	Removable  bool   `json:"rm"`
}

type StorageDiskTypeOutput struct {
	BlockDevices []BlockDevice `json:"blockdevices"`
}

// DiskUsage of path/to/disk
func DiskUsage(path string) (disk DiskStatus, err error) {
	fs := syscall.Statfs_t{}
	err = syscall.Statfs(path, &fs)
	if err != nil {
		return DiskStatus{}, err
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Avail = fs.Bavail * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	diskType, diskName, err := getDiskInfo()
	if err != nil {
		log.Printf("Utils:DiskUsage(), Failed to get disk info: %v", err)
		disk.DiskType = ""
		disk.DiskTotalReadSectors = 0
		disk.DiskTotalWriteSectors = 0
		return
	}
	disk.DiskType = diskType
	totalReadSectors, totalWriteSectors, err := getDiskSpeed(diskName)
	if err != nil {
		log.Printf("Utils:DiskUsage(), Failed to get disk speed: %v", err)
	}

	disk.DiskTotalReadSectors = totalReadSectors
	disk.DiskTotalWriteSectors = totalWriteSectors
	return
}

// getDiskInfo() return the type of the disk as ssd/hdd
func getDiskInfo() (string, string, error) {
	var diskType, diskName string

	// It is possible that it might return an empty string
	// If it not supported by the OS and the scenario is already handled from edge-core
	output, err := exec.Command(shell, "-c", storageDiskTypeCommand).Output()
	if err != nil {
		log.Printf("Utils:getDiskType(), Failed to execute %v: %v", storageDiskTypeCommand, err)
		return "", "", err
	}

	var lsblkOutput StorageDiskTypeOutput
	if err := json.Unmarshal(output, &lsblkOutput); err != nil {
		log.Printf("Utils:getDiskType(), Failed to marshall %v: %v", output, err)
		return "", "", err
	}

	for _, device := range lsblkOutput.BlockDevices {
		if device.Type == blockType && !device.Removable {
			diskType = ssdDisk
			if device.Rotational {
				diskType = hddDisk
			}
			diskName = device.Name
		}
	}
	return diskType, diskName, nil
}

// getDiskSpeed() returns the read/write sectors in bytes
func getDiskSpeed(name string) (uint64, uint64, error) {

	ioCounters, err := disk.IOCounters()
	if err != nil {
		log.Printf("Error getting disk I/O counters: %v\n", err)
		return 0, 0, err
	}
	totalReadSectors := ioCounters[name].ReadBytes
	totalWriteSectors := ioCounters[name].WriteBytes
	return totalReadSectors, totalWriteSectors, nil
}
