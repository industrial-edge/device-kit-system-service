/*
 * Copyright (c) 2021 Siemens AG
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"sync"

	"log"
	"net"

	v1 "systemservice/api/siemens_iedge_dmapi_v1"
	clientfct "systemservice/internal/clientfactory"
	"systemservice/internal/common/utils"
	sysController "systemservice/internal/systemcontroller"
	sysInfo "systemservice/internal/systeminfo"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//DeviceModelService interface to start operations
type DeviceModelService interface {
	StartGRPC(args []string)
	StartApp()
}

type systemServer struct {
	IsysController systemControllerAPI
	IsysInfo       systemInfoAPI
	Clients        *clientfct.ClientPack
	sync.Mutex
}

//MainApp represents the Main Application
type MainApp struct {
	serverInstance *systemServer
	done           chan bool
}

type systemInfoAPI interface {
	GetResourceStats() (*v1.Stats, error)
	GetModelNumber() (*v1.ModelNumber, error)
	GetLimits() (*v1.Limits, error)
	GetFirmwareInfo() (*v1.FirmwareInfo, error)
	GetLogFile(request *v1.LogRequest) (*v1.LogResponse, error)
}

type systemControllerAPI interface {
	RestartDevice() error
	ShutdownDevice() error
	HardReset(context.Context, *clientfct.ClientPack) error
}

//CreateServiceApp is used to start a new service application from main.go
func CreateServiceApp(factory clientfct.ClientFactory) *MainApp {
	app := MainApp{
		serverInstance: &systemServer{Clients: factory.CreateClients()},
	}
	var CtrlOsfs utils.FileSystem = utils.OsFS{}
	var CtrlUt utils.Utils = utils.OsUtils{}
	app.serverInstance.IsysController = sysController.NewSystemController(CtrlOsfs, CtrlUt)
	var InfoOsfs utils.FileSystem = utils.OsFS{}
	var InfoUt utils.Utils = utils.OsUtils{}
	app.serverInstance.IsysInfo = sysInfo.NewSystemInfo(InfoOsfs, InfoUt)

	app.done = make(chan bool)

	return &app
}

func chownSocket(address string, userName string, groupName string) error {
	us, err1 := user.Lookup(userName)
	group, err2 := user.LookupGroup(groupName)
	if err1 == nil && err2 == nil {
		uid, _ := strconv.Atoi(us.Uid)
		gid, _ := strconv.Atoi(group.Gid)
		err3 := os.Chmod(address, os.FileMode.Perm(0660))
		err4 := os.Chown(address, uid, gid)
		if err3 != nil || err4 != nil {
			return errors.New("File permissions failed")
		}
		log.Println(uid, " : ", gid)
		return nil

	}
	return errors.New("File permissions failed")

}

//StartGRPC initiates GRPC relevant operations
func (app *MainApp) StartGRPC(args []string) error {
	const message string = "ERROR: Could not start monitor with bad arguments! \n " +
		"Sample usage:\n  ./ntpservice unix /tmp/devicemodel/ntp.socket \n" +
		"  ./ntpservice tcp localhost:50006"

	if len(args) != 3 {
		fmt.Println(message)
		return errors.New("parameter not supported")
	}
	typeOfConnection := args[1]
	address := args[2]
	if typeOfConnection != "unix" && typeOfConnection != "tcp" {
		fmt.Println(message)
		return errors.New("parameter not supported: " + typeOfConnection)
	}

	if typeOfConnection == "unix" {

		if err := os.RemoveAll(os.Args[2]); err != nil {
			return errors.New("socket could not removed: " + typeOfConnection)
		}

	}

	lis, err := net.Listen(typeOfConnection, address)

	if err != nil {
		log.Println("Failed to listen: ", err.Error())
		return errors.New("Failed to listen: " + err.Error())

	}
	if typeOfConnection == "unix" {
		err = chownSocket(address, "root", "docker")
		if err != nil {
			return err
		}
	}

	log.Print("Started listening on : ", typeOfConnection, " - ", address)
	s := grpc.NewServer()

	v1.RegisterSystemServiceServer(s, app.serverInstance)
	if err := s.Serve(lis); err != nil {
		log.Printf("Failed to serve: %v", err)
		return errors.New("Failed to serve: " + err.Error())
	}

	return nil
}

//StartApp initiates App relevant operations and listens for the actions
func (app *MainApp) StartApp() {
	//wait for app.done signal
	go func() {
		for {
			select {
			case <-app.done:
				log.Println("app done!")
				return
			}
		}
	}()
}

// GRPC method implementations ################################################################################
// ############################################################################################################

//Implementation of RPC method given systemapi proto file
func (s systemServer) RestartDevice(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	log.Println("RestartDevice() enter:")

	defer log.Println("RestartDevice() leave")
	if err := s.IsysController.RestartDevice(); err != nil {
		log.Println("RestartDevice() RPC Failure err: ", err)
		return &empty.Empty{}, status.New(codes.Internal, "Failed to Restart").Err()
	}

	return &empty.Empty{}, status.New(codes.OK, "fine").Err()
}

//Implementation of RPC method given systemapi proto file
func (s systemServer) ShutdownDevice(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	log.Println("ShutdownDevice() enter:")

	defer log.Println("ShutdownDevice() leave")
	if err := s.IsysController.ShutdownDevice(); err != nil {
		log.Println("ShutdownDevice() RPC Failure err: ", err)
		return &empty.Empty{}, status.New(codes.Internal, "Failed to Shutdown").Err()
	}

	return &empty.Empty{}, status.New(codes.OK, "fine").Err()
}

//Implementation of RPC method given systemapi proto file
func (s systemServer) HardReset(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	log.Println("HardReset() enter:")

	defer log.Println("HardReset() leave")
	if err := s.IsysController.HardReset(ctx, s.Clients); err != nil {
		log.Println("HardReset() RPC Failure err: ", err)
		return &empty.Empty{}, status.New(codes.Internal, "Failed to HardReset err:"+err.Error()).Err()
	}

	return &empty.Empty{}, status.New(codes.OK, "fine").Err()
}

//Implementation of RPC method given systemapi proto file
func (s systemServer) GetModelNumber(ctx context.Context, e *empty.Empty) (*v1.ModelNumber, error) {
	log.Println("GetModelNumber() enter:")

	defer log.Println("GetModelNumber() leave")
	retval, err := s.IsysInfo.GetModelNumber()
	if err != nil {
		log.Println("GetModelNumber() RPC Failure err: ", err)
		return retval, status.New(codes.Internal, "Failed to GetModelNumber").Err()
	}

	return retval, status.New(codes.OK, "fine").Err()
}

//Implementation of RPC method given systemapi proto file
func (s systemServer) GetFirmwareInfo(ctx context.Context, e *empty.Empty) (*v1.FirmwareInfo, error) {
	log.Println("GetFirmwareInfo() enter:")
	defer log.Println("GetFirmwareInfo() leave")
	retval, err := s.IsysInfo.GetFirmwareInfo()
	if err != nil {
		log.Println("GetFirmwareInfo() RPC Failure err: ", err)
		return retval, status.New(codes.Internal, "Failed to GetFirmwareInfo").Err()
	}

	return retval, status.New(codes.OK, "fine").Err()
}

//Implementation of RPC method given systemapi proto file
func (s systemServer) GetResourceStats(ctx context.Context, e *empty.Empty) (*v1.Stats, error) {
	log.Println("GetResourceStats() enter:")

	defer log.Println("GetResourceStats() leave")
	retval, err := s.IsysInfo.GetResourceStats()
	if err != nil {
		log.Println("GetResourceStats() RPC Failure err: ", err)
		return retval, status.New(codes.Internal, "Failed to GetResourceStats").Err()
	}

	return retval, status.New(codes.OK, "fine").Err()
}

//Implementation of RPC method given systemapi proto file
func (s systemServer) GetLimits(ctx context.Context, e *empty.Empty) (*v1.Limits, error) {
	log.Println("GetLimits() enter:")

	defer log.Println("GetLimits() leave")
	retval, err := s.IsysInfo.GetLimits()
	if err != nil {
		log.Println("GetLimits() RPC Failure err: ", err)
		return retval, status.New(codes.Internal, "Failed to GetLimits").Err()
	}

	return retval, status.New(codes.OK, "fine").Err()
}

// GetLogFile  RPC method implementation
func (s systemServer) GetLogFile(ctx context.Context, req *v1.LogRequest) (*v1.LogResponse, error) {
	log.Println("GetLogFile() rpc method called")
	defer log.Println("GetLogFile() leave")
	s.Lock()
	retVal, err := s.IsysInfo.GetLogFile(req)
	if err != nil {
		log.Println("GetLogFile() RPC Failure err: ", err)
		return retVal, status.New(codes.Internal, "Failed to GetLogFile").Err()
	}
	s.Unlock()
	return retVal, status.New(codes.OK, "fine").Err()
}
