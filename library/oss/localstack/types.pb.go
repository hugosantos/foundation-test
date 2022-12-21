// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: library/oss/localstack/types.proto

package localstack

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	schema "namespacelabs.dev/foundation/schema"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClusterIntent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterIntent) Reset() {
	*x = ClusterIntent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_oss_localstack_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterIntent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterIntent) ProtoMessage() {}

func (x *ClusterIntent) ProtoReflect() protoreflect.Message {
	mi := &file_library_oss_localstack_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterIntent.ProtoReflect.Descriptor instead.
func (*ClusterIntent) Descriptor() ([]byte, []int) {
	return file_library_oss_localstack_types_proto_rawDescGZIP(), []int{0}
}

type ClusterInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *ClusterInstance) Reset() {
	*x = ClusterInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_oss_localstack_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterInstance) ProtoMessage() {}

func (x *ClusterInstance) ProtoReflect() protoreflect.Message {
	mi := &file_library_oss_localstack_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterInstance.ProtoReflect.Descriptor instead.
func (*ClusterInstance) Descriptor() ([]byte, []int) {
	return file_library_oss_localstack_types_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterInstance) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type ServerIntent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server *schema.PackageRef `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *ServerIntent) Reset() {
	*x = ServerIntent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_oss_localstack_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerIntent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerIntent) ProtoMessage() {}

func (x *ServerIntent) ProtoReflect() protoreflect.Message {
	mi := &file_library_oss_localstack_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerIntent.ProtoReflect.Descriptor instead.
func (*ServerIntent) Descriptor() ([]byte, []int) {
	return file_library_oss_localstack_types_proto_rawDescGZIP(), []int{2}
}

func (x *ServerIntent) GetServer() *schema.PackageRef {
	if x != nil {
		return x.Server
	}
	return nil
}

var File_library_oss_localstack_types_proto protoreflect.FileDescriptor

var file_library_oss_localstack_types_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6f, 0x73, 0x73, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x6f, 0x73,
	0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x1a, 0x14, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x0f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x22, 0x45, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x66, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x35, 0x5a, 0x33, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2f, 0x6f, 0x73, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_library_oss_localstack_types_proto_rawDescOnce sync.Once
	file_library_oss_localstack_types_proto_rawDescData = file_library_oss_localstack_types_proto_rawDesc
)

func file_library_oss_localstack_types_proto_rawDescGZIP() []byte {
	file_library_oss_localstack_types_proto_rawDescOnce.Do(func() {
		file_library_oss_localstack_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_library_oss_localstack_types_proto_rawDescData)
	})
	return file_library_oss_localstack_types_proto_rawDescData
}

var file_library_oss_localstack_types_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_library_oss_localstack_types_proto_goTypes = []interface{}{
	(*ClusterIntent)(nil),     // 0: library.oss.localstack.ClusterIntent
	(*ClusterInstance)(nil),   // 1: library.oss.localstack.ClusterInstance
	(*ServerIntent)(nil),      // 2: library.oss.localstack.ServerIntent
	(*schema.PackageRef)(nil), // 3: foundation.schema.PackageRef
}
var file_library_oss_localstack_types_proto_depIdxs = []int32{
	3, // 0: library.oss.localstack.ServerIntent.server:type_name -> foundation.schema.PackageRef
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_library_oss_localstack_types_proto_init() }
func file_library_oss_localstack_types_proto_init() {
	if File_library_oss_localstack_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_library_oss_localstack_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterIntent); i {
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
		file_library_oss_localstack_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterInstance); i {
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
		file_library_oss_localstack_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerIntent); i {
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
			RawDescriptor: file_library_oss_localstack_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_library_oss_localstack_types_proto_goTypes,
		DependencyIndexes: file_library_oss_localstack_types_proto_depIdxs,
		MessageInfos:      file_library_oss_localstack_types_proto_msgTypes,
	}.Build()
	File_library_oss_localstack_types_proto = out.File
	file_library_oss_localstack_types_proto_rawDesc = nil
	file_library_oss_localstack_types_proto_goTypes = nil
	file_library_oss_localstack_types_proto_depIdxs = nil
}