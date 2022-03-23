// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: Sdfsnode.proto

package sdfs_creator

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DeploymentType int32

const (
	DeploymentType_GKE   DeploymentType = 0
	DeploymentType_LOCAL DeploymentType = 1
)

// Enum value maps for DeploymentType.
var (
	DeploymentType_name = map[int32]string{
		0: "GKE",
		1: "LOCAL",
	}
	DeploymentType_value = map[string]int32{
		"GKE":   0,
		"LOCAL": 1,
	}
)

func (x DeploymentType) Enum() *DeploymentType {
	p := new(DeploymentType)
	*p = x
	return p
}

func (x DeploymentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeploymentType) Descriptor() protoreflect.EnumDescriptor {
	return file_Sdfsnode_proto_enumTypes[0].Descriptor()
}

func (DeploymentType) Type() protoreflect.EnumType {
	return &file_Sdfsnode_proto_enumTypes[0]
}

func (x DeploymentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeploymentType.Descriptor instead.
func (DeploymentType) EnumDescriptor() ([]byte, []int) {
	return file_Sdfsnode_proto_rawDescGZIP(), []int{0}
}

type StorageMount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MountPath string `protobuf:"bytes,1,opt,name=mountPath,proto3" json:"mountPath,omitempty"`
	Dev       string `protobuf:"bytes,2,opt,name=dev,proto3" json:"dev,omitempty"`
	Total     int64  `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Used      int64  `protobuf:"varint,4,opt,name=used,proto3" json:"used,omitempty"`
	Avail     int64  `protobuf:"varint,5,opt,name=avail,proto3" json:"avail,omitempty"`
	FsType    string `protobuf:"bytes,6,opt,name=fsType,proto3" json:"fsType,omitempty"`
}

func (x *StorageMount) Reset() {
	*x = StorageMount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Sdfsnode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageMount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageMount) ProtoMessage() {}

func (x *StorageMount) ProtoReflect() protoreflect.Message {
	mi := &file_Sdfsnode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageMount.ProtoReflect.Descriptor instead.
func (*StorageMount) Descriptor() ([]byte, []int) {
	return file_Sdfsnode_proto_rawDescGZIP(), []int{0}
}

func (x *StorageMount) GetMountPath() string {
	if x != nil {
		return x.MountPath
	}
	return ""
}

func (x *StorageMount) GetDev() string {
	if x != nil {
		return x.Dev
	}
	return ""
}

func (x *StorageMount) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *StorageMount) GetUsed() int64 {
	if x != nil {
		return x.Used
	}
	return 0
}

func (x *StorageMount) GetAvail() int64 {
	if x != nil {
		return x.Avail
	}
	return 0
}

func (x *StorageMount) GetFsType() string {
	if x != nil {
		return x.FsType
	}
	return ""
}

type StorageHost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostName       string         `protobuf:"bytes,1,opt,name=hostName,proto3" json:"hostName,omitempty"`
	UserName       string         `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Password       string         `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	PubKey         string         `protobuf:"bytes,4,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	PrivKey        string         `protobuf:"bytes,5,opt,name=privKey,proto3" json:"privKey,omitempty"`
	Description    string         `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	BaseUrl        string         `protobuf:"bytes,7,opt,name=baseUrl,proto3" json:"baseUrl,omitempty"`
	Id             string         `protobuf:"bytes,9,opt,name=id,proto3" json:"id,omitempty"`
	Standalone     bool           `protobuf:"varint,10,opt,name=standalone,proto3" json:"standalone,omitempty"`
	Connected      bool           `protobuf:"varint,11,opt,name=connected,proto3" json:"connected,omitempty"`
	DeploymentType DeploymentType `protobuf:"varint,12,opt,name=deploymentType,proto3,enum=DeploymentType" json:"deploymentType,omitempty"`
	Running        bool           `protobuf:"varint,13,opt,name=running,proto3" json:"running,omitempty"`
	Tls            bool           `protobuf:"varint,14,opt,name=tls,proto3" json:"tls,omitempty"`
	Mtls           bool           `protobuf:"varint,15,opt,name=mtls,proto3" json:"mtls,omitempty"`
}

func (x *StorageHost) Reset() {
	*x = StorageHost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Sdfsnode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageHost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageHost) ProtoMessage() {}

func (x *StorageHost) ProtoReflect() protoreflect.Message {
	mi := &file_Sdfsnode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageHost.ProtoReflect.Descriptor instead.
func (*StorageHost) Descriptor() ([]byte, []int) {
	return file_Sdfsnode_proto_rawDescGZIP(), []int{1}
}

func (x *StorageHost) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *StorageHost) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *StorageHost) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *StorageHost) GetPubKey() string {
	if x != nil {
		return x.PubKey
	}
	return ""
}

func (x *StorageHost) GetPrivKey() string {
	if x != nil {
		return x.PrivKey
	}
	return ""
}

func (x *StorageHost) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *StorageHost) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *StorageHost) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StorageHost) GetStandalone() bool {
	if x != nil {
		return x.Standalone
	}
	return false
}

func (x *StorageHost) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

func (x *StorageHost) GetDeploymentType() DeploymentType {
	if x != nil {
		return x.DeploymentType
	}
	return DeploymentType_GKE
}

func (x *StorageHost) GetRunning() bool {
	if x != nil {
		return x.Running
	}
	return false
}

func (x *StorageHost) GetTls() bool {
	if x != nil {
		return x.Tls
	}
	return false
}

func (x *StorageHost) GetMtls() bool {
	if x != nil {
		return x.Mtls
	}
	return false
}

type StorageLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MetaData  string `protobuf:"bytes,2,opt,name=metaData,proto3" json:"metaData,omitempty"`
	ChunkData string `protobuf:"bytes,3,opt,name=chunkData,proto3" json:"chunkData,omitempty"`
	HashDB    string `protobuf:"bytes,4,opt,name=hashDB,proto3" json:"hashDB,omitempty"`
}

func (x *StorageLocation) Reset() {
	*x = StorageLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Sdfsnode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageLocation) ProtoMessage() {}

func (x *StorageLocation) ProtoReflect() protoreflect.Message {
	mi := &file_Sdfsnode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageLocation.ProtoReflect.Descriptor instead.
func (*StorageLocation) Descriptor() ([]byte, []int) {
	return file_Sdfsnode_proto_rawDescGZIP(), []int{2}
}

func (x *StorageLocation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StorageLocation) GetMetaData() string {
	if x != nil {
		return x.MetaData
	}
	return ""
}

func (x *StorageLocation) GetChunkData() string {
	if x != nil {
		return x.ChunkData
	}
	return ""
}

func (x *StorageLocation) GetHashDB() string {
	if x != nil {
		return x.HashDB
	}
	return ""
}

type SystemInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCPU    float64 `protobuf:"fixed64,1,opt,name=totalCPU,proto3" json:"totalCPU,omitempty"`
	Cores       int32   `protobuf:"varint,2,opt,name=cores,proto3" json:"cores,omitempty"`
	TotalMemory int64   `protobuf:"varint,3,opt,name=totalMemory,proto3" json:"totalMemory,omitempty"`
	FreeMemory  int64   `protobuf:"varint,4,opt,name=freeMemory,proto3" json:"freeMemory,omitempty"`
	Version     string  `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Os          string  `protobuf:"bytes,6,opt,name=os,proto3" json:"os,omitempty"`
	Uptime      string  `protobuf:"bytes,7,opt,name=uptime,proto3" json:"uptime,omitempty"`
}

