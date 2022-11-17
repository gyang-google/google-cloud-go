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
// source: google/cloud/aiplatform/v1beta1/tensorboard_data.proto

package aiplatformpb

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

// All the data stored in a TensorboardTimeSeries.
type TimeSeriesData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the TensorboardTimeSeries, which will become the final component
	// of the TensorboardTimeSeries' resource name
	TensorboardTimeSeriesId string `protobuf:"bytes,1,opt,name=tensorboard_time_series_id,json=tensorboardTimeSeriesId,proto3" json:"tensorboard_time_series_id,omitempty"`
	// Required. Immutable. The value type of this time series. All the values in this time series data
	// must match this value type.
	ValueType TensorboardTimeSeries_ValueType `protobuf:"varint,2,opt,name=value_type,json=valueType,proto3,enum=google.cloud.aiplatform.v1beta1.TensorboardTimeSeries_ValueType" json:"value_type,omitempty"`
	// Required. Data points in this time series.
	Values []*TimeSeriesDataPoint `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *TimeSeriesData) Reset() {
	*x = TimeSeriesData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeSeriesData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeriesData) ProtoMessage() {}

func (x *TimeSeriesData) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeriesData.ProtoReflect.Descriptor instead.
func (*TimeSeriesData) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDescGZIP(), []int{0}
}

func (x *TimeSeriesData) GetTensorboardTimeSeriesId() string {
	if x != nil {
		return x.TensorboardTimeSeriesId
	}
	return ""
}

func (x *TimeSeriesData) GetValueType() TensorboardTimeSeries_ValueType {
	if x != nil {
		return x.ValueType
	}
	return TensorboardTimeSeries_VALUE_TYPE_UNSPECIFIED
}

func (x *TimeSeriesData) GetValues() []*TimeSeriesDataPoint {
	if x != nil {
		return x.Values
	}
	return nil
}

// A TensorboardTimeSeries data point.
type TimeSeriesDataPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value of this time series data point.
	//
	// Types that are assignable to Value:
	//
	//	*TimeSeriesDataPoint_Scalar
	//	*TimeSeriesDataPoint_Tensor
	//	*TimeSeriesDataPoint_Blobs
	Value isTimeSeriesDataPoint_Value `protobuf_oneof:"value"`
	// Wall clock timestamp when this data point is generated by the end user.
	WallTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=wall_time,json=wallTime,proto3" json:"wall_time,omitempty"`
	// Step index of this data point within the run.
	Step int64 `protobuf:"varint,2,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *TimeSeriesDataPoint) Reset() {
	*x = TimeSeriesDataPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeSeriesDataPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeriesDataPoint) ProtoMessage() {}

func (x *TimeSeriesDataPoint) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeriesDataPoint.ProtoReflect.Descriptor instead.
func (*TimeSeriesDataPoint) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDescGZIP(), []int{1}
}

func (m *TimeSeriesDataPoint) GetValue() isTimeSeriesDataPoint_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *TimeSeriesDataPoint) GetScalar() *Scalar {
	if x, ok := x.GetValue().(*TimeSeriesDataPoint_Scalar); ok {
		return x.Scalar
	}
	return nil
}

func (x *TimeSeriesDataPoint) GetTensor() *TensorboardTensor {
	if x, ok := x.GetValue().(*TimeSeriesDataPoint_Tensor); ok {
		return x.Tensor
	}
	return nil
}

func (x *TimeSeriesDataPoint) GetBlobs() *TensorboardBlobSequence {
	if x, ok := x.GetValue().(*TimeSeriesDataPoint_Blobs); ok {
		return x.Blobs
	}
	return nil
}

func (x *TimeSeriesDataPoint) GetWallTime() *timestamppb.Timestamp {
	if x != nil {
		return x.WallTime
	}
	return nil
}

func (x *TimeSeriesDataPoint) GetStep() int64 {
	if x != nil {
		return x.Step
	}
	return 0
}

type isTimeSeriesDataPoint_Value interface {
	isTimeSeriesDataPoint_Value()
}

type TimeSeriesDataPoint_Scalar struct {
	// A scalar value.
	Scalar *Scalar `protobuf:"bytes,3,opt,name=scalar,proto3,oneof"`
}

type TimeSeriesDataPoint_Tensor struct {
	// A tensor value.
	Tensor *TensorboardTensor `protobuf:"bytes,4,opt,name=tensor,proto3,oneof"`
}

type TimeSeriesDataPoint_Blobs struct {
	// A blob sequence value.
	Blobs *TensorboardBlobSequence `protobuf:"bytes,5,opt,name=blobs,proto3,oneof"`
}

func (*TimeSeriesDataPoint_Scalar) isTimeSeriesDataPoint_Value() {}

func (*TimeSeriesDataPoint_Tensor) isTimeSeriesDataPoint_Value() {}

func (*TimeSeriesDataPoint_Blobs) isTimeSeriesDataPoint_Value() {}

// One point viewable on a scalar metric plot.
type Scalar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value of the point at this step / timestamp.
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Scalar) Reset() {
	*x = Scalar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scalar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scalar) ProtoMessage() {}

func (x *Scalar) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scalar.ProtoReflect.Descriptor instead.
func (*Scalar) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDescGZIP(), []int{2}
}

func (x *Scalar) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// One point viewable on a tensor metric plot.
type TensorboardTensor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Serialized form of
	// https://github.com/tensorflow/tensorflow/blob/master/tensorflow/core/framework/tensor.proto
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// Optional. Version number of TensorProto used to serialize [value][google.cloud.aiplatform.v1beta1.TensorboardTensor.value].
	VersionNumber int32 `protobuf:"varint,2,opt,name=version_number,json=versionNumber,proto3" json:"version_number,omitempty"`
}

func (x *TensorboardTensor) Reset() {
	*x = TensorboardTensor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TensorboardTensor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TensorboardTensor) ProtoMessage() {}

func (x *TensorboardTensor) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TensorboardTensor.ProtoReflect.Descriptor instead.
func (*TensorboardTensor) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDescGZIP(), []int{3}
}

func (x *TensorboardTensor) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *TensorboardTensor) GetVersionNumber() int32 {
	if x != nil {
		return x.VersionNumber
	}
	return 0
}

// One point viewable on a blob metric plot, but mostly just a wrapper message
// to work around repeated fields can't be used directly within `oneof` fields.
type TensorboardBlobSequence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of blobs contained within the sequence.
	Values []*TensorboardBlob `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *TensorboardBlobSequence) Reset() {
	*x = TensorboardBlobSequence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TensorboardBlobSequence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TensorboardBlobSequence) ProtoMessage() {}

