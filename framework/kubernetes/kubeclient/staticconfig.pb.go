// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: framework/kubernetes/kubeclient/staticconfig.proto

package kubeclient

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

type StaticConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointAddress          string `protobuf:"bytes,1,opt,name=endpoint_address,json=endpointAddress,proto3" json:"endpoint_address,omitempty"`
	CertificateAuthorityData []byte `protobuf:"bytes,2,opt,name=certificate_authority_data,json=certificateAuthorityData,proto3" json:"certificate_authority_data,omitempty"`
	ClientCertificateData    []byte `protobuf:"bytes,3,opt,name=client_certificate_data,json=clientCertificateData,proto3" json:"client_certificate_data,omitempty"`
	ClientKeyData            []byte `protobuf:"bytes,4,opt,name=client_key_data,json=clientKeyData,proto3" json:"client_key_data,omitempty"`
}

func (x *StaticConfig) Reset() {
	*x = StaticConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_framework_kubernetes_kubeclient_staticconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticConfig) ProtoMessage() {}

func (x *StaticConfig) ProtoReflect() protoreflect.Message {
	mi := &file_framework_kubernetes_kubeclient_staticconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticConfig.ProtoReflect.Descriptor instead.
func (*StaticConfig) Descriptor() ([]byte, []int) {
	return file_framework_kubernetes_kubeclient_staticconfig_proto_rawDescGZIP(), []int{0}
}

func (x *StaticConfig) GetEndpointAddress() string {
	if x != nil {
		return x.EndpointAddress
	}
	return ""
}

func (x *StaticConfig) GetCertificateAuthorityData() []byte {
	if x != nil {
		return x.CertificateAuthorityData
	}
	return nil
}

func (x *StaticConfig) GetClientCertificateData() []byte {
	if x != nil {
		return x.ClientCertificateData
	}
	return nil
}

func (x *StaticConfig) GetClientKeyData() []byte {
	if x != nil {
		return x.ClientKeyData
	}
	return nil
}

var File_framework_kubernetes_kubeclient_staticconfig_proto protoreflect.FileDescriptor

var file_framework_kubernetes_kubeclient_staticconfig_proto_rawDesc = []byte{
	0x0a, 0x32, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x22, 0xd7, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3c, 0x0a, 0x1a,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x18, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x15, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x44, 0x61, 0x74, 0x61, 0x42, 0x3e, 0x5a, 0x3c, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_framework_kubernetes_kubeclient_staticconfig_proto_rawDescOnce sync.Once
	file_framework_kubernetes_kubeclient_staticconfig_proto_rawDescData = file_framework_kubernetes_kubeclient_staticconfig_proto_rawDesc
)

func file_framework_kubernetes_kubeclient_staticconfig_proto_rawDescGZIP() []byte {
	file_framework_kubernetes_kubeclient_staticconfig_proto_rawDescOnce.Do(func() {
		file_framework_kubernetes_kubeclient_staticconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_framework_kubernetes_kubeclient_staticconfig_proto_rawDescData)
	})
	return file_framework_kubernetes_kubeclient_staticconfig_proto_rawDescData
}

var file_framework_kubernetes_kubeclient_staticconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_framework_kubernetes_kubeclient_staticconfig_proto_goTypes = []interface{}{
	(*StaticConfig)(nil), // 0: foundation.framework.kubernetes.kubeclient.StaticConfig
}
var file_framework_kubernetes_kubeclient_staticconfig_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_framework_kubernetes_kubeclient_staticconfig_proto_init() }
func file_framework_kubernetes_kubeclient_staticconfig_proto_init() {
	if File_framework_kubernetes_kubeclient_staticconfig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_framework_kubernetes_kubeclient_staticconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticConfig); i {
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
			RawDescriptor: file_framework_kubernetes_kubeclient_staticconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_framework_kubernetes_kubeclient_staticconfig_proto_goTypes,
		DependencyIndexes: file_framework_kubernetes_kubeclient_staticconfig_proto_depIdxs,
		MessageInfos:      file_framework_kubernetes_kubeclient_staticconfig_proto_msgTypes,
	}.Build()
	File_framework_kubernetes_kubeclient_staticconfig_proto = out.File
	file_framework_kubernetes_kubeclient_staticconfig_proto_rawDesc = nil
	file_framework_kubernetes_kubeclient_staticconfig_proto_goTypes = nil
	file_framework_kubernetes_kubeclient_staticconfig_proto_depIdxs = nil
}