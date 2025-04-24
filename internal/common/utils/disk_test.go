/*
 * Copyright © Siemens 2024 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package utils

import (
	"testing"
)

func Test_DiskUsage(t *testing.T) {
	path := "/"

	diskStatus, _ := DiskUsage(path)
	if diskStatus.All == 0 {
		t.Errorf("Expected disk total %v, but got %d", diskStatus, diskStatus.All)
	}
}

func Test_getDiskType(t *testing.T) {
	notExpectedValue := ""
	diskType, diskName, _ := getDiskInfo()
	if diskType == notExpectedValue {
		t.Errorf("Expected disk type %s, but got %s", notExpectedValue, diskType)
	}
	if diskName == notExpectedValue {
		t.Errorf("Expected disk name %s, but got %s", notExpectedValue, diskName)
	}
}

func Test_getDiskSpeed(t *testing.T) {
	_, diskName, _ := getDiskInfo()
	readSectors, writeSectors, _ := getDiskSpeed(diskName)
	if readSectors == 0 {
		t.Errorf("Expected read sectors to be %d, but got 0", readSectors)
	}

	if writeSectors == 0 {
		t.Errorf("Expected write sectors to be %d, but got 0", writeSectors)
	}
}

func Test_DiskUsage_Error(t *testing.T) {
	path := "/abc"

	diskStatus, _ := DiskUsage(path)
	if diskStatus.All != 0 {
		t.Errorf("Expected disk total %v, but got %d", diskStatus, diskStatus.All)
	}
}
