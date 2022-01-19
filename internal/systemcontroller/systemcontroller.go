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

/*
type OsUtils struct{}

func (OsUtils) Commander(command string) ([]byte, error) {
	out, err := exec.Command(shell, "-c", command).Output()
	return out, err
}
*/
const shell = "bash"
const systemctl = "systemctl"
const reboot = "reboot"
const shutdown = "poweroff"
const journalctl = "journalctl"
const dockerDataRootDir = "docker info --format '{{json .DockerRootDir}}'"
const removeDuringHardReset = "edge-iot-core/Data/edgeresetflag"
const remove = "rm -rf"


//SystemController is the controller utility
type SystemController struct {
	fs utils.FileSystem
	ut utils.Utils
}

//NewSystemController to create a new instance of SystemController
func NewSystemController(fsVal utils.FileSystem, utVal utils.Utils) *SystemController {
	var systemcontroller = SystemController{fs: fsVal, ut: utVal}
	return &systemcontroller
}

//RestartDevice method performs device restart operation
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

//ShutdownDevice method performs device shutdown operation
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

//HardReset method performs facroty reset operation
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
	toBeRemoved := []string{}
	out, errDocker := s.ut.Commander(dockerDataRootDir)
	if errDocker != nil {
		log.Printf("docker info can not be queried, output : %s, error : %s",
			strings.Trim((strings.TrimSuffix(string(out), "\n")), "\""), errDocker.Error())
		return errDocker
	}
	log.Printf("Queried docker info, found DockerRootDir : %s", string(out))

	// Construct the absolute path of file to be removed
	// removes both "" and \n
	toBeRemoved = append(toBeRemoved, strings.Trim((strings.TrimSuffix(string(out), "\n")), "\""))
	toBeRemoved = append(toBeRemoved, removeDuringHardReset)

	removeThis := strings.Join(toBeRemoved, "/")
	log.Printf("Will be removed at the end of the HardReset : %s", removeThis)

	// Remove
	command := remove + " " + removeThis
	out, errRemove := s.ut.Commander(command)
	if errRemove != nil {
		log.Printf("%s can not be removed, output : %s, error : %s", removeThis, string(out), errRemove.Error())
		return errRemove
	}

	// Exit
	log.Printf("%s is removed succesfully", removeThis)
	return nil

}

