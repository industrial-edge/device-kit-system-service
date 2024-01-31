/*
 * Copyright (c) 2021 Siemens AG
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package systemcontroller

import (
	"context"
	"fmt"
	"log"
	"strings"
	"systemservice/internal/clientfactory"
	"systemservice/internal/common"
	"systemservice/internal/common/utils"
	"time"
)

const (
	blank           = " "
	systemctl       = "systemctl" + blank
	shutdown        = systemctl + "poweroff"
	reboot          = systemctl + "reboot"
	removeCommand   = "rm -rf" + blank
	dockerDataQuery = "docker info --format '{{json .DockerRootDir}}'"
	pathDockerData  = "edge-iot-core/Data/edgeresetflag"
)

// SystemController is the controller utility
type SystemController struct {
	fs utils.FileSystem
	ut utils.Utils
}

// NewSystemController to create a new instance of SystemController
func NewSystemController(fsVal utils.FileSystem, utVal utils.Utils) *SystemController {
	return &SystemController{fs: fsVal, ut: utVal}
}

// ShutdownDevice method performs device shutdown operation
func (s SystemController) ShutdownDevice() error {
	//Call shutdown device
	log.Println("systemcontroller:ShutdownDevice(), Enter")
	out, err := s.ut.Commander(shutdown)

	if err != nil {
		log.Printf("systemcontroller:ShutdownDevice(), output: %s error: %s ", string(out), err.Error())
	}
	log.Println("systemcontroller:ShutdownDevice(), Leave")
	return err
}

// RestartDevice method performs device restart operation
func (s SystemController) RestartDevice() error {
	log.Println("systemcontroller:RestartDevice(), Enter")
	// Call reboot device
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("systemcontroller:RestartDevice(), Restart in Progress... ")

		output, err := s.ut.Commander(reboot)
		if err != nil {
			log.Printf("Error while rebooting device: %s\nOutput: %s", err, string(output))
			return
		}

		log.Println("systemcontroller:RestartDevice(), Restart completed successfully")
	}()

	log.Println("systemcontroller:RestartDevice() Leave")
	return nil
}

// HardReset method performs factory reset operation
func (s SystemController) HardReset(ctx context.Context, clients *clientfactory.ClientPack) error {
	log.Println("All the contents regarding users/applications should be deleted from regarding directories/mount-points.")

	if err := s.RemoveContent(); err != nil {
		log.Printf("systemcontroller HardReset() RemoveContent Failed, error: %s", err.Error())
		return err
	}

	// Restart Device
	if err := s.RestartDevice(); err != nil {
		log.Printf("systemcontroller HardReset() RestartDevice Failed, error: %s", err.Error())
		return err
	}

	log.Println("HardReset completed successfully.")
	return nil
}

// RemoveContent removes specific files.
func (s SystemController) RemoveContent() error {
	dockerRootDir, err := s.GetDockerRootDir()
	if err != nil {
		log.Printf("Error occured: DockerRootDir could not be obtained: %s", err)
		return err
	}

	// Construct the absolute path of the file to be removed
	absolutePathOfDockerData := fmt.Sprintf("%s/%s", dockerRootDir, pathDockerData)
	filesToRemove := common.GetMandatoryPaths()
	filesToRemove = append(filesToRemove, absolutePathOfDockerData)

	// Remove specific files
	for _, file := range filesToRemove {
		if err := s.RemoveFile(file); err != nil {
			log.Printf("Error occured: '%s' could not be removed: %s", file, err)
			return err
		}
	}

	return nil
}

// RemoveFile removes a file
func (s SystemController) RemoveFile(filePath string) error {
	removeCommand := removeCommand + filePath
	output, err := s.ut.Commander(removeCommand)
	if err != nil {
		log.Printf("Error occured: %s encountered an error during removal, output: %s, error: %s", filePath, string(output), err)
		return err
	}

	log.Printf("%s has been successfully removed", filePath)
	return nil
}

// GetDockerRootDir gets the Docker root directory
func (s SystemController) GetDockerRootDir() (string, error) {
	out, err := s.ut.Commander(dockerDataQuery)
	if err != nil {
		return "", err
	}
	return strings.Trim(strings.TrimSuffix(string(out), "\n"), "\""), nil
}
