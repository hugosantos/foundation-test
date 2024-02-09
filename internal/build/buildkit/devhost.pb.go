// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: internal/build/buildkit/devhost.proto

package buildkit

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

// Used for configuration purposes.
type Overrides struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerName         string                 `protobuf:"bytes,1,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	BuildkitAddr          string                 `protobuf:"bytes,2,opt,name=buildkit_addr,json=buildkitAddr,proto3" json:"buildkit_addr,omitempty"`
	HostedBuildCluster    *HostedBuildCluster    `protobuf:"bytes,3,opt,name=hosted_build_cluster,json=hostedBuildCluster,proto3" json:"hosted_build_cluster,omitempty"`
	ColocatedBuildCluster *ColocatedBuildCluster `protobuf:"bytes,4,opt,name=colocated_build_cluster,json=colocatedBuildCluster,proto3" json:"colocated_build_cluster,omitempty"`
}

func (x *Overrides) Reset() {
	*x = Overrides{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_build_buildkit_devhost_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Overrides) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Overrides) ProtoMessage() {}

func (x *Overrides) ProtoReflect() protoreflect.Message {
	mi := &file_internal_build_buildkit_devhost_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Overrides.ProtoReflect.Descriptor instead.
func (*Overrides) Descriptor() ([]byte, []int) {
	return file_internal_build_buildkit_devhost_proto_rawDescGZIP(), []int{0}
}

func (x *Overrides) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *Overrides) GetBuildkitAddr() string {
	if x != nil {
		return x.BuildkitAddr
	}
	return ""
}

func (x *Overrides) GetHostedBuildCluster() *HostedBuildCluster {
	if x != nil {
		return x.HostedBuildCluster
	}
	return nil
}

func (x *Overrides) GetColocatedBuildCluster() *ColocatedBuildCluster {
	if x != nil {
		return x.ColocatedBuildCluster
	}
	return nil
}

type HostedBuildCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId  string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	TargetPort int32  `protobuf:"varint,2,opt,name=target_port,json=targetPort,proto3" json:"target_port,omitempty"`
	Endpoint   string `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *HostedBuildCluster) Reset() {
	*x = HostedBuildCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_build_buildkit_devhost_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostedBuildCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostedBuildCluster) ProtoMessage() {}

func (x *HostedBuildCluster) ProtoReflect() protoreflect.Message {
	mi := &file_internal_build_buildkit_devhost_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostedBuildCluster.ProtoReflect.Descriptor instead.
func (*HostedBuildCluster) Descriptor() ([]byte, []int) {
	return file_internal_build_buildkit_devhost_proto_rawDescGZIP(), []int{1}
}

func (x *HostedBuildCluster) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *HostedBuildCluster) GetTargetPort() int32 {
	if x != nil {
		return x.TargetPort
	}
	return 0
}

