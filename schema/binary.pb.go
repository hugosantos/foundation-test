// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/binary.proto

package schema

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

type Binary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The package name (computed).
	PackageName string        `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	Name        string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Repository  string        `protobuf:"bytes,3,opt,name=repository,proto3" json:"repository,omitempty"`
	Digest      string        `protobuf:"bytes,4,opt,name=digest,proto3" json:"digest,omitempty"`
	From        *Binary_From  `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`     // Build instructions.
	Config      *BinaryConfig `protobuf:"bytes,6,opt,name=config,proto3" json:"config,omitempty"` // Run instructions.
}

func (x *Binary) Reset() {
	*x = Binary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_binary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Binary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Binary) ProtoMessage() {}

func (x *Binary) ProtoReflect() protoreflect.Message {
	mi := &file_schema_binary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Binary.ProtoReflect.Descriptor instead.
func (*Binary) Descriptor() ([]byte, []int) {
	return file_schema_binary_proto_rawDescGZIP(), []int{0}
}

func (x *Binary) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *Binary) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Binary) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *Binary) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *Binary) GetFrom() *Binary_From {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Binary) GetConfig() *BinaryConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type BinaryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command []string              `protobuf:"bytes,1,rep,name=command,proto3" json:"command,omitempty"`
	Args    []string              `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	Env     []*BinaryConfig_Entry `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
}

func (x *BinaryConfig) Reset() {
	*x = BinaryConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_binary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryConfig) ProtoMessage() {}

func (x *BinaryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_schema_binary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryConfig.ProtoReflect.Descriptor instead.
func (*BinaryConfig) Descriptor() ([]byte, []int) {
	return file_schema_binary_proto_rawDescGZIP(), []int{1}
}

func (x *BinaryConfig) GetCommand() []string {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *BinaryConfig) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *BinaryConfig) GetEnv() []*BinaryConfig_Entry {
	if x != nil {
		return x.Env
	}
	return nil
}

// XXX make this an arbitrary input.
type Binary_From struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoPackage   string `protobuf:"bytes,1,opt,name=go_package,json=goPackage,proto3" json:"go_package,omitempty"`         // Use go binary builder.
	Dockerfile  string `protobuf:"bytes,2,opt,name=dockerfile,proto3" json:"dockerfile,omitempty"`                        // Use Dockerfile builder.
	WebBuild    string `protobuf:"bytes,3,opt,name=web_build,json=webBuild,proto3" json:"web_build,omitempty"`            // Use Web build (temporary).
	LlbGoBinary string `protobuf:"bytes,4,opt,name=llb_go_binary,json=llbGoBinary,proto3" json:"llb_go_binary,omitempty"` // Build a go binary which itself produces LLB.
	NixFlake    string `protobuf:"bytes,5,opt,name=nix_flake,json=nixFlake,proto3" json:"nix_flake,omitempty"`            // Build a docker image from a nix flake.
}

func (x *Binary_From) Reset() {
	*x = Binary_From{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_binary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Binary_From) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Binary_From) ProtoMessage() {}

func (x *Binary_From) ProtoReflect() protoreflect.Message {
	mi := &file_schema_binary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Binary_From.ProtoReflect.Descriptor instead.
func (*Binary_From) Descriptor() ([]byte, []int) {
	return file_schema_binary_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Binary_From) GetGoPackage() string {
	if x != nil {
		return x.GoPackage
	}
	return ""
}

func (x *Binary_From) GetDockerfile() string {
	if x != nil {
		return x.Dockerfile
	}
	return ""
}

func (x *Binary_From) GetWebBuild() string {
	if x != nil {
		return x.WebBuild
	}
	return ""
}

func (x *Binary_From) GetLlbGoBinary() string {
	if x != nil {
		return x.LlbGoBinary
	}
	return ""
}

func (x *Binary_From) GetNixFlake() string {
	if x != nil {
		return x.NixFlake
	}
	return ""
}

type BinaryConfig_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BinaryConfig_Entry) Reset() {
	*x = BinaryConfig_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_binary_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryConfig_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryConfig_Entry) ProtoMessage() {}

func (x *BinaryConfig_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_schema_binary_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryConfig_Entry.ProtoReflect.Descriptor instead.
func (*BinaryConfig_Entry) Descriptor() ([]byte, []int) {
	return file_schema_binary_proto_rawDescGZIP(), []int{1, 0}
}

func (x *BinaryConfig_Entry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BinaryConfig_Entry) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_schema_binary_proto protoreflect.FileDescriptor

var file_schema_binary_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x8a, 0x03, 0x0a, 0x06, 0x42, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x46, 0x72, 0x6f, 0x6d,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x37, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a,
	0xa3, 0x01, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x6f, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x6f,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x65, 0x62, 0x5f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x65, 0x62, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x6c, 0x62, 0x5f, 0x67, 0x6f, 0x5f, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6c, 0x62,
	0x47, 0x6f, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x78, 0x5f,
	0x66, 0x6c, 0x61, 0x6b, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x78,
	0x46, 0x6c, 0x61, 0x6b, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x0c, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x72, 0x67, 0x73, 0x12, 0x37, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x1a, 0x31, 0x0a,
	0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x25, 0x5a, 0x23, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62,
	0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_binary_proto_rawDescOnce sync.Once
	file_schema_binary_proto_rawDescData = file_schema_binary_proto_rawDesc
)

func file_schema_binary_proto_rawDescGZIP() []byte {
	file_schema_binary_proto_rawDescOnce.Do(func() {
		file_schema_binary_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_binary_proto_rawDescData)
	})
	return file_schema_binary_proto_rawDescData
}

var file_schema_binary_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_schema_binary_proto_goTypes = []interface{}{
	(*Binary)(nil),             // 0: foundation.schema.Binary
	(*BinaryConfig)(nil),       // 1: foundation.schema.BinaryConfig
	(*Binary_From)(nil),        // 2: foundation.schema.Binary.From
	(*BinaryConfig_Entry)(nil), // 3: foundation.schema.BinaryConfig.Entry
}
var file_schema_binary_proto_depIdxs = []int32{
	2, // 0: foundation.schema.Binary.from:type_name -> foundation.schema.Binary.From
	1, // 1: foundation.schema.Binary.config:type_name -> foundation.schema.BinaryConfig
	3, // 2: foundation.schema.BinaryConfig.env:type_name -> foundation.schema.BinaryConfig.Entry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_schema_binary_proto_init() }
func file_schema_binary_proto_init() {
	if File_schema_binary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_binary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Binary); i {
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
		file_schema_binary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryConfig); i {
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
		file_schema_binary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Binary_From); i {
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
		file_schema_binary_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryConfig_Entry); i {
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
			RawDescriptor: file_schema_binary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_binary_proto_goTypes,
		DependencyIndexes: file_schema_binary_proto_depIdxs,
		MessageInfos:      file_schema_binary_proto_msgTypes,
	}.Build()
	File_schema_binary_proto = out.File
	file_schema_binary_proto_rawDesc = nil
	file_schema_binary_proto_goTypes = nil
	file_schema_binary_proto_depIdxs = nil
}
