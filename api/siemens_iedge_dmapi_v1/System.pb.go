//
// Copyright (c) Siemens 2023
// Licensed under the MIT license
// See LICENSE file in the top-level directory

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: System.proto

package siemens_iedge_dmapi_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ModelNumber type indicates device specific model information.
type ModelNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelNumber string `protobuf:"bytes,1,opt,name=modelNumber,proto3" json:"modelNumber,omitempty"` // Can be MLFB for SIEMENS devices, for 3rd party vendors it can be any model information.
}

func (x *ModelNumber) Reset() {
	*x = ModelNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_System_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelNumber) ProtoMessage() {}

func (x *ModelNumber) ProtoReflect() protoreflect.Message {
	mi := &file_System_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelNumber.ProtoReflect.Descriptor instead.
func (*ModelNumber) Descriptor() ([]byte, []int) {
	return file_System_proto_rawDescGZIP(), []int{0}
}

func (x *ModelNumber) GetModelNumber() string {
	if x != nil {
		return x.ModelNumber
	}
	return ""
}

// FirmwareInfo contains Firmware Version.
type FirmwareInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"` // Firmware version.
}

func (x *FirmwareInfo) Reset() {
	*x = FirmwareInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_System_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirmwareInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirmwareInfo) ProtoMessage() {}

func (x *FirmwareInfo) ProtoReflect() protoreflect.Message {
	mi := &file_System_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirmwareInfo.ProtoReflect.Descriptor instead.
func (*FirmwareInfo) Descriptor() ([]byte, []int) {
	return file_System_proto_rawDescGZIP(), []int{1}
}

func (x *FirmwareInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// System Resource , memory or storage.
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalSpaceInGB      float32 `protobuf:"fixed32,1,opt,name=totalSpaceInGB,proto3" json:"totalSpaceInGB,omitempty"`           // Total capacity of storage device in Gigabytes e.g: 3.5
	FreeSpaceInGB       float32 `protobuf:"fixed32,2,opt,name=freeSpaceInGB,proto3" json:"freeSpaceInGB,omitempty"`             // Free space of storage device in Gigabytes e.g: 40.4
	UsedSpaceInGB       float32 `protobuf:"fixed32,3,opt,name=usedSpaceInGB,proto3" json:"usedSpaceInGB,omitempty"`             // Used space of storage device in Gigabytes e.g: 23.2
	PercentageFreeSpace float32 `protobuf:"fixed32,4,opt,name=percentageFreeSpace,proto3" json:"percentageFreeSpace,omitempty"` // Percentage of available space e.g: 3.5
	PercentageUsedSpace float32 `protobuf:"fixed32,5,opt,name=percentageUsedSpace,proto3" json:"percentageUsedSpace,omitempty"` // Percentage of used space e.g: 96.5
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_System_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_System_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_System_proto_rawDescGZIP(), []int{2}
}

func (x *Resource) GetTotalSpaceInGB() float32 {
	if x != nil {
		return x.TotalSpaceInGB
	}
	return 0
}

func (x *Resource) GetFreeSpaceInGB() float32 {
	if x != nil {
		return x.FreeSpaceInGB
	}
	return 0
}

func (x *Resource) GetUsedSpaceInGB() float32 {
	if x != nil {
		return x.UsedSpaceInGB
	}
	return 0
}

func (x *Resource) GetPercentageFreeSpace() float32 {
	if x != nil {
		return x.PercentageFreeSpace
	}
	return 0
}

func (x *Resource) GetPercentageUsedSpace() float32 {
	if x != nil {
		return x.PercentageUsedSpace
	}
	return 0
}

// Cpu type contains Cpu utilization at the current moment.
type Cpu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsedCpuPercentage float32 `protobuf:"fixed32,1,opt,name=usedCpuPercentage,proto3" json:"usedCpuPercentage,omitempty"` // Percentage of used CPU e.g: 20.0
	FreeCpuPercentage float32 `protobuf:"fixed32,2,opt,name=freeCpuPercentage,proto3" json:"freeCpuPercentage,omitempty"` // Percentage of available CPU e.g: 80.0
	CoreCount         int32   `protobuf:"varint,3,opt,name=coreCount,proto3" json:"coreCount,omitempty"`                  // Total available core count for CPU.e.g  2C/4T CPU value will be 4
	ModelInfo         string  `protobuf:"bytes,4,opt,name=modelInfo,proto3" json:"modelInfo,omitempty"`                   // intel x64 etc..
}

func (x *Cpu) Reset() {
	*x = Cpu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_System_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cpu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cpu) ProtoMessage() {}

