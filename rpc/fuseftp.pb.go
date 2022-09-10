// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.5
// source: rpc/fuseftp.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type VersionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Semver string `protobuf:"bytes,1,opt,name=semver,proto3" json:"semver,omitempty"`
}

func (x *VersionInfo) Reset() {
	*x = VersionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fuseftp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionInfo) ProtoMessage() {}

func (x *VersionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fuseftp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionInfo.ProtoReflect.Descriptor instead.
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return file_rpc_fuseftp_proto_rawDescGZIP(), []int{0}
}

func (x *VersionInfo) GetSemver() string {
	if x != nil {
		return x.Semver
	}
	return ""
}

type AddressAndPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   []byte `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *AddressAndPort) Reset() {
	*x = AddressAndPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fuseftp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressAndPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressAndPort) ProtoMessage() {}

func (x *AddressAndPort) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fuseftp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressAndPort.ProtoReflect.Descriptor instead.
func (*AddressAndPort) Descriptor() ([]byte, []int) {
	return file_rpc_fuseftp_proto_rawDescGZIP(), []int{1}
}

func (x *AddressAndPort) GetIp() []byte {
	if x != nil {
		return x.Ip
	}
	return nil
}

func (x *AddressAndPort) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type MountIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MountIdentifier) Reset() {
	*x = MountIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fuseftp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountIdentifier) ProtoMessage() {}

func (x *MountIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fuseftp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountIdentifier.ProtoReflect.Descriptor instead.
func (*MountIdentifier) Descriptor() ([]byte, []int) {
	return file_rpc_fuseftp_proto_rawDescGZIP(), []int{2}
}

func (x *MountIdentifier) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SetFtpServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *MountIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FtpServer *AddressAndPort  `protobuf:"bytes,2,opt,name=ftp_server,json=ftpServer,proto3" json:"ftp_server,omitempty"`
}

func (x *SetFtpServerRequest) Reset() {
	*x = SetFtpServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fuseftp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetFtpServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetFtpServerRequest) ProtoMessage() {}

func (x *SetFtpServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fuseftp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetFtpServerRequest.ProtoReflect.Descriptor instead.
func (*SetFtpServerRequest) Descriptor() ([]byte, []int) {
	return file_rpc_fuseftp_proto_rawDescGZIP(), []int{3}
}

func (x *SetFtpServerRequest) GetId() *MountIdentifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *SetFtpServerRequest) GetFtpServer() *AddressAndPort {
	if x != nil {
		return x.FtpServer
	}
	return nil
}

type MountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mount point on the local computer. Must be a drive letter on windows
	MountPoint string `protobuf:"bytes,1,opt,name=mount_point,json=mountPoint,proto3" json:"mount_point,omitempty"`
	// The ftp_server to connect to
	FtpServer *AddressAndPort `protobuf:"bytes,2,opt,name=ftp_server,json=ftpServer,proto3" json:"ftp_server,omitempty"`
	// Read timout
	ReadTimeout *durationpb.Duration `protobuf:"bytes,3,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	// The directory on the FTP server that gets mounted
	Directory string `protobuf:"bytes,4,opt,name=directory,proto3" json:"directory,omitempty"`
}

func (x *MountRequest) Reset() {
	*x = MountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fuseftp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountRequest) ProtoMessage() {}

func (x *MountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fuseftp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountRequest.ProtoReflect.Descriptor instead.
func (*MountRequest) Descriptor() ([]byte, []int) {
	return file_rpc_fuseftp_proto_rawDescGZIP(), []int{4}
}

func (x *MountRequest) GetMountPoint() string {
	if x != nil {
		return x.MountPoint
	}
	return ""
}

func (x *MountRequest) GetFtpServer() *AddressAndPort {
	if x != nil {
		return x.FtpServer
	}
	return nil
}

func (x *MountRequest) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *MountRequest) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

var File_rpc_fuseftp_proto protoreflect.FileDescriptor

