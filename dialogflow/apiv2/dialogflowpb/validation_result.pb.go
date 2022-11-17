// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: google/cloud/dialogflow/v2/validation_result.proto

package dialogflowpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents a level of severity.
type ValidationError_Severity int32

const (
	// Not specified. This value should never be used.
	ValidationError_SEVERITY_UNSPECIFIED ValidationError_Severity = 0
	// The agent doesn't follow Dialogflow best practices.
	ValidationError_INFO ValidationError_Severity = 1
	// The agent may not behave as expected.
	ValidationError_WARNING ValidationError_Severity = 2
	// The agent may experience partial failures.
	ValidationError_ERROR ValidationError_Severity = 3
	// The agent may completely fail.
	ValidationError_CRITICAL ValidationError_Severity = 4
)

// Enum value maps for ValidationError_Severity.
var (
	ValidationError_Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "INFO",
		2: "WARNING",
		3: "ERROR",
		4: "CRITICAL",
	}
	ValidationError_Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"INFO":                 1,
		"WARNING":              2,
		"ERROR":                3,
		"CRITICAL":             4,
	}
)

func (x ValidationError_Severity) Enum() *ValidationError_Severity {
	p := new(ValidationError_Severity)
	*p = x
	return p
}

func (x ValidationError_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValidationError_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_dialogflow_v2_validation_result_proto_enumTypes[0].Descriptor()
}

func (ValidationError_Severity) Type() protoreflect.EnumType {
	return &file_google_cloud_dialogflow_v2_validation_result_proto_enumTypes[0]
}

func (x ValidationError_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValidationError_Severity.Descriptor instead.
func (ValidationError_Severity) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2_validation_result_proto_rawDescGZIP(), []int{0, 0}
}

// Represents a single validation error.
type ValidationError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The severity of the error.
	Severity ValidationError_Severity `protobuf:"varint,1,opt,name=severity,proto3,enum=google.cloud.dialogflow.v2.ValidationError_Severity" json:"severity,omitempty"`
	// The names of the entries that the error is associated with.
	// Format:
	//
	// - "projects/<Project ID>/agent", if the error is associated with the entire
	// agent.
	// - "projects/<Project ID>/agent/intents/<Intent ID>", if the error is
	// associated with certain intents.
	// - "projects/<Project
	// ID>/agent/intents/<Intent Id>/trainingPhrases/<Training Phrase ID>", if the
	// error is associated with certain intent training phrases.
	// - "projects/<Project ID>/agent/intents/<Intent Id>/parameters/<Parameter
	// ID>", if the error is associated with certain intent parameters.
	// - "projects/<Project ID>/agent/entities/<Entity ID>", if the error is
	// associated with certain entities.
	Entries []string `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries,omitempty"`
	// The detailed error message.
	ErrorMessage string `protobuf:"bytes,4,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *ValidationError) Reset() {
	*x = ValidationError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_v2_validation_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationError) ProtoMessage() {}

func (x *ValidationError) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_v2_validation_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationError.ProtoReflect.Descriptor instead.
func (*ValidationError) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2_validation_result_proto_rawDescGZIP(), []int{0}
}

func (x *ValidationError) GetSeverity() ValidationError_Severity {
	if x != nil {
		return x.Severity
	}
	return ValidationError_SEVERITY_UNSPECIFIED
}

func (x *ValidationError) GetEntries() []string {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *ValidationError) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

// Represents the output of agent validation.
type ValidationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains all validation errors.
	ValidationErrors []*ValidationError `protobuf:"bytes,1,rep,name=validation_errors,json=validationErrors,proto3" json:"validation_errors,omitempty"`
}

func (x *ValidationResult) Reset() {
	*x = ValidationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_v2_validation_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationResult) ProtoMessage() {}

func (x *ValidationResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_v2_validation_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationResult.ProtoReflect.Descriptor instead.
func (*ValidationResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2_validation_result_proto_rawDescGZIP(), []int{1}
}

func (x *ValidationResult) GetValidationErrors() []*ValidationError {
	if x != nil {
		return x.ValidationErrors
	}
	return nil
}

var File_google_cloud_dialogflow_v2_validation_result_proto protoreflect.FileDescriptor

var file_google_cloud_dialogflow_v2_validation_result_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32,
	0x22, 0xf8, 0x01, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x50, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x76, 0x32, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x54, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49,
	0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x04, 0x22, 0x6c, 0x0a, 0x10, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x58, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0xa4, 0x01, 0x0a, 0x1e, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x42, 0x15, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x32,
	0x3b, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0xf8, 0x01, 0x01, 0xa2, 0x02,
	0x02, 0x44, 0x46, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x56, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dialogflow_v2_validation_result_proto_rawDescOnce sync.Once
	file_google_cloud_dialogflow_v2_validation_result_proto_rawDescData = file_google_cloud_dialogflow_v2_validation_result_proto_rawDesc
)

func file_google_cloud_dialogflow_v2_validation_result_proto_rawDescGZIP() []byte {
	file_google_cloud_dialogflow_v2_validation_result_proto_rawDescOnce.Do(func() {
		file_google_cloud_dialogflow_v2_validation_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dialogflow_v2_validation_result_proto_rawDescData)
	})
	return file_google_cloud_dialogflow_v2_validation_result_proto_rawDescData
}

var file_google_cloud_dialogflow_v2_validation_result_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_dialogflow_v2_validation_result_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_dialogflow_v2_validation_result_proto_goTypes = []interface{}{
	(ValidationError_Severity)(0), // 0: google.cloud.dialogflow.v2.ValidationError.Severity
	(*ValidationError)(nil),       // 1: google.cloud.dialogflow.v2.ValidationError
	(*ValidationResult)(nil),      // 2: google.cloud.dialogflow.v2.ValidationResult
}
var file_google_cloud_dialogflow_v2_validation_result_proto_depIdxs = []int32{
	0, // 0: google.cloud.dialogflow.v2.ValidationError.severity:type_name -> google.cloud.dialogflow.v2.ValidationError.Severity
	1, // 1: google.cloud.dialogflow.v2.ValidationResult.validation_errors:type_name -> google.cloud.dialogflow.v2.ValidationError
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_dialogflow_v2_validation_result_proto_init() }
func file_google_cloud_dialogflow_v2_validation_result_proto_init() {
	if File_google_cloud_dialogflow_v2_validation_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dialogflow_v2_validation_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationError); i {
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
		file_google_cloud_dialogflow_v2_validation_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationResult); i {
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
			RawDescriptor: file_google_cloud_dialogflow_v2_validation_result_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dialogflow_v2_validation_result_proto_goTypes,
		DependencyIndexes: file_google_cloud_dialogflow_v2_validation_result_proto_depIdxs,
		EnumInfos:         file_google_cloud_dialogflow_v2_validation_result_proto_enumTypes,
		MessageInfos:      file_google_cloud_dialogflow_v2_validation_result_proto_msgTypes,
	}.Build()
	File_google_cloud_dialogflow_v2_validation_result_proto = out.File
	file_google_cloud_dialogflow_v2_validation_result_proto_rawDesc = nil
	file_google_cloud_dialogflow_v2_validation_result_proto_goTypes = nil
	file_google_cloud_dialogflow_v2_validation_result_proto_depIdxs = nil
}