func (x *Cpu) ProtoReflect() protoreflect.Message {
	mi := &file_System_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cpu.ProtoReflect.Descriptor instead.
func (*Cpu) Descriptor() ([]byte, []int) {
	return file_System_proto_rawDescGZIP(), []int{3}
}

func (x *Cpu) GetUsedCpuPercentage() float32 {
	if x != nil {
		return x.UsedCpuPercentage
	}
	return 0
}

func (x *Cpu) GetFreeCpuPercentage() float32 {
	if x != nil {
		return x.FreeCpuPercentage
	}
	return 0
}

func (x *Cpu) GetCoreCount() int32 {
	if x != nil {
		return x.CoreCount
	}
	return 0
}

func (x *Cpu) GetModelInfo() string {
	if x != nil {
		return x.ModelInfo
	}
	return ""
}

// System Utilization type. Cpu, storage and memory utilization.
type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpu            *Cpu        `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty"`                       // Cpu Utilization
	StorageDevices []*Resource `protobuf:"bytes,2,rep,name=storageDevices,proto3" json:"storageDevices,omitempty"` // StorageDevices array of Resource type.
	Memory         *Resource   `protobuf:"bytes,3,opt,name=memory,proto3" json:"memory,omitempty"`                 // RAM Utilization Information
	UpTime         string      `protobuf:"bytes,4,opt,name=upTime,proto3" json:"upTime,omitempty"`                 // Elapsed time since the device is started.
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_System_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_System_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_System_proto_rawDescGZIP(), []int{4}
}

func (x *Stats) GetCpu() *Cpu {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *Stats) GetStorageDevices() []*Resource {
	if x != nil {
		return x.StorageDevices
	}
	return nil
}

func (x *Stats) GetMemory() *Resource {
	if x != nil {
		return x.Memory
	}
	return nil
}

func (x *Stats) GetUpTime() string {
	if x != nil {
		return x.UpTime
	}
	return ""
}

// System Limits for the EdgeRuntime.
type Limits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxInstalledApplications int32   `protobuf:"varint,1,opt,name=maxInstalledApplications,proto3" json:"maxInstalledApplications,omitempty"` // Maximum allowed number of installed edge applications.
	MaxRunningApplications   int32   `protobuf:"varint,2,opt,name=maxRunningApplications,proto3" json:"maxRunningApplications,omitempty"`     // Maximum allowed number of running edge applications.
	MaxMemoryUsageInGB       float32 `protobuf:"fixed32,3,opt,name=maxMemoryUsageInGB,proto3" json:"maxMemoryUsageInGB,omitempty"`            // Maximum allowed memory usage in Gigabytes.
	MaxStorageUsageInGB      float32 `protobuf:"fixed32,4,opt,name=maxStorageUsageInGB,proto3" json:"maxStorageUsageInGB,omitempty"`          // Maximum allowed disk usage in Gigabytes.
	MaxCpuUsagePerecentage   float32 `protobuf:"fixed32,5,opt,name=maxCpuUsagePerecentage,proto3" json:"maxCpuUsagePerecentage,omitempty"`    // Maximum allowed percentage of CPU usage.
}

func (x *Limits) Reset() {
	*x = Limits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_System_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Limits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Limits) ProtoMessage() {}

func (x *Limits) ProtoReflect() protoreflect.Message {
	mi := &file_System_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Limits.ProtoReflect.Descriptor instead.
func (*Limits) Descriptor() ([]byte, []int) {
	return file_System_proto_rawDescGZIP(), []int{5}
}

func (x *Limits) GetMaxInstalledApplications() int32 {
	if x != nil {
		return x.MaxInstalledApplications
	}
	return 0
}

func (x *Limits) GetMaxRunningApplications() int32 {
	if x != nil {
		return x.MaxRunningApplications
	}
	return 0
}

func (x *Limits) GetMaxMemoryUsageInGB() float32 {
	if x != nil {
		return x.MaxMemoryUsageInGB
	}
	return 0
}

func (x *Limits) GetMaxStorageUsageInGB() float32 {
	if x != nil {
		return x.MaxStorageUsageInGB
	}
	return 0
}

func (x *Limits) GetMaxCpuUsagePerecentage() float32 {
	if x != nil {
		return x.MaxCpuUsagePerecentage
	}
	return 0
}

// LogRequest type, determines the destination path for saving log file.
type LogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SaveFolderPath string `protobuf:"bytes,1,opt,name=saveFolderPath,proto3" json:"saveFolderPath,omitempty"` // Folder path for saving gathered logs.
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_System_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_System_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_System_proto_rawDescGZIP(), []int{6}
}

