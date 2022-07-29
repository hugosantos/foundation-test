// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/storage/tests.proto

package storage

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Next ID: 11
type TestBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentRunId      string                  `protobuf:"bytes,9,opt,name=parent_run_id,json=parentRunId,proto3" json:"parent_run_id,omitempty"`
	TestPackage      string                  `protobuf:"bytes,6,opt,name=test_package,json=testPackage,proto3" json:"test_package,omitempty"`
	TestName         string                  `protobuf:"bytes,7,opt,name=test_name,json=testName,proto3" json:"test_name,omitempty"`
	Result           *TestResult             `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	ServersUnderTest []string                `protobuf:"bytes,8,rep,name=servers_under_test,json=serversUnderTest,proto3" json:"servers_under_test,omitempty"`
	Created          *timestamppb.Timestamp  `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	Completed        *timestamppb.Timestamp  `protobuf:"bytes,5,opt,name=completed,proto3" json:"completed,omitempty"` // Regardless of success or failure.
	TestLog          *LogRef                 `protobuf:"bytes,2,opt,name=test_log,json=testLog,proto3" json:"test_log,omitempty"`
	ServerLog        []*LogRef               `protobuf:"bytes,3,rep,name=server_log,json=serverLog,proto3" json:"server_log,omitempty"`
	EnvDiagnostics   *EnvironmentDiagnostics `protobuf:"bytes,10,opt,name=env_diagnostics,json=envDiagnostics,proto3" json:"env_diagnostics,omitempty"`
}

func (x *TestBundle) Reset() {
	*x = TestBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_storage_tests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestBundle) ProtoMessage() {}

func (x *TestBundle) ProtoReflect() protoreflect.Message {
	mi := &file_schema_storage_tests_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestBundle.ProtoReflect.Descriptor instead.
func (*TestBundle) Descriptor() ([]byte, []int) {
	return file_schema_storage_tests_proto_rawDescGZIP(), []int{0}
}

func (x *TestBundle) GetParentRunId() string {
	if x != nil {
		return x.ParentRunId
	}
	return ""
}

func (x *TestBundle) GetTestPackage() string {
	if x != nil {
		return x.TestPackage
	}
	return ""
}

func (x *TestBundle) GetTestName() string {
	if x != nil {
		return x.TestName
	}
	return ""
}

func (x *TestBundle) GetResult() *TestResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *TestBundle) GetServersUnderTest() []string {
	if x != nil {
		return x.ServersUnderTest
	}
	return nil
}

func (x *TestBundle) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *TestBundle) GetCompleted() *timestamppb.Timestamp {
	if x != nil {
		return x.Completed
	}
	return nil
}

func (x *TestBundle) GetTestLog() *LogRef {
	if x != nil {
		return x.TestLog
	}
	return nil
}

func (x *TestBundle) GetServerLog() []*LogRef {
	if x != nil {
		return x.ServerLog
	}
	return nil
}

func (x *TestBundle) GetEnvDiagnostics() *EnvironmentDiagnostics {
	if x != nil {
		return x.EnvDiagnostics
	}
	return nil
}

type TestRuns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Run              []*TestRuns_Run              `protobuf:"bytes,1,rep,name=run,proto3" json:"run,omitempty"`
	IncompatibleTest []*TestRuns_IncompatibleTest `protobuf:"bytes,2,rep,name=incompatible_test,json=incompatibleTest,proto3" json:"incompatible_test,omitempty"`
}

func (x *TestRuns) Reset() {
	*x = TestRuns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_storage_tests_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRuns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRuns) ProtoMessage() {}

