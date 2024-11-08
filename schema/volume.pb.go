// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/volume.proto

package schema

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Mount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner     string      `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"` // Package that declared the mount.
	Path      string      `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Readonly  bool        `protobuf:"varint,4,opt,name=readonly,proto3" json:"readonly,omitempty"`
	VolumeRef *PackageRef `protobuf:"bytes,5,opt,name=volume_ref,json=volumeRef,proto3" json:"volume_ref,omitempty"`
}

func (x *Mount) Reset() {
	*x = Mount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_volume_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mount) ProtoMessage() {}

func (x *Mount) ProtoReflect() protoreflect.Message {
	mi := &file_schema_volume_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mount.ProtoReflect.Descriptor instead.
func (*Mount) Descriptor() ([]byte, []int) {
	return file_schema_volume_proto_rawDescGZIP(), []int{0}
}

func (x *Mount) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Mount) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Mount) GetReadonly() bool {
	if x != nil {
		return x.Readonly
	}
	return false
}

func (x *Mount) GetVolumeRef() *PackageRef {
	if x != nil {
		return x.VolumeRef
	}
	return nil
}

type Volume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"` // Package that declared the volume.
	// The type of volume.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Explicit if defined at the file level, auto-generated if inlined.
	// Volumes then referenced by the name from "mounts".
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// If true, runtimes may generate a new, compliant name for this volume and its mounts.
	Inline     bool       `protobuf:"varint,5,opt,name=inline,proto3" json:"inline,omitempty"`
	Definition *anypb.Any `protobuf:"bytes,4,opt,name=definition,proto3" json:"definition,omitempty"`
}

func (x *Volume) Reset() {
	*x = Volume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_volume_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Volume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Volume) ProtoMessage() {}

func (x *Volume) ProtoReflect() protoreflect.Message {
	mi := &file_schema_volume_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Volume.ProtoReflect.Descriptor instead.
func (*Volume) Descriptor() ([]byte, []int) {
	return file_schema_volume_proto_rawDescGZIP(), []int{1}
}

func (x *Volume) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Volume) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Volume) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Volume) GetInline() bool {
	if x != nil {
		return x.Inline
	}
	return false
}

func (x *Volume) GetDefinition() *anypb.Any {
	if x != nil {
		return x.Definition
	}
	return nil
}

type EphemeralVolume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SizeBytes uint64 `protobuf:"varint,1,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
}

func (x *EphemeralVolume) Reset() {
	*x = EphemeralVolume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_volume_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EphemeralVolume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EphemeralVolume) ProtoMessage() {}

func (x *EphemeralVolume) ProtoReflect() protoreflect.Message {
	mi := &file_schema_volume_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EphemeralVolume.ProtoReflect.Descriptor instead.
func (*EphemeralVolume) Descriptor() ([]byte, []int) {
	return file_schema_volume_proto_rawDescGZIP(), []int{2}
}

func (x *EphemeralVolume) GetSizeBytes() uint64 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

type PersistentVolume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SizeBytes    uint64 `protobuf:"varint,2,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	Template     bool   `protobuf:"varint,3,opt,name=template,proto3" json:"template,omitempty"`
	StorageClass string `protobuf:"bytes,4,opt,name=storage_class,json=storageClass,proto3" json:"storage_class,omitempty"`
}

func (x *PersistentVolume) Reset() {
	*x = PersistentVolume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_volume_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersistentVolume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersistentVolume) ProtoMessage() {}

func (x *PersistentVolume) ProtoReflect() protoreflect.Message {
	mi := &file_schema_volume_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersistentVolume.ProtoReflect.Descriptor instead.
func (*PersistentVolume) Descriptor() ([]byte, []int) {
	return file_schema_volume_proto_rawDescGZIP(), []int{3}
}

func (x *PersistentVolume) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PersistentVolume) GetSizeBytes() uint64 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

func (x *PersistentVolume) GetTemplate() bool {
	if x != nil {
		return x.Template
	}
	return false
}

func (x *PersistentVolume) GetStorageClass() string {
	if x != nil {
		return x.StorageClass
	}
	return ""
}