func (x *LogRequest) GetSaveFolderPath() string {
	if x != nil {
		return x.SaveFolderPath
	}
	return ""
}

// LogResponse type, contains the full path for the collected log archive.
type LogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogPath string `protobuf:"bytes,1,opt,name=logPath,proto3" json:"logPath,omitempty"` // Full file path for collected log archive.
}

func (x *LogResponse) Reset() {
	*x = LogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_System_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogResponse) ProtoMessage() {}

func (x *LogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_System_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogResponse.ProtoReflect.Descriptor instead.
func (*LogResponse) Descriptor() ([]byte, []int) {
	return file_System_proto_rawDescGZIP(), []int{7}
}

func (x *LogResponse) GetLogPath() string {
	if x != nil {
		return x.LogPath
	}
	return ""
}

var File_System_proto protoreflect.FileDescriptor

var file_System_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x28, 0x0a, 0x0c, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61,
	0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0xe2, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x47, 0x42, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x49, 0x6e, 0x47, 0x42, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x72, 0x65, 0x65, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x49, 0x6e, 0x47, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x66, 0x72,
	0x65, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x47, 0x42, 0x12, 0x24, 0x0a, 0x0d, 0x75,
	0x73, 0x65, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x47, 0x42, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x47,
	0x42, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x46,
	0x72, 0x65, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13,
	0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x46, 0x72, 0x65, 0x65, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x55, 0x73, 0x65, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x13, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65, 0x64,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x03, 0x43, 0x70, 0x75, 0x12, 0x2c, 0x0a,
	0x11, 0x75, 0x73, 0x65, 0x64, 0x43, 0x70, 0x75, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x75, 0x73, 0x65, 0x64, 0x43, 0x70,
	0x75, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x66,
	0x72, 0x65, 0x65, 0x43, 0x70, 0x75, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x66, 0x72, 0x65, 0x65, 0x43, 0x70, 0x75, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x72,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f,
	0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x34, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73,
	0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x70, 0x75,
	0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x4f, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73,
	0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x96, 0x02, 0x0a, 0x06, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x18, 0x6d, 0x61,
	0x78, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18, 0x6d, 0x61,
	0x78, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36, 0x0a, 0x16, 0x6d, 0x61, 0x78, 0x52, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x6d, 0x61, 0x78, 0x52, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e,
	0x0a, 0x12, 0x6d, 0x61, 0x78, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x47, 0x42, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x4d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x47, 0x42, 0x12, 0x30,
	0x0a, 0x13, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x47, 0x42, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x6d, 0x61, 0x78,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x47, 0x42,
	0x12, 0x36, 0x0a, 0x16, 0x6d, 0x61, 0x78, 0x43, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x50,
	0x65, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x16, 0x6d, 0x61, 0x78, 0x43, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x22, 0x34, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x61, 0x76, 0x65, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x73, 0x61, 0x76, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x22, 0x27,
	0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x32, 0x88, 0x06, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x52, 0x65, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x0e, 0x53, 0x68,
	0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x09,
	0x48, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x54, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x56, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2b, 0x2e, 0x73, 0x69, 0x65,
	0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x72, 0x6d, 0x77,
	0x61, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x50, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x24, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x4a, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x25,
	0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64,
	0x6d, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x12, 0x43, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x63, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x29, 0x2e, 0x73, 0x69,
	0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73,
	0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x2e, 0x3b, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x5f,
	0x69, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_System_proto_rawDescOnce sync.Once
	file_System_proto_rawDescData = file_System_proto_rawDesc
)

func file_System_proto_rawDescGZIP() []byte {
	file_System_proto_rawDescOnce.Do(func() {
		file_System_proto_rawDescData = protoimpl.X.CompressGZIP(file_System_proto_rawDescData)
	})
	return file_System_proto_rawDescData
}

