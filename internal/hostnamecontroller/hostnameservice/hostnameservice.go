/*
 * Copyright © Siemens 2024 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package hostnameservice

import (
	"fmt"
	"strings"

	"systemservice/internal/common/utils"
)

// HostnameService struct to hold hostname details and dependencies
type HostnameService struct {
	hostname string
	OsUtils  utils.Utils
	FileSys  utils.FileSystem
}

const (
	hostsFilePath    = "/etc/hosts"
	hostnameFilePath = "/etc/hostname"
)

// NewHostnameService constructor
func NewHostnameService(osUtils utils.Utils, fileSys utils.FileSystem) *HostnameService {
	return &HostnameService{
		OsUtils: osUtils,
		FileSys: fileSys,
	}
}

func (h *HostnameService) Get() (string, error) {
	hostname, err := h.OsUtils.OsHostname()
	if err != nil {
		return "", fmt.Errorf("failed to get hostname: %w", err)
	}
	return hostname, nil
}

// Update is the main method to set the hostname
func (h *HostnameService) Update(hostname string) error {
	h.hostname = hostname
	if err := h.updateWithHostname(); err != nil {
		return err
	}
	return nil
}

// updateWithHostname uses the hostname command to set the hostname
func (h *HostnameService) updateWithHostname() error {
	if err := h.writeToFile(hostnameFilePath, h.hostname); err != nil {
		return err
	}

	if _, err := h.OsUtils.Commander(fmt.Sprintf("hostname %s", h.hostname)); err != nil {
		return fmt.Errorf("failed to set hostname with hostname command: %w", err)
	}

	if err := h.updateHostsFile(); err == nil {
		return h.OsUtils.SetHostnameEnv("HOSTNAME", h.hostname)
	} else {
		return err
	}
}

func (h *HostnameService) writeToFile(filePath, content string) error {
	originalContent, err := h.FileSys.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", filePath, err)
	}

	if len(originalContent) > 0 {
		if _, err := h.OsUtils.Commander(fmt.Sprintf("sed -i '/^.*$/d' %s", filePath)); err != nil {
			return fmt.Errorf("failed to clear %s: %w", filePath, err)
		}
	}

	if err := h.FileSys.WriteFile(filePath, []byte(content), 0644); err != nil {
		if writeErr := h.FileSys.WriteFile(filePath, originalContent, 0644); writeErr != nil {
			return fmt.Errorf("failed to write to %s: %w; additionally, failed to rollback: %v", filePath, err, writeErr)
		}
		return fmt.Errorf("failed to write to %s: %w", filePath, err)
	}
	return nil
}

func (h *HostnameService) updateHostsFile() error {
	originalContent, err := h.FileSys.ReadFile(hostsFilePath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", hostsFilePath, err)
	}

	newContent := h.removeExistingEntries(originalContent)

	newContent = append(newContent, []byte(fmt.Sprintf("\n127.0.1.1 %s", h.hostname))...)
	if err := h.FileSys.WriteFile(hostsFilePath, newContent, 0644); err != nil {
		if writeErr := h.FileSys.WriteFile(hostsFilePath, originalContent, 0644); writeErr != nil {
			return fmt.Errorf("failed to update %s: %w; additionally, failed to rollback: %v", hostsFilePath, err, writeErr)
		}
		return fmt.Errorf("failed to update %s: %w", hostsFilePath, err)
	}
	return nil
}

func (h *HostnameService) removeExistingEntries(content []byte) []byte {
	lines := strings.Split(string(content), "\n")
	var newLines []string
	for _, line := range lines {
		if !strings.Contains(line, "127.0.1.1") {
			newLines = append(newLines, line)
		}
	}
	return []byte(strings.Join(newLines, "\n"))
}
