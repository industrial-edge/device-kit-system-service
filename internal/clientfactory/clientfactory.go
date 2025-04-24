/*
 * Copyright © Siemens 2020 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package clientfactory

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	socketDir string = "/var/run/devicemodel/"
)

type ClientPack struct {
	/*
		// If any clients needed, can be declared on this phase
		// such as:
		LoggingClient v1.LoggingClient
	*/
}

type ClientFactory interface {
	CreateClients() *ClientPack
}

type ClientFactoryImpl struct {
}

func createConnection(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			return net.Dial("unix", addr)
		}))
	if err != nil {
		log.Printf("Could not connect to: %v", err)
	} else {
		log.Printf("connected to " + address)
	}
	return conn
}

func (o ClientFactoryImpl) CreateClients() *ClientPack {
	/*
		//If any clients needed, can be initialized on this phase
		lClient := v1.NewLogginClient()
		pack := ClientPack{
			LogginClient : lClient,
		}
	*/

	pack := ClientPack{}
	return &pack
}