var file_System_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_System_proto_goTypes = []interface{}{
	(*ModelNumber)(nil),   // 0: siemens.iedge.dmapi.system.v1.ModelNumber
	(*FirmwareInfo)(nil),  // 1: siemens.iedge.dmapi.system.v1.FirmwareInfo
	(*Resource)(nil),      // 2: siemens.iedge.dmapi.system.v1.Resource
	(*Cpu)(nil),           // 3: siemens.iedge.dmapi.system.v1.Cpu
	(*Stats)(nil),         // 4: siemens.iedge.dmapi.system.v1.Stats
	(*Limits)(nil),        // 5: siemens.iedge.dmapi.system.v1.Limits
	(*LogRequest)(nil),    // 6: siemens.iedge.dmapi.system.v1.LogRequest
	(*LogResponse)(nil),   // 7: siemens.iedge.dmapi.system.v1.LogResponse
	(*emptypb.Empty)(nil), // 8: google.protobuf.Empty
	(*anypb.Any)(nil),     // 9: google.protobuf.Any
}
var file_System_proto_depIdxs = []int32{
	3,  // 0: siemens.iedge.dmapi.system.v1.Stats.cpu:type_name -> siemens.iedge.dmapi.system.v1.Cpu
	2,  // 1: siemens.iedge.dmapi.system.v1.Stats.storageDevices:type_name -> siemens.iedge.dmapi.system.v1.Resource
	2,  // 2: siemens.iedge.dmapi.system.v1.Stats.memory:type_name -> siemens.iedge.dmapi.system.v1.Resource
	8,  // 3: siemens.iedge.dmapi.system.v1.SystemService.RestartDevice:input_type -> google.protobuf.Empty
	8,  // 4: siemens.iedge.dmapi.system.v1.SystemService.ShutdownDevice:input_type -> google.protobuf.Empty
	8,  // 5: siemens.iedge.dmapi.system.v1.SystemService.HardReset:input_type -> google.protobuf.Empty
	8,  // 6: siemens.iedge.dmapi.system.v1.SystemService.GetModelNumber:input_type -> google.protobuf.Empty
	8,  // 7: siemens.iedge.dmapi.system.v1.SystemService.GetFirmwareInfo:input_type -> google.protobuf.Empty
	8,  // 8: siemens.iedge.dmapi.system.v1.SystemService.GetResourceStats:input_type -> google.protobuf.Empty
	8,  // 9: siemens.iedge.dmapi.system.v1.SystemService.GetLimits:input_type -> google.protobuf.Empty
	8,  // 10: siemens.iedge.dmapi.system.v1.SystemService.GetCustomSettings:input_type -> google.protobuf.Empty
	9,  // 11: siemens.iedge.dmapi.system.v1.SystemService.ApplyCustomSettings:input_type -> google.protobuf.Any
	6,  // 12: siemens.iedge.dmapi.system.v1.SystemService.GetLogFile:input_type -> siemens.iedge.dmapi.system.v1.LogRequest
	8,  // 13: siemens.iedge.dmapi.system.v1.SystemService.RestartDevice:output_type -> google.protobuf.Empty
	8,  // 14: siemens.iedge.dmapi.system.v1.SystemService.ShutdownDevice:output_type -> google.protobuf.Empty
	8,  // 15: siemens.iedge.dmapi.system.v1.SystemService.HardReset:output_type -> google.protobuf.Empty
	0,  // 16: siemens.iedge.dmapi.system.v1.SystemService.GetModelNumber:output_type -> siemens.iedge.dmapi.system.v1.ModelNumber
	1,  // 17: siemens.iedge.dmapi.system.v1.SystemService.GetFirmwareInfo:output_type -> siemens.iedge.dmapi.system.v1.FirmwareInfo
	4,  // 18: siemens.iedge.dmapi.system.v1.SystemService.GetResourceStats:output_type -> siemens.iedge.dmapi.system.v1.Stats
	5,  // 19: siemens.iedge.dmapi.system.v1.SystemService.GetLimits:output_type -> siemens.iedge.dmapi.system.v1.Limits
	9,  // 20: siemens.iedge.dmapi.system.v1.SystemService.GetCustomSettings:output_type -> google.protobuf.Any
	8,  // 21: siemens.iedge.dmapi.system.v1.SystemService.ApplyCustomSettings:output_type -> google.protobuf.Empty
	7,  // 22: siemens.iedge.dmapi.system.v1.SystemService.GetLogFile:output_type -> siemens.iedge.dmapi.system.v1.LogResponse
	13, // [13:23] is the sub-list for method output_type
	3,  // [3:13] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_System_proto_init() }
func file_System_proto_init() {
	if File_System_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_System_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelNumber); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_System_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirmwareInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_System_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_System_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cpu); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_System_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_System_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Limits); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_System_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_System_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_System_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_System_proto_goTypes,
		DependencyIndexes: file_System_proto_depIdxs,
		MessageInfos:      file_System_proto_msgTypes,
	}.Build()
	File_System_proto = out.File
	file_System_proto_rawDesc = nil
	file_System_proto_goTypes = nil
	file_System_proto_depIdxs = nil
}