func (x *TestRuns) ProtoReflect() protoreflect.Message {
	mi := &file_schema_storage_tests_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRuns.ProtoReflect.Descriptor instead.
func (*TestRuns) Descriptor() ([]byte, []int) {
	return file_schema_storage_tests_proto_rawDescGZIP(), []int{1}
}

func (x *TestRuns) GetRun() []*TestRuns_Run {
	if x != nil {
		return x.Run
	}
	return nil
}

func (x *TestRuns) GetIncompatibleTest() []*TestRuns_IncompatibleTest {
	if x != nil {
		return x.IncompatibleTest
	}
	return nil
}

type TestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success      bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrorCode    int32  `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *TestResult) Reset() {
	*x = TestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_storage_tests_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResult) ProtoMessage() {}

func (x *TestResult) ProtoReflect() protoreflect.Message {
	mi := &file_schema_storage_tests_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResult.ProtoReflect.Descriptor instead.
func (*TestResult) Descriptor() ([]byte, []int) {
	return file_schema_storage_tests_proto_rawDescGZIP(), []int{2}
}

func (x *TestResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *TestResult) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *TestResult) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

// Next ID: 7
type TestResultBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result                 *TestResult                    `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	DeployPlan             *schema.DeployPlan             `protobuf:"bytes,4,opt,name=deploy_plan,json=deployPlan,proto3" json:"deploy_plan,omitempty"`
	ComputedConfigurations *schema.ComputedConfigurations `protobuf:"bytes,5,opt,name=computed_configurations,json=computedConfigurations,proto3" json:"computed_configurations,omitempty"`
	TestLog                *TestResultBundle_InlineLog    `protobuf:"bytes,2,opt,name=test_log,json=testLog,proto3" json:"test_log,omitempty"`
	ServerLog              []*TestResultBundle_InlineLog  `protobuf:"bytes,3,rep,name=server_log,json=serverLog,proto3" json:"server_log,omitempty"`
	EnvDiagnostics         *EnvironmentDiagnostics        `protobuf:"bytes,6,opt,name=env_diagnostics,json=envDiagnostics,proto3" json:"env_diagnostics,omitempty"`
}

func (x *TestResultBundle) Reset() {
	*x = TestResultBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_storage_tests_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResultBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResultBundle) ProtoMessage() {}

func (x *TestResultBundle) ProtoReflect() protoreflect.Message {
	mi := &file_schema_storage_tests_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResultBundle.ProtoReflect.Descriptor instead.
func (*TestResultBundle) Descriptor() ([]byte, []int) {
	return file_schema_storage_tests_proto_rawDescGZIP(), []int{3}
}

func (x *TestResultBundle) GetResult() *TestResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *TestResultBundle) GetDeployPlan() *schema.DeployPlan {
	if x != nil {
		return x.DeployPlan
	}
	return nil
}

func (x *TestResultBundle) GetComputedConfigurations() *schema.ComputedConfigurations {
	if x != nil {
		return x.ComputedConfigurations
	}
	return nil
}

func (x *TestResultBundle) GetTestLog() *TestResultBundle_InlineLog {
	if x != nil {
		return x.TestLog
	}
	return nil
}

func (x *TestResultBundle) GetServerLog() []*TestResultBundle_InlineLog {
	if x != nil {
		return x.ServerLog
	}
	return nil
}

func (x *TestResultBundle) GetEnvDiagnostics() *EnvironmentDiagnostics {
	if x != nil {
		return x.EnvDiagnostics
	}
	return nil
}

type TestRuns_Run struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestBundleId string            `protobuf:"bytes,1,opt,name=test_bundle_id,json=testBundleId,proto3" json:"test_bundle_id,omitempty"`
	TestSummary  *TestBundle       `protobuf:"bytes,2,opt,name=test_summary,json=testSummary,proto3" json:"test_summary,omitempty"` // XXX replace with test_results.
	TestResults  *TestResultBundle `protobuf:"bytes,3,opt,name=test_results,json=testResults,proto3" json:"test_results,omitempty"`
}

