// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: internal/build/registry/types.proto

package registry

import (
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

type Registry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url              string             `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Insecure         bool               `protobuf:"varint,2,opt,name=insecure,proto3" json:"insecure,omitempty"`
	UseDockerAuth    bool               `protobuf:"varint,3,opt,name=use_docker_auth,json=useDockerAuth,proto3" json:"use_docker_auth,omitempty"`        // If true, the credentials stored by Docker are used to access this repository.
	SingleRepository bool               `protobuf:"varint,4,opt,name=single_repository,json=singleRepository,proto3" json:"single_repository,omitempty"` // If true, all images are stored in a single repository, rather than creating a repository per image.
	Transport        *RegistryTransport `protobuf:"bytes,5,opt,name=transport,proto3" json:"transport,omitempty"`
}

func (x *Registry) Reset() {
	*x = Registry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_build_registry_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_internal_build_registry_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry.ProtoReflect.Descriptor instead.
func (*Registry) Descriptor() ([]byte, []int) {
	return file_internal_build_registry_types_proto_rawDescGZIP(), []int{0}
}

func (x *Registry) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Registry) GetInsecure() bool {
	if x != nil {
		return x.Insecure
	}
	return false
}

func (x *Registry) GetUseDockerAuth() bool {
	if x != nil {
		return x.UseDockerAuth
	}
	return false
}

func (x *Registry) GetSingleRepository() bool {
	if x != nil {
		return x.SingleRepository
	}
	return false
}

func (x *Registry) GetTransport() *RegistryTransport {
	if x != nil {
		return x.Transport
	}
	return nil
}

type Provider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *Provider) Reset() {
	*x = Provider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_build_registry_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provider) ProtoMessage() {}

func (x *Provider) ProtoReflect() protoreflect.Message {
	mi := &file_internal_build_registry_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provider.ProtoReflect.Descriptor instead.
func (*Provider) Descriptor() ([]byte, []int) {
	return file_internal_build_registry_types_proto_rawDescGZIP(), []int{1}
}

func (x *Provider) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

type RegistryTransport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ssh *RegistryTransport_SSH `protobuf:"bytes,1,opt,name=ssh,proto3" json:"ssh,omitempty"`
}

func (x *RegistryTransport) Reset() {
	*x = RegistryTransport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_build_registry_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryTransport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryTransport) ProtoMessage() {}

func (x *RegistryTransport) ProtoReflect() protoreflect.Message {
	mi := &file_internal_build_registry_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryTransport.ProtoReflect.Descriptor instead.
func (*RegistryTransport) Descriptor() ([]byte, []int) {
	return file_internal_build_registry_types_proto_rawDescGZIP(), []int{2}
}

func (x *RegistryTransport) GetSsh() *RegistryTransport_SSH {
	if x != nil {
		return x.Ssh
	}
	return nil
}

type RegistryTransport_SSH struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User           string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	PrivateKeyPath string `protobuf:"bytes,2,opt,name=private_key_path,json=privateKeyPath,proto3" json:"private_key_path,omitempty"`
	SshAddr        string `protobuf:"bytes,3,opt,name=ssh_addr,json=sshAddr,proto3" json:"ssh_addr,omitempty"`
	RemoteAddr     string `protobuf:"bytes,4,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"`
	AgentSockPath  string `protobuf:"bytes,5,opt,name=agent_sock_path,json=agentSockPath,proto3" json:"agent_sock_path,omitempty"`
}

func (x *RegistryTransport_SSH) Reset() {
	*x = RegistryTransport_SSH{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_build_registry_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryTransport_SSH) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryTransport_SSH) ProtoMessage() {}

func (x *RegistryTransport_SSH) ProtoReflect() protoreflect.Message {
	mi := &file_internal_build_registry_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryTransport_SSH.ProtoReflect.Descriptor instead.
func (*RegistryTransport_SSH) Descriptor() ([]byte, []int) {
	return file_internal_build_registry_types_proto_rawDescGZIP(), []int{2, 0}
}

func (x *RegistryTransport_SSH) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *RegistryTransport_SSH) GetPrivateKeyPath() string {
	if x != nil {
		return x.PrivateKeyPath
	}
	return ""
}

func (x *RegistryTransport_SSH) GetSshAddr() string {
	if x != nil {
		return x.SshAddr
	}
	return ""
}

func (x *RegistryTransport_SSH) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *RegistryTransport_SSH) GetAgentSockPath() string {
	if x != nil {
		return x.AgentSockPath
	}
	return ""
}

var File_internal_build_registry_types_proto protoreflect.FileDescriptor

var file_internal_build_registry_types_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x22, 0xd9, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x75,
	0x73, 0x65, 0x5f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x4a, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x26, 0x0a, 0x08,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x22, 0x81, 0x02, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x42, 0x0a, 0x03, 0x73, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x53, 0x53, 0x48, 0x52, 0x03, 0x73, 0x73, 0x68, 0x1a, 0xa7,
	0x01, 0x0a, 0x03, 0x53, 0x53, 0x48, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x73, 0x68, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x73, 0x68, 0x41, 0x64, 0x64, 0x72, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72,
	0x12, 0x26, 0x0a, 0x0f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x53, 0x6f, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x42, 0x36, 0x5a, 0x34, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_build_registry_types_proto_rawDescOnce sync.Once
	file_internal_build_registry_types_proto_rawDescData = file_internal_build_registry_types_proto_rawDesc
)

func file_internal_build_registry_types_proto_rawDescGZIP() []byte {
	file_internal_build_registry_types_proto_rawDescOnce.Do(func() {
		file_internal_build_registry_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_build_registry_types_proto_rawDescData)
	})
	return file_internal_build_registry_types_proto_rawDescData
}

var file_internal_build_registry_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_build_registry_types_proto_goTypes = []interface{}{
	(*Registry)(nil),              // 0: foundation.build.registry.Registry
	(*Provider)(nil),              // 1: foundation.build.registry.Provider
	(*RegistryTransport)(nil),     // 2: foundation.build.registry.RegistryTransport
	(*RegistryTransport_SSH)(nil), // 3: foundation.build.registry.RegistryTransport.SSH
}
var file_internal_build_registry_types_proto_depIdxs = []int32{
	2, // 0: foundation.build.registry.Registry.transport:type_name -> foundation.build.registry.RegistryTransport
	3, // 1: foundation.build.registry.RegistryTransport.ssh:type_name -> foundation.build.registry.RegistryTransport.SSH
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_build_registry_types_proto_init() }
func file_internal_build_registry_types_proto_init() {
	if File_internal_build_registry_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_build_registry_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry); i {
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
		file_internal_build_registry_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provider); i {
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
		file_internal_build_registry_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryTransport); i {
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
		file_internal_build_registry_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryTransport_SSH); i {
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
			RawDescriptor: file_internal_build_registry_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_build_registry_types_proto_goTypes,
		DependencyIndexes: file_internal_build_registry_types_proto_depIdxs,
		MessageInfos:      file_internal_build_registry_types_proto_msgTypes,
	}.Build()
	File_internal_build_registry_types_proto = out.File
	file_internal_build_registry_types_proto_rawDesc = nil
	file_internal_build_registry_types_proto_goTypes = nil
	file_internal_build_registry_types_proto_depIdxs = nil
}