func (x *SystemInfo) Reset() {
	*x = SystemInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Sdfsnode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemInfo) ProtoMessage() {}

func (x *SystemInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Sdfsnode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemInfo.ProtoReflect.Descriptor instead.
func (*SystemInfo) Descriptor() ([]byte, []int) {
	return file_Sdfsnode_proto_rawDescGZIP(), []int{3}
}

func (x *SystemInfo) GetTotalCPU() float64 {
	if x != nil {
		return x.TotalCPU
	}
	return 0
}

func (x *SystemInfo) GetCores() int32 {
	if x != nil {
		return x.Cores
	}
	return 0
}

func (x *SystemInfo) GetTotalMemory() int64 {
	if x != nil {
		return x.TotalMemory
	}
	return 0
}

func (x *SystemInfo) GetFreeMemory() int64 {
	if x != nil {
		return x.FreeMemory
	}
	return 0
}

func (x *SystemInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *SystemInfo) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *SystemInfo) GetUptime() string {
	if x != nil {
		return x.Uptime
	}
	return ""
}

var File_Sdfsnode_proto protoreflect.FileDescriptor

var file_Sdfsnode_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x64, 0x66, 0x73, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x96, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x65, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x65,
	0x76, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x73, 0x54, 0x79, 0x70, 0x65, 0x22, 0x96, 0x03, 0x0a, 0x0b, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x4b, 0x65, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x4b, 0x65, 0x79, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x6c, 0x6f, 0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x6c, 0x6f, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x0e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x6c, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x74, 0x6c, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6d, 0x74,
	0x6c, 0x73, 0x22, 0x73, 0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x44, 0x42, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x68, 0x61, 0x73, 0x68, 0x44, 0x42, 0x22, 0xc2, 0x01, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x50, 0x55, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x50, 0x55, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x72,
	0x65, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x66, 0x72, 0x65, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2a, 0x24, 0x0a, 0x0e,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07,
	0x0a, 0x03, 0x47, 0x4b, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x43, 0x41, 0x4c,
	0x10, 0x01, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x75, 0x70, 0x2f, 0x73, 0x64, 0x66, 0x73, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Sdfsnode_proto_rawDescOnce sync.Once
	file_Sdfsnode_proto_rawDescData = file_Sdfsnode_proto_rawDesc
)

