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
	systemapi "systemcontroller/api"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/grpc"
)

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

	if len(os.Args) < 4 {
		log.Println("Unexpected argument! please Enter : restart, shutdown or hard")
		return
	}
	if os.Args[3] == "hard" {
		response, err := sysClient.HardReset(context.Background(), new(emptypb.Empty))
		if err != nil {
			log.Printf("action could not performed %v", err)
		} else {
			log.Println("message sent : ", response)
			log.Println("DONE!")
		}
	} else if os.Args[3] == "restart" {
		response, err := sysClient.RestartDevice(context.Background(), new(emptypb.Empty))
		if err != nil {
			log.Printf("action could not performed %v", err)
		} else {
			log.Println("message sent : ", response)
			log.Println("DONE!")
		}
	} else if os.Args[3] == "shutdown" {
		response, err := sysClient.ShutdownDevice(context.Background(), new(emptypb.Empty))
		if err != nil {
			log.Printf("action could not performed %v", err)
		} else {
			log.Println("message sent : ", response)
			log.Println("DONE!")
		}
	} else {
		log.Println("Unexpected argument! please Enter : restart, shutdown or hard")
	}

}