func (x *HostedBuildCluster) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type ColocatedBuildCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace         string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	MatchingPodLabels map[string]string `protobuf:"bytes,2,rep,name=matching_pod_labels,json=matchingPodLabels,proto3" json:"matching_pod_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TargetPort        int32             `protobuf:"varint,3,opt,name=target_port,json=targetPort,proto3" json:"target_port,omitempty"`
}

func (x *ColocatedBuildCluster) Reset() {
	*x = ColocatedBuildCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_build_buildkit_devhost_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColocatedBuildCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColocatedBuildCluster) ProtoMessage() {}

func (x *ColocatedBuildCluster) ProtoReflect() protoreflect.Message {
	mi := &file_internal_build_buildkit_devhost_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColocatedBuildCluster.ProtoReflect.Descriptor instead.
func (*ColocatedBuildCluster) Descriptor() ([]byte, []int) {
	return file_internal_build_buildkit_devhost_proto_rawDescGZIP(), []int{2}
}

func (x *ColocatedBuildCluster) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ColocatedBuildCluster) GetMatchingPodLabels() map[string]string {
	if x != nil {
		return x.MatchingPodLabels
	}
	return nil
}

func (x *ColocatedBuildCluster) GetTargetPort() int32 {
	if x != nil {
		return x.TargetPort
	}
	return 0
}

var File_internal_build_buildkit_devhost_proto protoreflect.FileDescriptor

var file_internal_build_buildkit_devhost_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6b, 0x69, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x68, 0x6f, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6b,
	0x69, 0x74, 0x22, 0xa2, 0x02, 0x0a, 0x09, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x6b, 0x69, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x6b, 0x69, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x5f, 0x0a, 0x14,
	0x68, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x6b, 0x69, 0x74, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x12, 0x68, 0x6f, 0x73, 0x74, 0x65,
	0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x68, 0x0a,
	0x17, 0x63, 0x6f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6b, 0x69, 0x74, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x15, 0x63, 0x6f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x70, 0x0a, 0x12, 0x48, 0x6f, 0x73, 0x74, 0x65,
	0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x95, 0x02, 0x0a, 0x15, 0x43, 0x6f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x77, 0x0a, 0x13, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f,
	0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47,
	0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6b, 0x69, 0x74, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x64, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x50, 0x6f, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x1a, 0x44, 0x0a, 0x16, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x36, 0x5a, 0x34, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61,
	0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6b, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_internal_build_buildkit_devhost_proto_rawDescOnce sync.Once
	file_internal_build_buildkit_devhost_proto_rawDescData = file_internal_build_buildkit_devhost_proto_rawDesc
)

func file_internal_build_buildkit_devhost_proto_rawDescGZIP() []byte {
	file_internal_build_buildkit_devhost_proto_rawDescOnce.Do(func() {
		file_internal_build_buildkit_devhost_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_build_buildkit_devhost_proto_rawDescData)
	})
	return file_internal_build_buildkit_devhost_proto_rawDescData
}

var file_internal_build_buildkit_devhost_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_build_buildkit_devhost_proto_goTypes = []interface{}{
	(*Overrides)(nil),             // 0: foundation.build.buildkit.Overrides
	(*HostedBuildCluster)(nil),    // 1: foundation.build.buildkit.HostedBuildCluster
	(*ColocatedBuildCluster)(nil), // 2: foundation.build.buildkit.ColocatedBuildCluster
	nil,                           // 3: foundation.build.buildkit.ColocatedBuildCluster.MatchingPodLabelsEntry
}
var file_internal_build_buildkit_devhost_proto_depIdxs = []int32{
	1, // 0: foundation.build.buildkit.Overrides.hosted_build_cluster:type_name -> foundation.build.buildkit.HostedBuildCluster
	2, // 1: foundation.build.buildkit.Overrides.colocated_build_cluster:type_name -> foundation.build.buildkit.ColocatedBuildCluster
	3, // 2: foundation.build.buildkit.ColocatedBuildCluster.matching_pod_labels:type_name -> foundation.build.buildkit.ColocatedBuildCluster.MatchingPodLabelsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_build_buildkit_devhost_proto_init() }
func file_internal_build_buildkit_devhost_proto_init() {
	if File_internal_build_buildkit_devhost_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_build_buildkit_devhost_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Overrides); i {
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
		file_internal_build_buildkit_devhost_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostedBuildCluster); i {
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
		file_internal_build_buildkit_devhost_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColocatedBuildCluster); i {
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
			RawDescriptor: file_internal_build_buildkit_devhost_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_build_buildkit_devhost_proto_goTypes,
		DependencyIndexes: file_internal_build_buildkit_devhost_proto_depIdxs,
		MessageInfos:      file_internal_build_buildkit_devhost_proto_msgTypes,
	}.Build()
	File_internal_build_buildkit_devhost_proto = out.File
	file_internal_build_buildkit_devhost_proto_rawDesc = nil
	file_internal_build_buildkit_devhost_proto_goTypes = nil
	file_internal_build_buildkit_devhost_proto_depIdxs = nil
}