func (x *TestRuns_Run) Reset() {
	*x = TestRuns_Run{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_storage_tests_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRuns_Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRuns_Run) ProtoMessage() {}

func (x *TestRuns_Run) ProtoReflect() protoreflect.Message {
	mi := &file_schema_storage_tests_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRuns_Run.ProtoReflect.Descriptor instead.
func (*TestRuns_Run) Descriptor() ([]byte, []int) {
	return file_schema_storage_tests_proto_rawDescGZIP(), []int{1, 0}
}

func (x *TestRuns_Run) GetTestBundleId() string {
	if x != nil {
		return x.TestBundleId
	}
	return ""
}

func (x *TestRuns_Run) GetTestSummary() *TestBundle {
	if x != nil {
		return x.TestSummary
	}
	return nil
}

func (x *TestRuns_Run) GetTestResults() *TestResultBundle {
	if x != nil {
		return x.TestResults
	}
	return nil
}

type TestRuns_IncompatibleTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestPackage       string                    `protobuf:"bytes,1,opt,name=test_package,json=testPackage,proto3" json:"test_package,omitempty"`
	ServerPackage     string                    `protobuf:"bytes,2,opt,name=server_package,json=serverPackage,proto3" json:"server_package,omitempty"`
	RequirementOwner  string                    `protobuf:"bytes,3,opt,name=requirement_owner,json=requirementOwner,proto3" json:"requirement_owner,omitempty"`
	RequiredLabel     *schema.Environment_Label `protobuf:"bytes,4,opt,name=required_label,json=requiredLabel,proto3" json:"required_label,omitempty"`
	IncompatibleLabel *schema.Environment_Label `protobuf:"bytes,5,opt,name=incompatible_label,json=incompatibleLabel,proto3" json:"incompatible_label,omitempty"`
}

func (x *TestRuns_IncompatibleTest) Reset() {
	*x = TestRuns_IncompatibleTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_storage_tests_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRuns_IncompatibleTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRuns_IncompatibleTest) ProtoMessage() {}

func (x *TestRuns_IncompatibleTest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_storage_tests_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRuns_IncompatibleTest.ProtoReflect.Descriptor instead.
func (*TestRuns_IncompatibleTest) Descriptor() ([]byte, []int) {
	return file_schema_storage_tests_proto_rawDescGZIP(), []int{1, 1}
}

func (x *TestRuns_IncompatibleTest) GetTestPackage() string {
	if x != nil {
		return x.TestPackage
	}
	return ""
}

func (x *TestRuns_IncompatibleTest) GetServerPackage() string {
	if x != nil {
		return x.ServerPackage
	}
	return ""
}

func (x *TestRuns_IncompatibleTest) GetRequirementOwner() string {
	if x != nil {
		return x.RequirementOwner
	}
	return ""
}

func (x *TestRuns_IncompatibleTest) GetRequiredLabel() *schema.Environment_Label {
	if x != nil {
		return x.RequiredLabel
	}
	return nil
}

func (x *TestRuns_IncompatibleTest) GetIncompatibleLabel() *schema.Environment_Label {
	if x != nil {
		return x.IncompatibleLabel
	}
	return nil
}

type TestResultBundle_InlineLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageName   string               `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	ContainerName string               `protobuf:"bytes,3,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	ContainerKind schema.ContainerKind `protobuf:"varint,4,opt,name=container_kind,json=containerKind,proto3,enum=foundation.schema.ContainerKind" json:"container_kind,omitempty"`
	Output        []byte               `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *TestResultBundle_InlineLog) Reset() {
	*x = TestResultBundle_InlineLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_storage_tests_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResultBundle_InlineLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResultBundle_InlineLog) ProtoMessage() {}

func (x *TestResultBundle_InlineLog) ProtoReflect() protoreflect.Message {
	mi := &file_schema_storage_tests_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResultBundle_InlineLog.ProtoReflect.Descriptor instead.
func (*TestResultBundle_InlineLog) Descriptor() ([]byte, []int) {
	return file_schema_storage_tests_proto_rawDescGZIP(), []int{3, 0}
}

func (x *TestResultBundle_InlineLog) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *TestResultBundle_InlineLog) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *TestResultBundle_InlineLog) GetContainerKind() schema.ContainerKind {
	if x != nil {
		return x.ContainerKind
	}
	return schema.ContainerKind(0)
}

func (x *TestResultBundle_InlineLog) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_schema_storage_tests_proto protoreflect.FileDescriptor

var file_schema_storage_tests_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa9, 0x04, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x75, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x52, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x65, 0x73,
	0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x5f,
	0x75, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x65,
	0x73, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x12, 0x3c, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x66, 0x52, 0x07, 0x74, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67,
	0x12, 0x40, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x66, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x12, 0x5a, 0x0a, 0x0f, 0x65, 0x6e, 0x76, 0x5f, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x0e,
	0x65, 0x6e, 0x76, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x22, 0x9e,
	0x05, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x03, 0x72,
	0x75, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x73, 0x2e, 0x52, 0x75,
	0x6e, 0x52, 0x03, 0x72, 0x75, 0x6e, 0x12, 0x61, 0x0a, 0x11, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x75, 0x6e, 0x73, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69,
	0x62, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x10, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x74, 0x69, 0x62, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x1a, 0xc5, 0x01, 0x0a, 0x03, 0x52, 0x75,
	0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x52, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x12, 0x4e, 0x0a, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x52, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x1a, 0xab, 0x02, 0x0a, 0x10, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62,
	0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x65,
	0x73, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x4b, 0x0a,
	0x0e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x0d, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x53, 0x0a, 0x12, 0x69, 0x6e,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x11, 0x69, 0x6e,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x22,
	0x6a, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xb2, 0x05, 0x0a, 0x10,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x12, 0x3d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x12,
	0x62, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x16, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x50, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x2e, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x07, 0x74, 0x65,
	0x73, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x54, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x6c, 0x6f, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x6f, 0x67,
	0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x12, 0x5a, 0x0a, 0x0f, 0x65,
	0x6e, 0x76, 0x5f, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x61, 0x67,
	0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x0e, 0x65, 0x6e, 0x76, 0x44, 0x69, 0x61, 0x67,
	0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x1a, 0xb6, 0x01, 0x0a, 0x09, 0x49, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x47, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x42, 0x2d, 0x5a, 0x2b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62,
	0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_storage_tests_proto_rawDescOnce sync.Once
	file_schema_storage_tests_proto_rawDescData = file_schema_storage_tests_proto_rawDesc
)

func file_schema_storage_tests_proto_rawDescGZIP() []byte {
	file_schema_storage_tests_proto_rawDescOnce.Do(func() {
		file_schema_storage_tests_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_storage_tests_proto_rawDescData)
	})
	return file_schema_storage_tests_proto_rawDescData
}

var file_schema_storage_tests_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_schema_storage_tests_proto_goTypes = []interface{}{
	(*TestBundle)(nil),                    // 0: foundation.schema.storage.TestBundle
	(*TestRuns)(nil),                      // 1: foundation.schema.storage.TestRuns
	(*TestResult)(nil),                    // 2: foundation.schema.storage.TestResult
	(*TestResultBundle)(nil),              // 3: foundation.schema.storage.TestResultBundle
	(*TestRuns_Run)(nil),                  // 4: foundation.schema.storage.TestRuns.Run
	(*TestRuns_IncompatibleTest)(nil),     // 5: foundation.schema.storage.TestRuns.IncompatibleTest
	(*TestResultBundle_InlineLog)(nil),    // 6: foundation.schema.storage.TestResultBundle.InlineLog
	(*timestamppb.Timestamp)(nil),         // 7: google.protobuf.Timestamp
	(*LogRef)(nil),                        // 8: foundation.schema.storage.LogRef
	(*EnvironmentDiagnostics)(nil),        // 9: foundation.schema.storage.EnvironmentDiagnostics
	(*schema.DeployPlan)(nil),             // 10: foundation.schema.DeployPlan
	(*schema.ComputedConfigurations)(nil), // 11: foundation.schema.ComputedConfigurations
	(*schema.Environment_Label)(nil),      // 12: foundation.schema.Environment.Label
	(schema.ContainerKind)(0),             // 13: foundation.schema.ContainerKind
}
var file_schema_storage_tests_proto_depIdxs = []int32{
	2,  // 0: foundation.schema.storage.TestBundle.result:type_name -> foundation.schema.storage.TestResult
	7,  // 1: foundation.schema.storage.TestBundle.created:type_name -> google.protobuf.Timestamp
	7,  // 2: foundation.schema.storage.TestBundle.completed:type_name -> google.protobuf.Timestamp
	8,  // 3: foundation.schema.storage.TestBundle.test_log:type_name -> foundation.schema.storage.LogRef
	8,  // 4: foundation.schema.storage.TestBundle.server_log:type_name -> foundation.schema.storage.LogRef
	9,  // 5: foundation.schema.storage.TestBundle.env_diagnostics:type_name -> foundation.schema.storage.EnvironmentDiagnostics
	4,  // 6: foundation.schema.storage.TestRuns.run:type_name -> foundation.schema.storage.TestRuns.Run
	5,  // 7: foundation.schema.storage.TestRuns.incompatible_test:type_name -> foundation.schema.storage.TestRuns.IncompatibleTest
	2,  // 8: foundation.schema.storage.TestResultBundle.result:type_name -> foundation.schema.storage.TestResult
	10, // 9: foundation.schema.storage.TestResultBundle.deploy_plan:type_name -> foundation.schema.DeployPlan
	11, // 10: foundation.schema.storage.TestResultBundle.computed_configurations:type_name -> foundation.schema.ComputedConfigurations
	6,  // 11: foundation.schema.storage.TestResultBundle.test_log:type_name -> foundation.schema.storage.TestResultBundle.InlineLog
	6,  // 12: foundation.schema.storage.TestResultBundle.server_log:type_name -> foundation.schema.storage.TestResultBundle.InlineLog
	9,  // 13: foundation.schema.storage.TestResultBundle.env_diagnostics:type_name -> foundation.schema.storage.EnvironmentDiagnostics
	0,  // 14: foundation.schema.storage.TestRuns.Run.test_summary:type_name -> foundation.schema.storage.TestBundle
	3,  // 15: foundation.schema.storage.TestRuns.Run.test_results:type_name -> foundation.schema.storage.TestResultBundle
	12, // 16: foundation.schema.storage.TestRuns.IncompatibleTest.required_label:type_name -> foundation.schema.Environment.Label
	12, // 17: foundation.schema.storage.TestRuns.IncompatibleTest.incompatible_label:type_name -> foundation.schema.Environment.Label
	13, // 18: foundation.schema.storage.TestResultBundle.InlineLog.container_kind:type_name -> foundation.schema.ContainerKind
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_schema_storage_tests_proto_init() }
func file_schema_storage_tests_proto_init() {
	if File_schema_storage_tests_proto != nil {
		return
	}
	file_schema_storage_logs_proto_init()
	file_schema_storage_diagnostics_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_schema_storage_tests_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestBundle); i {
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
		file_schema_storage_tests_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRuns); i {
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
		file_schema_storage_tests_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResult); i {
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
		file_schema_storage_tests_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResultBundle); i {
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
		file_schema_storage_tests_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRuns_Run); i {
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
		file_schema_storage_tests_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRuns_IncompatibleTest); i {
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
		file_schema_storage_tests_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResultBundle_InlineLog); i {
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
			RawDescriptor: file_schema_storage_tests_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_storage_tests_proto_goTypes,
		DependencyIndexes: file_schema_storage_tests_proto_depIdxs,
		MessageInfos:      file_schema_storage_tests_proto_msgTypes,
	}.Build()
	File_schema_storage_tests_proto = out.File
	file_schema_storage_tests_proto_rawDesc = nil
	file_schema_storage_tests_proto_goTypes = nil
	file_schema_storage_tests_proto_depIdxs = nil
}
