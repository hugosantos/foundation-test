// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: universe/db/postgres/provider.proto

package postgres

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DatabaseArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client           string               `protobuf:"bytes,7,opt,name=client,proto3" json:"client,omitempty"`
	ResourceRef      string               `protobuf:"bytes,3,opt,name=resource_ref,json=resourceRef,proto3" json:"resource_ref,omitempty"`
	MaxConns         int32                `protobuf:"varint,4,opt,name=max_conns,json=maxConns,proto3" json:"max_conns,omitempty"` // Set if > 0.
	MaxConnsFromEnv  string               `protobuf:"bytes,6,opt,name=max_conns_from_env,json=maxConnsFromEnv,proto3" json:"max_conns_from_env,omitempty"`
	MaxConnsIdleTime *durationpb.Duration `protobuf:"bytes,5,opt,name=max_conns_idle_time,json=maxConnsIdleTime,proto3" json:"max_conns_idle_time,omitempty"`
}

func (x *DatabaseArgs) Reset() {
	*x = DatabaseArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_universe_db_postgres_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseArgs) ProtoMessage() {}

func (x *DatabaseArgs) ProtoReflect() protoreflect.Message {
	mi := &file_universe_db_postgres_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseArgs.ProtoReflect.Descriptor instead.
func (*DatabaseArgs) Descriptor() ([]byte, []int) {
	return file_universe_db_postgres_provider_proto_rawDescGZIP(), []int{0}
}

func (x *DatabaseArgs) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *DatabaseArgs) GetResourceRef() string {
	if x != nil {
		return x.ResourceRef
	}
	return ""
}

func (x *DatabaseArgs) GetMaxConns() int32 {
	if x != nil {
		return x.MaxConns
	}
	return 0
}

func (x *DatabaseArgs) GetMaxConnsFromEnv() string {
	if x != nil {
		return x.MaxConnsFromEnv
	}
	return ""
}

func (x *DatabaseArgs) GetMaxConnsIdleTime() *durationpb.Duration {
	if x != nil {
		return x.MaxConnsIdleTime
	}
	return nil
}

type DatabaseReferenceArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterRef string `protobuf:"bytes,1,opt,name=cluster_ref,json=clusterRef,proto3" json:"cluster_ref,omitempty"`
	Database   string `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
}

func (x *DatabaseReferenceArgs) Reset() {
	*x = DatabaseReferenceArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_universe_db_postgres_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseReferenceArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseReferenceArgs) ProtoMessage() {}

func (x *DatabaseReferenceArgs) ProtoReflect() protoreflect.Message {
	mi := &file_universe_db_postgres_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseReferenceArgs.ProtoReflect.Descriptor instead.
func (*DatabaseReferenceArgs) Descriptor() ([]byte, []int) {
	return file_universe_db_postgres_provider_proto_rawDescGZIP(), []int{1}
}

func (x *DatabaseReferenceArgs) GetClusterRef() string {
	if x != nil {
		return x.ClusterRef
	}
	return ""
}

func (x *DatabaseReferenceArgs) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

type FactoryArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client string `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *FactoryArgs) Reset() {
	*x = FactoryArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_universe_db_postgres_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FactoryArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FactoryArgs) ProtoMessage() {}

func (x *FactoryArgs) ProtoReflect() protoreflect.Message {
	mi := &file_universe_db_postgres_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FactoryArgs.ProtoReflect.Descriptor instead.
func (*FactoryArgs) Descriptor() ([]byte, []int) {
	return file_universe_db_postgres_provider_proto_rawDescGZIP(), []int{2}
}

func (x *FactoryArgs) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

var File_universe_db_postgres_provider_proto protoreflect.FileDescriptor

var file_universe_db_postgres_provider_proto_rawDesc = []byte{
	0x0a, 0x23, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f, 0x64, 0x62, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x64, 0x62, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x01, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x66, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x12,
	0x2b, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x73, 0x5f, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x65, 0x6e, 0x76, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x61, 0x78,
	0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x45, 0x6e, 0x76, 0x12, 0x48, 0x0a, 0x13,
	0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x73, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x49, 0x64,
	0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02,
	0x10, 0x03, 0x22, 0x54, 0x0a, 0x15, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x66, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x0b, 0x46, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x41, 0x72, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42,
	0x33, 0x5a, 0x31, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73,
	0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f, 0x64, 0x62, 0x2f, 0x70, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_universe_db_postgres_provider_proto_rawDescOnce sync.Once
	file_universe_db_postgres_provider_proto_rawDescData = file_universe_db_postgres_provider_proto_rawDesc
)

func file_universe_db_postgres_provider_proto_rawDescGZIP() []byte {
	file_universe_db_postgres_provider_proto_rawDescOnce.Do(func() {
		file_universe_db_postgres_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_universe_db_postgres_provider_proto_rawDescData)
	})
	return file_universe_db_postgres_provider_proto_rawDescData
}

var file_universe_db_postgres_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_universe_db_postgres_provider_proto_goTypes = []interface{}{
	(*DatabaseArgs)(nil),          // 0: foundation.universe.db.postgres.DatabaseArgs
	(*DatabaseReferenceArgs)(nil), // 1: foundation.universe.db.postgres.DatabaseReferenceArgs
	(*FactoryArgs)(nil),           // 2: foundation.universe.db.postgres.FactoryArgs
	(*durationpb.Duration)(nil),   // 3: google.protobuf.Duration
}
var file_universe_db_postgres_provider_proto_depIdxs = []int32{
	3, // 0: foundation.universe.db.postgres.DatabaseArgs.max_conns_idle_time:type_name -> google.protobuf.Duration
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_universe_db_postgres_provider_proto_init() }
func file_universe_db_postgres_provider_proto_init() {
	if File_universe_db_postgres_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_universe_db_postgres_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseArgs); i {
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
		file_universe_db_postgres_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseReferenceArgs); i {
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
		file_universe_db_postgres_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FactoryArgs); i {
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
			RawDescriptor: file_universe_db_postgres_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_universe_db_postgres_provider_proto_goTypes,
		DependencyIndexes: file_universe_db_postgres_provider_proto_depIdxs,
		MessageInfos:      file_universe_db_postgres_provider_proto_msgTypes,
	}.Build()
	File_universe_db_postgres_provider_proto = out.File
	file_universe_db_postgres_provider_proto_rawDesc = nil
	file_universe_db_postgres_provider_proto_goTypes = nil
	file_universe_db_postgres_provider_proto_depIdxs = nil
}
