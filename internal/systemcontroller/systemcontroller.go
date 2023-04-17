/*
 * Copyright (c) 2021 Siemens AG
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package systemcontroller

import (
	"context"
	"log"
	"strings"
	clientfactory "systemservice/internal/clientfactory"
	"systemservice/internal/common/utils"
	"time"
)

const shell = "bash"
const systemctl = "systemctl"
const reboot = "reboot"
const shutdown = "poweroff"
const journalctl = "journalctl"
const dockerDataRootDir = "docker info --format '{{json .DockerRootDir}}'"
const pathDockerData = "edge-iot-core/Data/edgeresetflag"
const pathDeviceName = "/var/device.name"
const pathDeviceId = "/var/device.id"
const commandOfRemove = "rm -rf"

// SystemController is the controller utility
type SystemController struct {
	fs utils.FileSystem
	ut utils.Utils
}

// NewSystemController to create a new instance of SystemController
func NewSystemController(fsVal utils.FileSystem, utVal utils.Utils) *SystemController {
	var systemcontroller = SystemController{fs: fsVal, ut: utVal}
	return &systemcontroller
}

// RestartDevice method performs device restart operation
func (s SystemController) RestartDevice() error {
	log.Println("systemcontroller:RestartDevice(), Enter")
	//Call reboot device
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("systemcontroller:RestartDevice(), Restart in Progress... ")
		command := systemctl + " " + reboot
		s.ut.Commander(command)
	}()
	log.Println("systemcontroller:RestartDevice() Leave")
	return nil
}

// ShutdownDevice method performs device shutdown operation
func (s SystemController) ShutdownDevice() error {
	//Call shutdown device
	log.Println("systemcontroller:ShutdownDevice(), Enter")
	command := systemctl + " " + shutdown
	out, err := s.ut.Commander(command)

	if err != nil {
		log.Printf("systemcontroller:ShutdownDevice(), output: %s error: %s ", string(out), err.Error())
	}
	log.Println("systemcontroller:ShutdownDevice(), Leave")
	return err
}

// HardReset method performs facroty reset operation
func (s SystemController) HardReset(ctx context.Context, clients *clientfactory.ClientPack) error {
	var err error
	var errflag error

	log.Println("All the contents regarding users/applications should be deleted from regarding directories/mountpoints.")

	if err = s.RemoveContent(); err != nil {
		log.Println("systemcontroller HardReset() RemoveContent Failed, error: ", err.Error())
	}

	//Restart Device
	s.RestartDevice()

	return errflag
}

// RemoveContent removes the edgeresetflag file
func (s SystemController) RemoveContent() error {

	// Get DataRootDir from 'docker info' output
	rootDirOfDockerData := []string{}
	out, errDocker := s.ut.Commander(dockerDataRootDir)
	if errDocker != nil {
		log.Printf("docker info can not be queried, output : %s, error : %s",
			strings.Trim((strings.TrimSuffix(string(out), "\n")), "\""), errDocker.Error())
		return errDocker
	}
	log.Printf("Queried docker info, found DockerRootDir : %s", string(out))

	// Construct the absolute path of file to be removed
	// removes both "" and \n
	rootDirOfDockerData = append(rootDirOfDockerData, strings.Trim((strings.TrimSuffix(string(out), "\n")), "\""))
	rootDirOfDockerData = append(rootDirOfDockerData, pathDockerData)

	absolutePathOfTheDockerData := strings.Join(rootDirOfDockerData, "/")

	filesToBeRemoved := []string{absolutePathOfTheDockerData, pathDeviceName, pathDeviceId}
	log.Printf("Will be removed at the end of the HardReset : %s, %s, %s ", absolutePathOfTheDockerData, pathDeviceName, pathDeviceId)

	// Remove Files
	err := s.RemoveFiles(filesToBeRemoved)
	if err != nil {
		return err
	}

	return nil
}

func (s SystemController) RemoveFiles(removeList []string) error {

	//Remove files on list
	for _, removeElement := range removeList {
		commandRemoveFile := commandOfRemove + " " + removeElement
		output, errRemoveFile := s.ut.Commander(commandRemoveFile)
		if errRemoveFile != nil {
			log.Printf("%s can not be removed, output : %s, error : %s", removeElement, string(output), errRemoveFile.Error())
			return errRemoveFile
		}

		log.Printf("%s is removed succesfully", removeElement)
	}

	return nil
}