func (x *TensorboardBlobSequence) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TensorboardBlobSequence.ProtoReflect.Descriptor instead.
func (*TensorboardBlobSequence) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDescGZIP(), []int{4}
}

func (x *TensorboardBlobSequence) GetValues() []*TensorboardBlob {
	if x != nil {
		return x.Values
	}
	return nil
}

// One blob (e.g, image, graph) viewable on a blob metric plot.
type TensorboardBlob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. A URI safe key uniquely identifying a blob. Can be used to locate the blob
	// stored in the Cloud Storage bucket of the consumer project.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Optional. The bytes of the blob is not present unless it's returned by the
	// ReadTensorboardBlobData endpoint.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TensorboardBlob) Reset() {
	*x = TensorboardBlob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TensorboardBlob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TensorboardBlob) ProtoMessage() {}

func (x *TensorboardBlob) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TensorboardBlob.ProtoReflect.Descriptor instead.
func (*TensorboardBlob) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDescGZIP(), []int{5}
}

func (x *TensorboardBlob) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TensorboardBlob) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_google_cloud_aiplatform_v1beta1_tensorboard_data_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x02, 0x0a, 0x0e, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a,
	0x1a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x17, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x64, 0x12,
	0x67, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x09, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x51, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xce, 0x02, 0x0a, 0x13,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x06, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x48, 0x00, 0x52, 0x06,
	0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x12, 0x4c, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x06, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x12, 0x50, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x42, 0x6c, 0x6f, 0x62, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x48, 0x00, 0x52,
	0x05, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x74, 0x65, 0x70, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x06,
	0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5a, 0x0a, 0x11,
	0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x54, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2a, 0x0a, 0x0e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x63, 0x0a, 0x17, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x53, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x3f, 0x0a,
	0x0f, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x42, 0x6c, 0x6f, 0x62,
	0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0xf1,
	0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x14, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41,
	0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_goTypes = []interface{}{
	(*TimeSeriesData)(nil),               // 0: google.cloud.aiplatform.v1beta1.TimeSeriesData
	(*TimeSeriesDataPoint)(nil),          // 1: google.cloud.aiplatform.v1beta1.TimeSeriesDataPoint
	(*Scalar)(nil),                       // 2: google.cloud.aiplatform.v1beta1.Scalar
	(*TensorboardTensor)(nil),            // 3: google.cloud.aiplatform.v1beta1.TensorboardTensor
	(*TensorboardBlobSequence)(nil),      // 4: google.cloud.aiplatform.v1beta1.TensorboardBlobSequence
	(*TensorboardBlob)(nil),              // 5: google.cloud.aiplatform.v1beta1.TensorboardBlob
	(TensorboardTimeSeries_ValueType)(0), // 6: google.cloud.aiplatform.v1beta1.TensorboardTimeSeries.ValueType
	(*timestamppb.Timestamp)(nil),        // 7: google.protobuf.Timestamp
}
var file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_depIdxs = []int32{
	6, // 0: google.cloud.aiplatform.v1beta1.TimeSeriesData.value_type:type_name -> google.cloud.aiplatform.v1beta1.TensorboardTimeSeries.ValueType
	1, // 1: google.cloud.aiplatform.v1beta1.TimeSeriesData.values:type_name -> google.cloud.aiplatform.v1beta1.TimeSeriesDataPoint
	2, // 2: google.cloud.aiplatform.v1beta1.TimeSeriesDataPoint.scalar:type_name -> google.cloud.aiplatform.v1beta1.Scalar
	3, // 3: google.cloud.aiplatform.v1beta1.TimeSeriesDataPoint.tensor:type_name -> google.cloud.aiplatform.v1beta1.TensorboardTensor
	4, // 4: google.cloud.aiplatform.v1beta1.TimeSeriesDataPoint.blobs:type_name -> google.cloud.aiplatform.v1beta1.TensorboardBlobSequence
	7, // 5: google.cloud.aiplatform.v1beta1.TimeSeriesDataPoint.wall_time:type_name -> google.protobuf.Timestamp
	5, // 6: google.cloud.aiplatform.v1beta1.TensorboardBlobSequence.values:type_name -> google.cloud.aiplatform.v1beta1.TensorboardBlob
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_init() }
func file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_tensorboard_data_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1beta1_tensorboard_time_series_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSeriesData); i {
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
		file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSeriesDataPoint); i {
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
		file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scalar); i {
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
		file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TensorboardTensor); i {
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
		file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TensorboardBlobSequence); i {
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
		file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TensorboardBlob); i {
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
	file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*TimeSeriesDataPoint_Scalar)(nil),
		(*TimeSeriesDataPoint_Tensor)(nil),
		(*TimeSeriesDataPoint_Blobs)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_tensorboard_data_proto = out.File
	file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_tensorboard_data_proto_depIdxs = nil
}
