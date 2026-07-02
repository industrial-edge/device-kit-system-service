/*
 * Copyright © Siemens 2025 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package mocks

import (
	"errors"
	"fmt"
	"log"
	"strings"
	systemapi "systemservice/api/siemens_iedge_dmapi_v1"
)

// MockSystemInfo extends SystemInfo for testing purposes
type MockSystemInfo struct {
	MockCPUStatsFunc     func() (*systemapi.Cpu, error)
	MockMemoryStatFunc   func() (*systemapi.Resource, error)
	MockStorageStatsFunc func() []*systemapi.Resource
	MockUpTimeFunc       func() (string, error)
}

// GetResourceStats mimics the real SystemInfo.GetResourceStats behavior
func (m *MockSystemInfo) GetResourceStats() (*systemapi.Stats, error) {
	var errMsgs []string
	var finalError error

	cpu, cpuErr := m.getCPUStats()
	storage := m.getStorageStats()
	memory, memoryErr := m.getMemoryStat()
	upTime, uptimeErr := m.getUpTime()

	if cpuErr != nil {
		log.Printf("mock:GetResourceStats(), CPU error occurred: %v", cpuErr)
		errMsgs = append(errMsgs, fmt.Sprintf("CPU: %v", cpuErr.Error()))
	}
	if memoryErr != nil {
		log.Printf("mock:GetResourceStats(), Memory error occurred: %v", memoryErr)
		errMsgs = append(errMsgs, fmt.Sprintf("Memory: %v", memoryErr.Error()))
	}
	if uptimeErr != nil {
		log.Printf("mock:GetResourceStats(), Uptime error occurred: %v", uptimeErr)
		errMsgs = append(errMsgs, fmt.Sprintf("Uptime: %v", uptimeErr.Error()))
	}

	if len(errMsgs) > 0 {
		finalError = errors.New("mock:GetResourceStats() errors: " + strings.Join(errMsgs, "; "))
	}

	return &systemapi.Stats{
		Cpu:            cpu,
		StorageDevices: storage,
		Memory:         memory,
		UpTime:         upTime,
	}, finalError
}

// Override the SystemInfo methods for testing
func (m *MockSystemInfo) getCPUStats() (*systemapi.Cpu, error) {
	if m.MockCPUStatsFunc != nil {
		return m.MockCPUStatsFunc()
	}
	// Always return a default CPU object, never nil (matching real implementation)
	return &systemapi.Cpu{
		UsedCpuPercentage: 0,
		FreeCpuPercentage: 0,
		CoreCount:         0,
		ModelInfo:         "NOTDEFINED",
		IdleTime:          0,
		Frequency:         0,
	}, nil
}

func (m *MockSystemInfo) getMemoryStat() (*systemapi.Resource, error) {
	if m.MockMemoryStatFunc != nil {
		return m.MockMemoryStatFunc()
	}
	return &systemapi.Resource{}, nil
}

func (m *MockSystemInfo) getStorageStats() []*systemapi.Resource {
	if m.MockStorageStatsFunc != nil {
		return m.MockStorageStatsFunc()
	}
	return []*systemapi.Resource{}
}

func (m *MockSystemInfo) getUpTime() (string, error) {
	if m.MockUpTimeFunc != nil {
		return m.MockUpTimeFunc()
	}
	return "0 days, 0 hours, 0 minutes", nil
}