type HostPathVolume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Directory string `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
}

func (x *HostPathVolume) Reset() {
	*x = HostPathVolume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_volume_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostPathVolume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostPathVolume) ProtoMessage() {}

func (x *HostPathVolume) ProtoReflect() protoreflect.Message {
	mi := &file_schema_volume_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostPathVolume.ProtoReflect.Descriptor instead.
func (*HostPathVolume) Descriptor() ([]byte, []int) {
	return file_schema_volume_proto_rawDescGZIP(), []int{4}
}

func (x *HostPathVolume) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

type WorkspaceSyncVolume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A directory, relative to package.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *WorkspaceSyncVolume) Reset() {
	*x = WorkspaceSyncVolume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_volume_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceSyncVolume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceSyncVolume) ProtoMessage() {}

func (x *WorkspaceSyncVolume) ProtoReflect() protoreflect.Message {
	mi := &file_schema_volume_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceSyncVolume.ProtoReflect.Descriptor instead.
func (*WorkspaceSyncVolume) Descriptor() ([]byte, []int) {
	return file_schema_volume_proto_rawDescGZIP(), []int{5}
}

func (x *WorkspaceSyncVolume) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ConfigurableVolume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*ConfigurableVolume_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *ConfigurableVolume) Reset() {
	*x = ConfigurableVolume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_volume_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigurableVolume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigurableVolume) ProtoMessage() {}

func (x *ConfigurableVolume) ProtoReflect() protoreflect.Message {
	mi := &file_schema_volume_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigurableVolume.ProtoReflect.Descriptor instead.
func (*ConfigurableVolume) Descriptor() ([]byte, []int) {
	return file_schema_volume_proto_rawDescGZIP(), []int{6}
}

func (x *ConfigurableVolume) GetEntries() []*ConfigurableVolume_Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type ResourceSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource []*FileContents `protobuf:"bytes,1,rep,name=resource,proto3" json:"resource,omitempty"`
}

func (x *ResourceSet) Reset() {
	*x = ResourceSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_volume_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceSet) ProtoMessage() {}

func (x *ResourceSet) ProtoReflect() protoreflect.Message {
	mi := &file_schema_volume_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceSet.ProtoReflect.Descriptor instead.
func (*ResourceSet) Descriptor() ([]byte, []int) {
	return file_schema_volume_proto_rawDescGZIP(), []int{7}
}

func (x *ResourceSet) GetResource() []*FileContents {
	if x != nil {
		return x.Resource
	}
	return nil
}

type ConfigurableVolume_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A single file or a directory, relative to the mount path.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// One of.
	Inline              *FileContents                                 `protobuf:"bytes,2,opt,name=inline,proto3" json:"inline,omitempty"`
	InlineSet           *ResourceSet                                  `protobuf:"bytes,4,opt,name=inline_set,json=inlineSet,proto3" json:"inline_set,omitempty"`
	SecretRef           *PackageRef                                   `protobuf:"bytes,6,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
	KubernetesSecretRef *ConfigurableVolume_Entry_KubernetesSecretRef `protobuf:"bytes,5,opt,name=kubernetes_secret_ref,json=kubernetesSecretRef,proto3" json:"kubernetes_secret_ref,omitempty"`
}

func (x *ConfigurableVolume_Entry) Reset() {
	*x = ConfigurableVolume_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_volume_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigurableVolume_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigurableVolume_Entry) ProtoMessage() {}

func (x *ConfigurableVolume_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_schema_volume_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigurableVolume_Entry.ProtoReflect.Descriptor instead.
func (*ConfigurableVolume_Entry) Descriptor() ([]byte, []int) {
	return file_schema_volume_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ConfigurableVolume_Entry) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ConfigurableVolume_Entry) GetInline() *FileContents {
	if x != nil {
		return x.Inline
	}
	return nil
}

func (x *ConfigurableVolume_Entry) GetInlineSet() *ResourceSet {
	if x != nil {
		return x.InlineSet
	}
	return nil
}

func (x *ConfigurableVolume_Entry) GetSecretRef() *PackageRef {
	if x != nil {
		return x.SecretRef
	}
	return nil
}

func (x *ConfigurableVolume_Entry) GetKubernetesSecretRef() *ConfigurableVolume_Entry_KubernetesSecretRef {
	if x != nil {
		return x.KubernetesSecretRef
	}
	return nil
}

type ConfigurableVolume_Entry_KubernetesSecretRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretName string `protobuf:"bytes,1,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
	SecretKey  string `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
}

func (x *ConfigurableVolume_Entry_KubernetesSecretRef) Reset() {
	*x = ConfigurableVolume_Entry_KubernetesSecretRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_volume_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigurableVolume_Entry_KubernetesSecretRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigurableVolume_Entry_KubernetesSecretRef) ProtoMessage() {}