var file_rpc_fuseftp_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x75, 0x73, 0x65, 0x66, 0x74, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x64, 0x61, 0x74, 0x61, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x66, 0x75,
	0x73, 0x65, 0x66, 0x74, 0x70, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x41, 0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22,
	0x21, 0x0a, 0x0f, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x46, 0x74, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x77, 0x69, 0x72,
	0x65, 0x2e, 0x66, 0x75, 0x73, 0x65, 0x66, 0x74, 0x70, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3f, 0x0a,
	0x0a, 0x66, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x66, 0x75, 0x73,
	0x65, 0x66, 0x74, 0x70, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x41, 0x6e, 0x64, 0x50,
	0x6f, 0x72, 0x74, 0x52, 0x09, 0x66, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0xcc,
	0x01, 0x0a, 0x0c, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x3f, 0x0a, 0x0a, 0x66, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x77, 0x69, 0x72, 0x65, 0x2e,
	0x66, 0x75, 0x73, 0x65, 0x66, 0x74, 0x70, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x41,
	0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x09, 0x66, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x32, 0xac, 0x02,
	0x0a, 0x07, 0x46, 0x75, 0x73, 0x65, 0x46, 0x54, 0x50, 0x12, 0x40, 0x0a, 0x07, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x66, 0x75, 0x73, 0x65, 0x66, 0x74, 0x70, 0x2e,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4a, 0x0a, 0x05, 0x4d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x77, 0x69, 0x72, 0x65, 0x2e,
	0x66, 0x75, 0x73, 0x65, 0x66, 0x74, 0x70, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x77, 0x69, 0x72, 0x65, 0x2e,
	0x66, 0x75, 0x73, 0x65, 0x66, 0x74, 0x70, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x07, 0x55, 0x6e, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x66, 0x75,
	0x73, 0x65, 0x66, 0x74, 0x70, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4d, 0x0a,
	0x0c, 0x53, 0x65, 0x74, 0x46, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x25, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x66, 0x75, 0x73, 0x65, 0x66, 0x74, 0x70,
	0x2e, 0x53, 0x65, 0x74, 0x46, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x24, 0x5a, 0x22,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x77,
	0x69, 0x72, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x66, 0x75, 0x73, 0x65, 0x66, 0x74, 0x70, 0x2f, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_fuseftp_proto_rawDescOnce sync.Once
	file_rpc_fuseftp_proto_rawDescData = file_rpc_fuseftp_proto_rawDesc
)

func file_rpc_fuseftp_proto_rawDescGZIP() []byte {
	file_rpc_fuseftp_proto_rawDescOnce.Do(func() {
		file_rpc_fuseftp_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_fuseftp_proto_rawDescData)
	})
	return file_rpc_fuseftp_proto_rawDescData
}

var file_rpc_fuseftp_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpc_fuseftp_proto_goTypes = []interface{}{
	(*VersionInfo)(nil),         // 0: datawire.fuseftp.VersionInfo
	(*AddressAndPort)(nil),      // 1: datawire.fuseftp.AddressAndPort
	(*MountIdentifier)(nil),     // 2: datawire.fuseftp.MountIdentifier
	(*SetFtpServerRequest)(nil), // 3: datawire.fuseftp.SetFtpServerRequest
	(*MountRequest)(nil),        // 4: datawire.fuseftp.MountRequest
	(*durationpb.Duration)(nil), // 5: google.protobuf.Duration
	(*emptypb.Empty)(nil),       // 6: google.protobuf.Empty
}
var file_rpc_fuseftp_proto_depIdxs = []int32{
	2, // 0: datawire.fuseftp.SetFtpServerRequest.id:type_name -> datawire.fuseftp.MountIdentifier
	1, // 1: datawire.fuseftp.SetFtpServerRequest.ftp_server:type_name -> datawire.fuseftp.AddressAndPort
	1, // 2: datawire.fuseftp.MountRequest.ftp_server:type_name -> datawire.fuseftp.AddressAndPort
	5, // 3: datawire.fuseftp.MountRequest.read_timeout:type_name -> google.protobuf.Duration
	6, // 4: datawire.fuseftp.FuseFTP.Version:input_type -> google.protobuf.Empty
	4, // 5: datawire.fuseftp.FuseFTP.Mount:input_type -> datawire.fuseftp.MountRequest
	2, // 6: datawire.fuseftp.FuseFTP.Unmount:input_type -> datawire.fuseftp.MountIdentifier
	3, // 7: datawire.fuseftp.FuseFTP.SetFtpServer:input_type -> datawire.fuseftp.SetFtpServerRequest
	0, // 8: datawire.fuseftp.FuseFTP.Version:output_type -> datawire.fuseftp.VersionInfo
	2, // 9: datawire.fuseftp.FuseFTP.Mount:output_type -> datawire.fuseftp.MountIdentifier
	6, // 10: datawire.fuseftp.FuseFTP.Unmount:output_type -> google.protobuf.Empty
	6, // 11: datawire.fuseftp.FuseFTP.SetFtpServer:output_type -> google.protobuf.Empty
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_fuseftp_proto_init() }
func file_rpc_fuseftp_proto_init() {
	if File_rpc_fuseftp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_fuseftp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionInfo); i {
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
		file_rpc_fuseftp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressAndPort); i {
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
		file_rpc_fuseftp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountIdentifier); i {
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
		file_rpc_fuseftp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetFtpServerRequest); i {
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
		file_rpc_fuseftp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountRequest); i {
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
			RawDescriptor: file_rpc_fuseftp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_fuseftp_proto_goTypes,
		DependencyIndexes: file_rpc_fuseftp_proto_depIdxs,
		MessageInfos:      file_rpc_fuseftp_proto_msgTypes,
	}.Build()
	File_rpc_fuseftp_proto = out.File
	file_rpc_fuseftp_proto_rawDesc = nil
	file_rpc_fuseftp_proto_goTypes = nil
	file_rpc_fuseftp_proto_depIdxs = nil
}
