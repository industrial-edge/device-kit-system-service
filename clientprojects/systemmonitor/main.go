/*
 * Copyright (c) 2021 Siemens AG
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	systemapi "systemmonitor/api"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	// Set up a connection to the server.
	if len(os.Args) < 3 || os.Args[1] != "unix" && os.Args[1] != "tcp" {
		fmt.Println("You must give an argument when running golang file." +
			"[uds /var/run/devicemodel/edge.sock OR tcp localhost:50006]\nUsage:" +
			" go run main.go uds /var/run/devicemodel/edge.sock OR go run main.go tcp localhost:50006")
		return
	}

	var conn *grpc.ClientConn
	var err error
	if os.Args[1] == "tcp" {
		conn, err = grpc.Dial(os.Args[2], grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Did not connect: %v", err)
		} else {
			log.Print("connected to " + os.Args[2])
		}
	}
	if os.Args[1] == "unix" {
		conn, err = grpc.Dial(
			os.Args[2],
			grpc.WithInsecure(),
			grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
				return net.DialTimeout("unix", addr, timeout)
			}))
		if err != nil {
			log.Fatalf("Did not connect: %v", err)
		} else {
			log.Print("connected to " + os.Args[2])
		}
	}

	defer conn.Close()
	sysClient := systemapi.NewSystemServiceClient(conn)

	go func() {
		time.Sleep(time.Second * 2)
		//clear screen
		for {
			response, _ := sysClient.GetLimits(context.Background(), new(emptypb.Empty))
			response2, _ := sysClient.GetModelNumber(context.Background(), new(emptypb.Empty))
			response3, _ := sysClient.GetFirmwareInfo(context.Background(), new(emptypb.Empty))
			response4, _ := sysClient.GetResourceStats(context.Background(), new(emptypb.Empty))
			CallClear()
			log.Println("----------- Firmware Info -----------")
			log.Println("Firmware Version: ", response3.GetVersion())
			log.Println()
			log.Println()
			log.Println("----------- Model Info -----------")
			log.Println("Model Number: ", response2.GetModelNumber())
			log.Println()
			log.Println()
			log.Println("----------- Limit Values -----------")
			log.Println("Max Installed Applications: ", response.GetMaxInstalledApplications())
			log.Println("Max Cpu Usage Percentage:", response.GetMaxCpuUsagePerecentage())
			log.Println("Max Memory Usage :", response.GetMaxMemoryUsageInGB())
			log.Println("Max Running Applications:", response.GetMaxRunningApplications())
			log.Println("Max Storage Usage :", response.GetMaxStorageUsageInGB())
			log.Println()
			log.Println()
			log.Println("----------- Resource Monitoring -----------")
			log.Println("Uptime: ", response4.GetUpTime())
			log.Println()
			log.Println("##### CPU STATS")
			log.Println("Core Count: ", response4.Cpu.CoreCount)
			log.Println("Cpu ModelInfo: ", response4.Cpu.ModelInfo)
			log.Println("Used Cpu: %", response4.Cpu.UsedCpuPercentage)
			log.Println("Free Cpu: %", response4.Cpu.FreeCpuPercentage)
			log.Println()
			log.Println("##### MEMORY STATS")
			log.Printf("Mem Total Space: %6.3f GB", response4.Memory.TotalSpaceInGB)
			log.Printf("Mem Used Space: %6.3f GB", response4.Memory.UsedSpaceInGB)
			log.Printf("Mem Free Space: %6.3f GB", response4.Memory.FreeSpaceInGB)
			log.Printf("Mem Used: %6.2f ", response4.Memory.PercentageUsedSpace)
			log.Printf("Mem Free: %6.2f ", response4.Memory.PercentageFreeSpace)
			log.Println()
			log.Println("##### STORAGE STATS")
			log.Printf("Mem Total Space: %6.3f GB", response4.StorageDevices[0].TotalSpaceInGB)
			log.Printf("Mem Used Space: %6.3f GB", response4.StorageDevices[0].UsedSpaceInGB)
			log.Printf("Mem Free Space: %6.3f GB", response4.StorageDevices[0].FreeSpaceInGB)
			log.Printf("Mem Used: %6.2f ", response4.StorageDevices[0].PercentageUsedSpace)
			log.Printf("Mem Free: %6.2f ", response4.StorageDevices[0].PercentageFreeSpace)

			time.Sleep(time.Second * 1)
		}
	}()
	select {}
}
