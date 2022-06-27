// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/storage/runs.proto

package storage

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IndividualRun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentRunId string                 `protobuf:"bytes,1,opt,name=parent_run_id,json=parentRunId,proto3" json:"parent_run_id,omitempty"`
	Status      *status.Status         `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Created     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Completed   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=completed,proto3" json:"completed,omitempty"` // Regardless of success or failure.
	Attachment  []*anypb.Any           `protobuf:"bytes,5,rep,name=attachment,proto3" json:"attachment,omitempty"`
}

func (x *IndividualRun) Reset() {
	*x = IndividualRun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_storage_runs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndividualRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndividualRun) ProtoMessage() {}

func (x *IndividualRun) ProtoReflect() protoreflect.Message {
	mi := &file_schema_storage_runs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndividualRun.ProtoReflect.Descriptor instead.
func (*IndividualRun) Descriptor() ([]byte, []int) {
	return file_schema_storage_runs_proto_rawDescGZIP(), []int{0}
}

func (x *IndividualRun) GetParentRunId() string {
	if x != nil {
		return x.ParentRunId
	}
	return ""
}

func (x *IndividualRun) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *IndividualRun) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *IndividualRun) GetCompleted() *timestamppb.Timestamp {
	if x != nil {
		return x.Completed
	}
	return nil
}

func (x *IndividualRun) GetAttachment() []*anypb.Any {
	if x != nil {
		return x.Attachment
	}
	return nil
}

var File_schema_storage_runs_proto protoreflect.FileDescriptor

var file_schema_storage_runs_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x72, 0x75, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x02, 0x0a, 0x0d,
	0x49, 0x6e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x52, 0x75, 0x6e, 0x12, 0x22, 0x0a,
	0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x75, 0x6e, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a,
	0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_storage_runs_proto_rawDescOnce sync.Once
	file_schema_storage_runs_proto_rawDescData = file_schema_storage_runs_proto_rawDesc
)

func file_schema_storage_runs_proto_rawDescGZIP() []byte {
	file_schema_storage_runs_proto_rawDescOnce.Do(func() {
		file_schema_storage_runs_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_storage_runs_proto_rawDescData)
	})
	return file_schema_storage_runs_proto_rawDescData
}

var file_schema_storage_runs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_schema_storage_runs_proto_goTypes = []interface{}{
	(*IndividualRun)(nil),         // 0: foundation.schema.storage.IndividualRun
	(*status.Status)(nil),         // 1: google.rpc.Status
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 3: google.protobuf.Any
}
var file_schema_storage_runs_proto_depIdxs = []int32{
	1, // 0: foundation.schema.storage.IndividualRun.status:type_name -> google.rpc.Status
	2, // 1: foundation.schema.storage.IndividualRun.created:type_name -> google.protobuf.Timestamp
	2, // 2: foundation.schema.storage.IndividualRun.completed:type_name -> google.protobuf.Timestamp
	3, // 3: foundation.schema.storage.IndividualRun.attachment:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_schema_storage_runs_proto_init() }
func file_schema_storage_runs_proto_init() {
	if File_schema_storage_runs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_storage_runs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndividualRun); i {
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
			RawDescriptor: file_schema_storage_runs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_storage_runs_proto_goTypes,
		DependencyIndexes: file_schema_storage_runs_proto_depIdxs,
		MessageInfos:      file_schema_storage_runs_proto_msgTypes,
	}.Build()
	File_schema_storage_runs_proto = out.File
	file_schema_storage_runs_proto_rawDesc = nil
	file_schema_storage_runs_proto_goTypes = nil
	file_schema_storage_runs_proto_depIdxs = nil
}
