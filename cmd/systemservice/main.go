/*
 * Copyright © Siemens 2020 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package main

import (
	"os"
	systemservice "systemservice/app"
	"systemservice/internal/clientfactory"
)

func main() {

	systemServiceApp := systemservice.CreateServiceApp(clientfactory.ClientFactoryImpl{})
	systemServiceApp.StartApp()
	systemServiceApp.StartGRPC(os.Args)
}
