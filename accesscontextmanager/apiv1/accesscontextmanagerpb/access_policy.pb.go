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
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: google/identity/accesscontextmanager/v1/access_policy.proto

package accesscontextmanagerpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// `AccessPolicy` is a container for `AccessLevels` (which define the necessary
// attributes to use Google Cloud services) and `ServicePerimeters` (which
// define regions of services able to freely pass data within a perimeter). An
// access policy is globally visible within an organization, and the
// restrictions it specifies apply to all projects within an organization.
type AccessPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Resource name of the `AccessPolicy`. Format:
	// `accessPolicies/{access_policy}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The parent of this `AccessPolicy` in the Cloud Resource
	// Hierarchy. Currently immutable once created. Format:
	// `organizations/{organization_id}`
	Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. Human readable title. Does not affect behavior.
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// The scopes of a policy define which resources an ACM policy can restrict,
	// and where ACM resources can be referenced.
	// For example, a policy with scopes=["folders/123"] has the following
	// behavior:
	// - vpcsc perimeters can only restrict projects within folders/123
	// - access levels can only be referenced by resources within folders/123.
	// If empty, there are no limitations on which resources can be restricted by
	// an ACM policy, and there are no limitations on where ACM resources can be
	// referenced.
	// Only one policy can include a given scope (attempting to create a second
	// policy which includes "folders/123" will result in an error).
	// Currently, scopes cannot be modified after a policy is created.
	// Currently, policies can only have a single scope.
	// Format: list of `folders/{folder_number}` or `projects/{project_number}`
	Scopes []string `protobuf:"bytes,7,rep,name=scopes,proto3" json:"scopes,omitempty"`
	// Output only. Time the `AccessPolicy` was created in UTC.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Time the `AccessPolicy` was updated in UTC.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. An opaque identifier for the current version of the
	// `AccessPolicy`. This will always be a strongly validated etag, meaning that
	// two Access Polices will be identical if and only if their etags are
	// identical. Clients should not expect this to be in any specific format.
	Etag string `protobuf:"bytes,6,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *AccessPolicy) Reset() {
	*x = AccessPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_identity_accesscontextmanager_v1_access_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessPolicy) ProtoMessage() {}

func (x *AccessPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_google_identity_accesscontextmanager_v1_access_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessPolicy.ProtoReflect.Descriptor instead.
func (*AccessPolicy) Descriptor() ([]byte, []int) {
	return file_google_identity_accesscontextmanager_v1_access_policy_proto_rawDescGZIP(), []int{0}
}

func (x *AccessPolicy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccessPolicy) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *AccessPolicy) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AccessPolicy) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *AccessPolicy) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *AccessPolicy) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *AccessPolicy) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

var File_google_identity_accesscontextmanager_v1_access_policy_proto protoreflect.FileDescriptor

var file_google_identity_accesscontextmanager_v1_access_policy_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcd, 0x02, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x3b, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x3a, 0x55, 0xea, 0x41, 0x52,
	0x0a, 0x30, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x1e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x7d, 0x42, 0xa2, 0x02, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x42, 0x0b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x5c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x62, 0xa2,
	0x02, 0x04, 0x47, 0x41, 0x43, 0x4d, 0xaa, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x2a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_identity_accesscontextmanager_v1_access_policy_proto_rawDescOnce sync.Once
	file_google_identity_accesscontextmanager_v1_access_policy_proto_rawDescData = file_google_identity_accesscontextmanager_v1_access_policy_proto_rawDesc
)

func file_google_identity_accesscontextmanager_v1_access_policy_proto_rawDescGZIP() []byte {
	file_google_identity_accesscontextmanager_v1_access_policy_proto_rawDescOnce.Do(func() {
		file_google_identity_accesscontextmanager_v1_access_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_identity_accesscontextmanager_v1_access_policy_proto_rawDescData)
	})
	return file_google_identity_accesscontextmanager_v1_access_policy_proto_rawDescData
}

var file_google_identity_accesscontextmanager_v1_access_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_identity_accesscontextmanager_v1_access_policy_proto_goTypes = []interface{}{
	(*AccessPolicy)(nil),          // 0: google.identity.accesscontextmanager.v1.AccessPolicy
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_google_identity_accesscontextmanager_v1_access_policy_proto_depIdxs = []int32{
	1, // 0: google.identity.accesscontextmanager.v1.AccessPolicy.create_time:type_name -> google.protobuf.Timestamp
	1, // 1: google.identity.accesscontextmanager.v1.AccessPolicy.update_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_identity_accesscontextmanager_v1_access_policy_proto_init() }
func file_google_identity_accesscontextmanager_v1_access_policy_proto_init() {
	if File_google_identity_accesscontextmanager_v1_access_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_identity_accesscontextmanager_v1_access_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessPolicy); i {
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
			RawDescriptor: file_google_identity_accesscontextmanager_v1_access_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_identity_accesscontextmanager_v1_access_policy_proto_goTypes,
		DependencyIndexes: file_google_identity_accesscontextmanager_v1_access_policy_proto_depIdxs,
		MessageInfos:      file_google_identity_accesscontextmanager_v1_access_policy_proto_msgTypes,
	}.Build()
	File_google_identity_accesscontextmanager_v1_access_policy_proto = out.File
	file_google_identity_accesscontextmanager_v1_access_policy_proto_rawDesc = nil
	file_google_identity_accesscontextmanager_v1_access_policy_proto_goTypes = nil
	file_google_identity_accesscontextmanager_v1_access_policy_proto_depIdxs = nil
}
