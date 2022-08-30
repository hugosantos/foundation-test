// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: runtime/interface.proto

package runtime

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type ContainerWaitStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Containers   []*ContainerUnitWaitStatus `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty"`
	Initializers []*ContainerUnitWaitStatus `protobuf:"bytes,2,rep,name=initializers,proto3" json:"initializers,omitempty"`
}

func (x *ContainerWaitStatus) Reset() {
	*x = ContainerWaitStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runtime_interface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerWaitStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerWaitStatus) ProtoMessage() {}

func (x *ContainerWaitStatus) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_interface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerWaitStatus.ProtoReflect.Descriptor instead.
func (*ContainerWaitStatus) Descriptor() ([]byte, []int) {
	return file_runtime_interface_proto_rawDescGZIP(), []int{0}
}

func (x *ContainerWaitStatus) GetContainers() []*ContainerUnitWaitStatus {
	if x != nil {
		return x.Containers
	}
	return nil
}

func (x *ContainerWaitStatus) GetInitializers() []*ContainerUnitWaitStatus {
	if x != nil {
		return x.Initializers
	}
	return nil
}

type ContainerUnitWaitStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reference   *ContainerReference `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`
	Name        string              `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	StatusLabel string              `protobuf:"bytes,3,opt,name=status_label,json=statusLabel,proto3" json:"status_label,omitempty"`
	Status      *Diagnostics        `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ContainerUnitWaitStatus) Reset() {
	*x = ContainerUnitWaitStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runtime_interface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerUnitWaitStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerUnitWaitStatus) ProtoMessage() {}

func (x *ContainerUnitWaitStatus) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_interface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerUnitWaitStatus.ProtoReflect.Descriptor instead.
func (*ContainerUnitWaitStatus) Descriptor() ([]byte, []int) {
	return file_runtime_interface_proto_rawDescGZIP(), []int{1}
}

func (x *ContainerUnitWaitStatus) GetReference() *ContainerReference {
	if x != nil {
		return x.Reference
	}
	return nil
}

func (x *ContainerUnitWaitStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContainerUnitWaitStatus) GetStatusLabel() string {
	if x != nil {
		return x.StatusLabel
	}
	return ""
}

func (x *ContainerUnitWaitStatus) GetStatus() *Diagnostics {
	if x != nil {
		return x.Status
	}
	return nil
}

type ContainerReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueId       string               `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	HumanReference string               `protobuf:"bytes,2,opt,name=HumanReference,proto3" json:"HumanReference,omitempty"`
	Kind           schema.ContainerKind `protobuf:"varint,3,opt,name=Kind,proto3,enum=foundation.schema.ContainerKind" json:"Kind,omitempty"`
	Opaque         *anypb.Any           `protobuf:"bytes,4,opt,name=opaque,proto3" json:"opaque,omitempty"`
}

func (x *ContainerReference) Reset() {
	*x = ContainerReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runtime_interface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerReference) ProtoMessage() {}

func (x *ContainerReference) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_interface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerReference.ProtoReflect.Descriptor instead.
func (*ContainerReference) Descriptor() ([]byte, []int) {
	return file_runtime_interface_proto_rawDescGZIP(), []int{2}
}

func (x *ContainerReference) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

func (x *ContainerReference) GetHumanReference() string {
	if x != nil {
		return x.HumanReference
	}
	return ""
}

func (x *ContainerReference) GetKind() schema.ContainerKind {
	if x != nil {
		return x.Kind
	}
	return schema.ContainerKind(0)
}

func (x *ContainerReference) GetOpaque() *anypb.Any {
	if x != nil {
		return x.Opaque
	}
	return nil
}

type Diagnostics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Running          bool                   `protobuf:"varint,1,opt,name=running,proto3" json:"running,omitempty"`
	Started          *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=started,proto3" json:"started,omitempty"`
	Waiting          bool                   `protobuf:"varint,3,opt,name=waiting,proto3" json:"waiting,omitempty"`
	WaitingReason    string                 `protobuf:"bytes,4,opt,name=waiting_reason,json=waitingReason,proto3" json:"waiting_reason,omitempty"`
	Crashed          bool                   `protobuf:"varint,5,opt,name=crashed,proto3" json:"crashed,omitempty"`
	Terminated       bool                   `protobuf:"varint,6,opt,name=terminated,proto3" json:"terminated,omitempty"`
	TerminatedReason string                 `protobuf:"bytes,7,opt,name=terminated_reason,json=terminatedReason,proto3" json:"terminated_reason,omitempty"`
	ExitCode         int32                  `protobuf:"varint,8,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	RestartCount     int32                  `protobuf:"varint,9,opt,name=restart_count,json=restartCount,proto3" json:"restart_count,omitempty"`
}

func (x *Diagnostics) Reset() {
	*x = Diagnostics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runtime_interface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Diagnostics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Diagnostics) ProtoMessage() {}

func (x *Diagnostics) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_interface_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Diagnostics.ProtoReflect.Descriptor instead.
func (*Diagnostics) Descriptor() ([]byte, []int) {
	return file_runtime_interface_proto_rawDescGZIP(), []int{3}
}

func (x *Diagnostics) GetRunning() bool {
	if x != nil {
		return x.Running
	}
	return false
}

func (x *Diagnostics) GetStarted() *timestamppb.Timestamp {
	if x != nil {
		return x.Started
	}
	return nil
}

func (x *Diagnostics) GetWaiting() bool {
	if x != nil {
		return x.Waiting
	}
	return false
}

func (x *Diagnostics) GetWaitingReason() string {
	if x != nil {
		return x.WaitingReason
	}
	return ""
}

func (x *Diagnostics) GetCrashed() bool {
	if x != nil {
		return x.Crashed
	}
	return false
}

func (x *Diagnostics) GetTerminated() bool {
	if x != nil {
		return x.Terminated
	}
	return false
}

func (x *Diagnostics) GetTerminatedReason() string {
	if x != nil {
		return x.TerminatedReason
	}
	return ""
}

func (x *Diagnostics) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *Diagnostics) GetRestartCount() int32 {
	if x != nil {
		return x.RestartCount
	}
	return 0
}

var File_runtime_interface_proto protoreflect.FileDescriptor

var file_runtime_interface_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb3, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x57, 0x61, 0x69,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4b, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x57, 0x61,
	0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x73, 0x12, 0x4f, 0x0a, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x57, 0x61, 0x69,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x72, 0x73, 0x22, 0xcf, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x57, 0x61, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x44, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x37,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xbd, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x48,
	0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61,
	0x71, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x22, 0xc7, 0x02, 0x0a, 0x0b, 0x44, 0x69, 0x61, 0x67,
	0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x61, 0x69, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x77, 0x61, 0x69, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x25, 0x0a, 0x0e, 0x77, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x61, 0x69, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x61, 0x73,
	0x68, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x72, 0x61, 0x73, 0x68,
	0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x26, 0x5a, 0x24, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61,
	0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_runtime_interface_proto_rawDescOnce sync.Once
	file_runtime_interface_proto_rawDescData = file_runtime_interface_proto_rawDesc
)

func file_runtime_interface_proto_rawDescGZIP() []byte {
	file_runtime_interface_proto_rawDescOnce.Do(func() {
		file_runtime_interface_proto_rawDescData = protoimpl.X.CompressGZIP(file_runtime_interface_proto_rawDescData)
	})
	return file_runtime_interface_proto_rawDescData
}

var file_runtime_interface_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_runtime_interface_proto_goTypes = []interface{}{
	(*ContainerWaitStatus)(nil),     // 0: foundation.runtime.ContainerWaitStatus
	(*ContainerUnitWaitStatus)(nil), // 1: foundation.runtime.ContainerUnitWaitStatus
	(*ContainerReference)(nil),      // 2: foundation.runtime.ContainerReference
	(*Diagnostics)(nil),             // 3: foundation.runtime.Diagnostics
	(schema.ContainerKind)(0),       // 4: foundation.schema.ContainerKind
	(*anypb.Any)(nil),               // 5: google.protobuf.Any
	(*timestamppb.Timestamp)(nil),   // 6: google.protobuf.Timestamp
}
var file_runtime_interface_proto_depIdxs = []int32{
	1, // 0: foundation.runtime.ContainerWaitStatus.containers:type_name -> foundation.runtime.ContainerUnitWaitStatus
	1, // 1: foundation.runtime.ContainerWaitStatus.initializers:type_name -> foundation.runtime.ContainerUnitWaitStatus
	2, // 2: foundation.runtime.ContainerUnitWaitStatus.reference:type_name -> foundation.runtime.ContainerReference
	3, // 3: foundation.runtime.ContainerUnitWaitStatus.status:type_name -> foundation.runtime.Diagnostics
	4, // 4: foundation.runtime.ContainerReference.Kind:type_name -> foundation.schema.ContainerKind
	5, // 5: foundation.runtime.ContainerReference.opaque:type_name -> google.protobuf.Any
	6, // 6: foundation.runtime.Diagnostics.started:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_runtime_interface_proto_init() }
func file_runtime_interface_proto_init() {
	if File_runtime_interface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_runtime_interface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerWaitStatus); i {
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
		file_runtime_interface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerUnitWaitStatus); i {
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
		file_runtime_interface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerReference); i {
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
		file_runtime_interface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Diagnostics); i {
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
			RawDescriptor: file_runtime_interface_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_runtime_interface_proto_goTypes,
		DependencyIndexes: file_runtime_interface_proto_depIdxs,
		MessageInfos:      file_runtime_interface_proto_msgTypes,
	}.Build()
	File_runtime_interface_proto = out.File
	file_runtime_interface_proto_rawDesc = nil
	file_runtime_interface_proto_goTypes = nil
	file_runtime_interface_proto_depIdxs = nil
}
