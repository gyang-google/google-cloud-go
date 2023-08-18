// Copyright 2023 Google LLC
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
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: google/cloud/dialogflow/cx/v3beta1/import_strategy.proto

package cxpb

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

// Import strategies for the conflict resolution of resources (i.e. intents,
// entities, and webhooks) with identical display names during import
// operations.
type ImportStrategy int32

const (
	// Unspecified. Treated as 'CREATE_NEW'.
	ImportStrategy_IMPORT_STRATEGY_UNSPECIFIED ImportStrategy = 0
	// Create a new resource with a numeric suffix appended to the end of the
	// existing display name.
	ImportStrategy_IMPORT_STRATEGY_CREATE_NEW ImportStrategy = 1
	// Replace existing resource with incoming resource in the content to be
	// imported.
	ImportStrategy_IMPORT_STRATEGY_REPLACE ImportStrategy = 2
	// Keep existing resource and discard incoming resource in the content to be
	// imported.
	ImportStrategy_IMPORT_STRATEGY_KEEP ImportStrategy = 3
	// Combine existing and incoming resources when a conflict is encountered.
	ImportStrategy_IMPORT_STRATEGY_MERGE ImportStrategy = 4
	// Throw error if a conflict is encountered.
	ImportStrategy_IMPORT_STRATEGY_THROW_ERROR ImportStrategy = 5
)

// Enum value maps for ImportStrategy.
var (
	ImportStrategy_name = map[int32]string{
		0: "IMPORT_STRATEGY_UNSPECIFIED",
		1: "IMPORT_STRATEGY_CREATE_NEW",
		2: "IMPORT_STRATEGY_REPLACE",
		3: "IMPORT_STRATEGY_KEEP",
		4: "IMPORT_STRATEGY_MERGE",
		5: "IMPORT_STRATEGY_THROW_ERROR",
	}
	ImportStrategy_value = map[string]int32{
		"IMPORT_STRATEGY_UNSPECIFIED": 0,
		"IMPORT_STRATEGY_CREATE_NEW":  1,
		"IMPORT_STRATEGY_REPLACE":     2,
		"IMPORT_STRATEGY_KEEP":        3,
		"IMPORT_STRATEGY_MERGE":       4,
		"IMPORT_STRATEGY_THROW_ERROR": 5,
	}
)

func (x ImportStrategy) Enum() *ImportStrategy {
	p := new(ImportStrategy)
	*p = x
	return p
}

func (x ImportStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImportStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_enumTypes[0].Descriptor()
}

func (ImportStrategy) Type() protoreflect.EnumType {
	return &file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_enumTypes[0]
}

func (x ImportStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImportStrategy.Descriptor instead.
func (ImportStrategy) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_rawDescGZIP(), []int{0}
}

var File_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto protoreflect.FileDescriptor

var file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x76, 0x33, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2a, 0xc4,
	0x01, 0x0a, 0x0e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x52, 0x41,
	0x54, 0x45, 0x47, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x52,
	0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x45, 0x57,
	0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x52,
	0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x10, 0x02, 0x12,
	0x18, 0x0a, 0x14, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45,
	0x47, 0x59, 0x5f, 0x4b, 0x45, 0x45, 0x50, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4d, 0x50,
	0x4f, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x4d, 0x45, 0x52,
	0x47, 0x45, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x53,
	0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x54, 0x48, 0x52, 0x4f, 0x57, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x05, 0x42, 0xa4, 0x01, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x42, 0x13, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x33,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x78, 0x70, 0x62, 0x3b, 0x63, 0x78, 0x70, 0x62, 0xf8,
	0x01, 0x01, 0xa2, 0x02, 0x02, 0x44, 0x46, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x43, 0x78, 0x2e, 0x56, 0x33, 0x42, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_rawDescOnce sync.Once
	file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_rawDescData = file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_rawDesc
)

func file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_rawDescGZIP() []byte {
	file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_rawDescOnce.Do(func() {
		file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_rawDescData)
	})
	return file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_rawDescData
}

var file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_goTypes = []interface{}{
	(ImportStrategy)(0), // 0: google.cloud.dialogflow.cx.v3beta1.ImportStrategy
}
var file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_init() }
func file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_init() {
	if File_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_goTypes,
		DependencyIndexes: file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_depIdxs,
		EnumInfos:         file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_enumTypes,
	}.Build()
	File_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto = out.File
	file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_rawDesc = nil
	file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_goTypes = nil
	file_google_cloud_dialogflow_cx_v3beta1_import_strategy_proto_depIdxs = nil
}