func file_Sdfsnode_proto_rawDescGZIP() []byte {
	file_Sdfsnode_proto_rawDescOnce.Do(func() {
		file_Sdfsnode_proto_rawDescData = protoimpl.X.CompressGZIP(file_Sdfsnode_proto_rawDescData)
	})
	return file_Sdfsnode_proto_rawDescData
}

var file_Sdfsnode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Sdfsnode_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Sdfsnode_proto_goTypes = []interface{}{
	(DeploymentType)(0),     // 0: deploymentType
	(*StorageMount)(nil),    // 1: StorageMount
	(*StorageHost)(nil),     // 2: StorageHost
	(*StorageLocation)(nil), // 3: storageLocation
	(*SystemInfo)(nil),      // 4: SystemInfo
}
var file_Sdfsnode_proto_depIdxs = []int32{
	0, // 0: StorageHost.deploymentType:type_name -> deploymentType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Sdfsnode_proto_init() }
func file_Sdfsnode_proto_init() {
	if File_Sdfsnode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Sdfsnode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageMount); i {
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
		file_Sdfsnode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageHost); i {
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
		file_Sdfsnode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageLocation); i {
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
		file_Sdfsnode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemInfo); i {
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
			RawDescriptor: file_Sdfsnode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Sdfsnode_proto_goTypes,
		DependencyIndexes: file_Sdfsnode_proto_depIdxs,
		EnumInfos:         file_Sdfsnode_proto_enumTypes,
		MessageInfos:      file_Sdfsnode_proto_msgTypes,
	}.Build()
	File_Sdfsnode_proto = out.File
	file_Sdfsnode_proto_rawDesc = nil
	file_Sdfsnode_proto_goTypes = nil
	file_Sdfsnode_proto_depIdxs = nil
}
