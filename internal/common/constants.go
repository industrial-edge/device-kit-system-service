package common

const (
	pathDeviceName        = "/var/device.name"
	pathDeviceID          = "/var/device.id"
	pathCpuResourcePlugin = "/etc/ie-resource-plugins/*"
)

func GetMandatoryPaths() []string {
	return []string{
		pathDeviceName,
		pathDeviceID,
		pathCpuResourcePlugin,
	}
}
