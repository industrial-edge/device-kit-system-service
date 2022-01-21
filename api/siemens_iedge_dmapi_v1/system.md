# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [System.proto](#System.proto)
    - [Cpu](#siemens.iedge.dmapi.system.v1.Cpu)
    - [FirmwareInfo](#siemens.iedge.dmapi.system.v1.FirmwareInfo)
    - [Limits](#siemens.iedge.dmapi.system.v1.Limits)
    - [LogRequest](#siemens.iedge.dmapi.system.v1.LogRequest)
    - [LogResponse](#siemens.iedge.dmapi.system.v1.LogResponse)
    - [ModelNumber](#siemens.iedge.dmapi.system.v1.ModelNumber)
    - [Resource](#siemens.iedge.dmapi.system.v1.Resource)
    - [Stats](#siemens.iedge.dmapi.system.v1.Stats)
  
    - [SystemService](#siemens.iedge.dmapi.system.v1.SystemService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="System.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## System.proto



<a name="siemens.iedge.dmapi.system.v1.Cpu"></a>

### Cpu
Cpu type contains Cpu utilization at the current moment.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| usedCpuPercentage | [float](#float) |  | Percentage of used CPU e.g: 20.0 |
| freeCpuPercentage | [float](#float) |  | Percentage of available CPU e.g: 80.0 |
| coreCount | [int32](#int32) |  | Total available core count for CPU.e.g 2C/4T CPU value will be 4 |
| modelInfo | [string](#string) |  | intel x64 etc.. |






<a name="siemens.iedge.dmapi.system.v1.FirmwareInfo"></a>

### FirmwareInfo
FirmwareInfo contains Firmware Version.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  | Firmware version. |






<a name="siemens.iedge.dmapi.system.v1.Limits"></a>

### Limits
System Limits for the EdgeRuntime.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| maxInstalledApplications | [int32](#int32) |  | Maximum allowed number of installed edge applications. |
| maxRunningApplications | [int32](#int32) |  | Maximum allowed number of running edge applications. |
| maxMemoryUsageInGB | [float](#float) |  | Maximum allowed memory usage in Gigabytes. |
| maxStorageUsageInGB | [float](#float) |  | Maximum allowed disk usage in Gigabytes. |
| maxCpuUsagePerecentage | [float](#float) |  | Maximum allowed percentage of CPU usage. |






<a name="siemens.iedge.dmapi.system.v1.LogRequest"></a>

### LogRequest
LogRequest type, determines the destination path for saving log file.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| saveFolderPath | [string](#string) |  | Folder path for saving gathered logs. |






<a name="siemens.iedge.dmapi.system.v1.LogResponse"></a>

### LogResponse
LogResponse type, contains the full path for the collected log archive.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logPath | [string](#string) |  | Full file path for collected log archive. |






<a name="siemens.iedge.dmapi.system.v1.ModelNumber"></a>

### ModelNumber
ModelNumber type indicates device specific model information.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| modelNumber | [string](#string) |  | Can be MLFB for SIEMENS devices, for 3rd party vendors it can be any model information. |






<a name="siemens.iedge.dmapi.system.v1.Resource"></a>

### Resource
System Resource , memory or storage.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| totalSpaceInGB | [float](#float) |  | Total capacity of storage device in Gigabytes e.g: 3.5 |
| freeSpaceInGB | [float](#float) |  | Free space of storage device in Gigabytes e.g: 40.4 |
| usedSpaceInGB | [float](#float) |  | Used space of storage device in Gigabytes e.g: 23.2 |
| percentageFreeSpace | [float](#float) |  | Percentage of available space e.g: 3.5 |
| percentageUsedSpace | [float](#float) |  | Percentage of used space e.g: 96.5 |






<a name="siemens.iedge.dmapi.system.v1.Stats"></a>

### Stats
System Utilization type. Cpu, storage and memory utilization.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cpu | [Cpu](#siemens.iedge.dmapi.system.v1.Cpu) |  | Cpu Utilization |
| storageDevices | [Resource](#siemens.iedge.dmapi.system.v1.Resource) | repeated | StorageDevices array of Resource type. |
| memory | [Resource](#siemens.iedge.dmapi.system.v1.Resource) |  | RAM Utilization Information |
| upTime | [string](#string) |  | Elapsed time since the device is started. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="siemens.iedge.dmapi.system.v1.SystemService"></a>

### SystemService
System service ,uses a UNIX Domain Socket "/var/run/devicemodel/system.sock" for GRPC communication.
protoc  generates both client and server instance for this Service.
GRPC Status codes : https://developers.google.com/maps-booking/reference/grpc-api/status_codes .

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| RestartDevice | [.google.protobuf.Empty](#google.protobuf.Empty) | [.google.protobuf.Empty](#google.protobuf.Empty) | Restarts the device |
| ShutdownDevice | [.google.protobuf.Empty](#google.protobuf.Empty) | [.google.protobuf.Empty](#google.protobuf.Empty) | ShutsDown the device. |
| HardReset | [.google.protobuf.Empty](#google.protobuf.Empty) | [.google.protobuf.Empty](#google.protobuf.Empty) | Performs host side actions in addition to edge-core for hard reset. e.g: cleaning hard-reset flag(mandatory) ,custom device builder steps(optional) and finally reboots the system(mandatory). |
| GetModelNumber | [.google.protobuf.Empty](#google.protobuf.Empty) | [ModelNumber](#siemens.iedge.dmapi.system.v1.ModelNumber) | Returns model number (mlfb) for siemens or any type model for 3rd party vendors. |
| GetFirmwareInfo | [.google.protobuf.Empty](#google.protobuf.Empty) | [FirmwareInfo](#siemens.iedge.dmapi.system.v1.FirmwareInfo) | Returns firmware information of currently installed firmware |
| GetResourceStats | [.google.protobuf.Empty](#google.protobuf.Empty) | [Stats](#siemens.iedge.dmapi.system.v1.Stats) | Returns current Cpu, Memory, Uptime and Storage usage |
| GetLimits | [.google.protobuf.Empty](#google.protobuf.Empty) | [Limits](#siemens.iedge.dmapi.system.v1.Limits) | Returns limits for how many applications and how much cpu, ram and storage should be available for applications. |
| GetLogFile | [LogRequest](#siemens.iedge.dmapi.system.v1.LogRequest) | [LogResponse](#siemens.iedge.dmapi.system.v1.LogResponse) | Collects and compress all Journald logs (mandatory) from host ,(plus optional device specific log/report) and then returns a single file path for this new log archive. |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