func (x *ConfigurableVolume_Entry_KubernetesSecretRef) ProtoReflect() protoreflect.Message {
	mi := &file_schema_volume_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigurableVolume_Entry_KubernetesSecretRef.ProtoReflect.Descriptor instead.
func (*ConfigurableVolume_Entry_KubernetesSecretRef) Descriptor() ([]byte, []int) {
	return file_schema_volume_proto_rawDescGZIP(), []int{6, 0, 0}
}

func (x *ConfigurableVolume_Entry_KubernetesSecretRef) GetSecretName() string {
	if x != nil {
		return x.SecretName
	}
	return ""
}

func (x *ConfigurableVolume_Entry_KubernetesSecretRef) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

var File_schema_volume_proto protoreflect.FileDescriptor

var file_schema_volume_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x05, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64,
	0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64,
	0x6f, 0x6e, 0x6c, 0x79, 0x12, 0x3c, 0x0a, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x72,
	0x65, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x52, 0x09, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52,
	0x65, 0x66, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x94, 0x01, 0x0a, 0x06, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x30, 0x0a, 0x0f, 0x45, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x22, 0x82, 0x01, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x22, 0x2e, 0x0a, 0x0e, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x61,
	0x74, 0x68, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x29, 0x0a, 0x13, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x22, 0x81, 0x04, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x62,
	0x6c, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a,
	0xa3, 0x03, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x37, 0x0a,
	0x06, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x06,
	0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x5f, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x74, 0x52, 0x09, 0x69, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x53, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f,
	0x72, 0x65, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x52, 0x65, 0x66, 0x12, 0x73, 0x0a, 0x15, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x62, 0x6c, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x52, 0x65, 0x66, 0x52, 0x13, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x66, 0x1a, 0x55, 0x0a, 0x13, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x66, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x4a,
	0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x4a, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x65, 0x74, 0x12, 0x3b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x25, 0x5a, 0x23, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61,
	0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_volume_proto_rawDescOnce sync.Once
	file_schema_volume_proto_rawDescData = file_schema_volume_proto_rawDesc
)

func file_schema_volume_proto_rawDescGZIP() []byte {
	file_schema_volume_proto_rawDescOnce.Do(func() {
		file_schema_volume_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_volume_proto_rawDescData)
	})
	return file_schema_volume_proto_rawDescData
}

var file_schema_volume_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_schema_volume_proto_goTypes = []interface{}{
	(*Mount)(nil),                                        // 0: foundation.schema.Mount
	(*Volume)(nil),                                       // 1: foundation.schema.Volume
	(*EphemeralVolume)(nil),                              // 2: foundation.schema.EphemeralVolume
	(*PersistentVolume)(nil),                             // 3: foundation.schema.PersistentVolume
	(*HostPathVolume)(nil),                               // 4: foundation.schema.HostPathVolume
	(*WorkspaceSyncVolume)(nil),                          // 5: foundation.schema.WorkspaceSyncVolume
	(*ConfigurableVolume)(nil),                           // 6: foundation.schema.ConfigurableVolume
	(*ResourceSet)(nil),                                  // 7: foundation.schema.ResourceSet
	(*ConfigurableVolume_Entry)(nil),                     // 8: foundation.schema.ConfigurableVolume.Entry
	(*ConfigurableVolume_Entry_KubernetesSecretRef)(nil), // 9: foundation.schema.ConfigurableVolume.Entry.KubernetesSecretRef
	(*PackageRef)(nil),                                   // 10: foundation.schema.PackageRef
	(*anypb.Any)(nil),                                    // 11: google.protobuf.Any
	(*FileContents)(nil),                                 // 12: foundation.schema.FileContents
}
var file_schema_volume_proto_depIdxs = []int32{
	10, // 0: foundation.schema.Mount.volume_ref:type_name -> foundation.schema.PackageRef
	11, // 1: foundation.schema.Volume.definition:type_name -> google.protobuf.Any
	8,  // 2: foundation.schema.ConfigurableVolume.entries:type_name -> foundation.schema.ConfigurableVolume.Entry
	12, // 3: foundation.schema.ResourceSet.resource:type_name -> foundation.schema.FileContents
	12, // 4: foundation.schema.ConfigurableVolume.Entry.inline:type_name -> foundation.schema.FileContents
	7,  // 5: foundation.schema.ConfigurableVolume.Entry.inline_set:type_name -> foundation.schema.ResourceSet
	10, // 6: foundation.schema.ConfigurableVolume.Entry.secret_ref:type_name -> foundation.schema.PackageRef
	9,  // 7: foundation.schema.ConfigurableVolume.Entry.kubernetes_secret_ref:type_name -> foundation.schema.ConfigurableVolume.Entry.KubernetesSecretRef
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_schema_volume_proto_init() }
func file_schema_volume_proto_init() {
	if File_schema_volume_proto != nil {
		return
	}
	file_schema_package_proto_init()
	file_schema_filecontents_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_schema_volume_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mount); i {
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
		file_schema_volume_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Volume); i {
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
		file_schema_volume_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EphemeralVolume); i {
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
		file_schema_volume_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersistentVolume); i {
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
		file_schema_volume_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostPathVolume); i {
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
		file_schema_volume_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceSyncVolume); i {
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
		file_schema_volume_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigurableVolume); i {
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
		file_schema_volume_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceSet); i {
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
		file_schema_volume_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigurableVolume_Entry); i {
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
		file_schema_volume_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigurableVolume_Entry_KubernetesSecretRef); i {
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
			RawDescriptor: file_schema_volume_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_volume_proto_goTypes,
		DependencyIndexes: file_schema_volume_proto_depIdxs,
		MessageInfos:      file_schema_volume_proto_msgTypes,
	}.Build()
	File_schema_volume_proto = out.File
	file_schema_volume_proto_rawDesc = nil
	file_schema_volume_proto_goTypes = nil
	file_schema_volume_proto_depIdxs = nil
}
