/*
 * Copyright © Siemens 2023 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package common

const (
	pathDeviceName                  = "/var/device.name"
	pathDeviceID                    = "/var/device.id"
	pathResourcePluginConfiguration = "/etc/ie-resource-plugins/*"
	pathResourceManagerDb           = "/var/lib/ie-resource-manager/Allocation_Persistence.json"
	pathNetworkResourcePluginDb     = "/var/lib/ie-docker-network-plugin/network/files/iedge-kv.db"
)

func GetMandatoryDeletionPaths() []string {
	return []string{
		pathDeviceName,
		pathDeviceID,
		pathResourcePluginConfiguration,
		pathResourceManagerDb,
	}
}

func GetMandatoryTruncatePaths() []string {
	return []string{
		pathNetworkResourcePluginDb,
	}
}
