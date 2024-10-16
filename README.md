# Introduction

The IE Device Kit API provides the abstraction layer that decouples the Industrial Edge Runtime from the underlying Linux systems. This allows to adapt the runtime and its behavior to serve for the specific needs of different Industrial Edge products. 

The IE Device Kit API is based on gRPC which provides a modern intermediate process communication style for building distributed applications and microservices. The Industrial Edge platform provides and maintains the protobuf specification files for the APIs contained in the IE Device Kit. These protobuf specifications can be used to create stub implementations for both client and server in various programming languages. The Industrial Edge Runtime ships with a client side implementation of these APIs and expects the host system to provide a server side implementation.

Purpose of these repositories is to share reference implementation of IE Device Kit APIs. You can use existing implementation or adapt it based on your needs.
# IEDK System Service

System Service is a gRPC & Go based system resource tracker and system controller. It also includes a console application that shows the system status -`System Resource Stats`, `Model Number`, `System Limit Values`- of devices and a demo client for controlling device as following;

```bash

    //Restarts the device
    rpc RestartDevice(google.protobuf.Empty) returns(google.protobuf.Empty);
    
    //ShutsDown the device.
    rpc ShutdownDevice(google.protobuf.Empty) returns(google.protobuf.Empty);
    
    // Performs host side actions in addition to edge-core for hard reset. e.g: cleaning hard-reset flag(mandatory) ,custom device builder steps(optional) and finally reboots the system(mandatory). 
    rpc HardReset(google.protobuf.Empty) returns(google.protobuf.Empty);
 
	//Returns model number (mlfb) for siemens or any type model for 3rd party vendors.
    rpc GetModelNumber(google.protobuf.Empty) returns(ModelNumber);
	
	//Returns firmware information of currently installed firmware
    rpc GetFirmwareInfo(google.protobuf.Empty) returns(FirmwareInfo);
    
    //Returns current Cpu, Memory, Uptime and Storage usage
    rpc GetResourceStats(google.protobuf.Empty) returns(Stats);
    
    //Returns limits for how many applications and how much cpu, ram and storage should be available for applications.
    rpc GetLimits(google.protobuf.Empty) returns(Limits);

    //Returns device specific custom settings.
    rpc GetCustomSettings(google.protobuf.Empty) returns(google.protobuf.Any);

    //Applies device specific custom settings.
    rpc ApplyCustomSettings(google.protobuf.Any) returns(google.protobuf.Empty);
    
    //Collects and compress all Journald logs (mandatory) from host ,(plus optional device specific log/report) and then returns a single file path for this new log archive. 
    rpc GetLogFile(LogRequest) returns(LogResponse);

```

## Overview

_System Service_ is developed in the Go programming language and gRPC. More information can be found [here](https://grpc.io/docs/). The _System Service_ runs as a systemd service within the device that has a debian-based operating system.


## Getting Started

### Prerequisities

> - Setting up Go
> - Additional requirements to _run_ or _develop_ the project

### Building the service and know-how about other features

> Instructions how to build the deb package:
>
> - For generating a deb package, [goreleaser](https://goreleaser.com/intro/) tool is heavily used. For proper execution, goreleaser needs a TAG identifier which indicates version of the deb package. __TAG__ must obey the semantic versioning rules and must include ```${major}.${minor}.${hotfix}``` release identifiers. Only two commands are needed for generating deb package: `cd build/package`, `TAG=X.Y.Z make deb`. After running these commands, a __dist__ directory will be created under the __build/package__ directory and the deb package will be in dist directory. To install this generated deb package on the device as a daemon `cd dist`, `sudo apt install ./dm-system_X.Y.Z_linux_amd64.deb`
>
> Instructions how to use the make command:
>
> - There is a Makefile file under the __build/package__ directory. Unit tests, code coverage and many other similar features can be used via this Makefile. The following commands are used to view all features: `cd build/package`, `make help`


### Running the service

> To see the status and logs of the deb package running as daemon(systemd service) directly from the command line, the following commands can be run: `systemctl status dm-service`, `journalctl -fu dm-service`
## FAQ

### How do I get a proper firmware info?
In order to get a proper firmware info from O/S through this service, it is needed to add a __*VARIANT*__ key and corresponding value into the /etc/os-release file. <br>
The content can proposed as below:
```bash
$ cat /etc/os-release
PRETTY_NAME="Your OS 2.1 (buster)"
NAME="Your OS"
VERSION_ID="2.1"
VERSION="2.1 (buster)"
VERSION_CODENAME=buster
ID=your-os
HOME_URL="https://www.some-url.com"
SUPPORT_URL="https://www.some-url.com/support"
BUG_REPORT_URL="https://www.some-url.com/support"
VARIANT="your-os-1.1.0-21-amd64-develop"
```

### How do I configure limits?
Limits regarding resource usage, such as maximum cpu usage, maximum storage usage, etc. can be specified in different metrics in the `./cmd/systemservice/default.json` file in the project directory and can easily be manipulated. The resources can be limited are indicated with the json object __"Limits"__ as can be seen below. 

```bash
$ cat ./cmd/systemservice/default.json
{
"Limits":
{
   "MaxInstalledApplications":20,
   "MaxRunningApplications":10,
   "MaxMemoryUsageInGB":8.0,
   "MaxStorageUsageInGB":20.0,
   "MaxCpuUsagePerecentage":100
},
"MonitoredStorage":
{
   "Path":"/"
}
}
```

### How do I indicate the `path` to be monitored?
As can be seen the json object above, in order to provide a directory to be reported, one can specify only directory on the `./cmd/systemservice/default.json` file with the json object __"MonitoredStorage"__ and key __"Path"__. As of now, `root` directory (indicated as __`"/"`__) is being monitored.

# Contributing IE Device Kit Repository
Please check our [contribution guideline](CONTRIBUTING.md). 

# Contribution License Agreement
If you haven't previously signed the [Siemens Contributor License Agreement](https://cla-assistant.io/industrial-edge/) (CLA), the system will automatically prompt you to do so when you submit your Pull Request. This can be conveniently done through the CLA Assistant's online platform.
Once the CLA is signed, your Pull Request will automatically be cleared and made ready for merging if all other test stages succeed.

# How to be part of Siemens Industrial Edge Ecosystem
Please check [this](https://new.siemens.com/global/en/products/automation/topic-areas/industrial-edge.html) page to learn more information about Industrial Edge.
