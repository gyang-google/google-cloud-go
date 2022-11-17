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
// source: google/cloud/retail/v2beta/model.proto

package retailpb

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

// The serving state of the model.
type Model_ServingState int32

const (
	// Unspecified serving state.
	Model_SERVING_STATE_UNSPECIFIED Model_ServingState = 0
	// The model is not serving.
	Model_INACTIVE Model_ServingState = 1
	// The model is serving and can be queried.
	Model_ACTIVE Model_ServingState = 2
	// The model is trained on tuned hyperparameters and can be
	// queried.
	Model_TUNED Model_ServingState = 3
)

// Enum value maps for Model_ServingState.
var (
	Model_ServingState_name = map[int32]string{
		0: "SERVING_STATE_UNSPECIFIED",
		1: "INACTIVE",
		2: "ACTIVE",
		3: "TUNED",
	}
	Model_ServingState_value = map[string]int32{
		"SERVING_STATE_UNSPECIFIED": 0,
		"INACTIVE":                  1,
		"ACTIVE":                    2,
		"TUNED":                     3,
	}
)

func (x Model_ServingState) Enum() *Model_ServingState {
	p := new(Model_ServingState)
	*p = x
	return p
}

func (x Model_ServingState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Model_ServingState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_retail_v2beta_model_proto_enumTypes[0].Descriptor()
}

func (Model_ServingState) Type() protoreflect.EnumType {
	return &file_google_cloud_retail_v2beta_model_proto_enumTypes[0]
}

func (x Model_ServingState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Model_ServingState.Descriptor instead.
func (Model_ServingState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2beta_model_proto_rawDescGZIP(), []int{0, 0}
}

// The training state of the model.
type Model_TrainingState int32

const (
	// Unspecified training state.
	Model_TRAINING_STATE_UNSPECIFIED Model_TrainingState = 0
	// The model training is paused.
	Model_PAUSED Model_TrainingState = 1
	// The model is training.
	Model_TRAINING Model_TrainingState = 2
)

// Enum value maps for Model_TrainingState.
var (
	Model_TrainingState_name = map[int32]string{
		0: "TRAINING_STATE_UNSPECIFIED",
		1: "PAUSED",
		2: "TRAINING",
	}
	Model_TrainingState_value = map[string]int32{
		"TRAINING_STATE_UNSPECIFIED": 0,
		"PAUSED":                     1,
		"TRAINING":                   2,
	}
)

func (x Model_TrainingState) Enum() *Model_TrainingState {
	p := new(Model_TrainingState)
	*p = x
	return p
}

func (x Model_TrainingState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Model_TrainingState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_retail_v2beta_model_proto_enumTypes[1].Descriptor()
}

func (Model_TrainingState) Type() protoreflect.EnumType {
	return &file_google_cloud_retail_v2beta_model_proto_enumTypes[1]
}

func (x Model_TrainingState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Model_TrainingState.Descriptor instead.
func (Model_TrainingState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2beta_model_proto_rawDescGZIP(), []int{0, 1}
}

// Describes whether periodic tuning is enabled for this model
// or not. Periodic tuning is scheduled at most every three months. You can
// start a tuning process manually by using the `TuneModel`
// method, which starts a tuning process immediately and resets the quarterly
// schedule. Enabling or disabling periodic tuning does not affect any
// current tuning processes.
type Model_PeriodicTuningState int32

const (
	// Unspecified default value, should never be explicitly set.
	Model_PERIODIC_TUNING_STATE_UNSPECIFIED Model_PeriodicTuningState = 0
	// The model has periodic tuning disabled. Tuning
	// can be reenabled by calling the `EnableModelPeriodicTuning`
	// method or by calling the `TuneModel` method.
	Model_PERIODIC_TUNING_DISABLED Model_PeriodicTuningState = 1
	// The model cannot be tuned with periodic tuning OR the
	// `TuneModel` method. Hide the options in customer UI and
	// reject any requests through the backend self serve API.
	Model_ALL_TUNING_DISABLED Model_PeriodicTuningState = 3
	// The model has periodic tuning enabled. Tuning
	// can be disabled by calling the `DisableModelPeriodicTuning`
	// method.
	Model_PERIODIC_TUNING_ENABLED Model_PeriodicTuningState = 2
)

// Enum value maps for Model_PeriodicTuningState.
var (
	Model_PeriodicTuningState_name = map[int32]string{
		0: "PERIODIC_TUNING_STATE_UNSPECIFIED",
		1: "PERIODIC_TUNING_DISABLED",
		3: "ALL_TUNING_DISABLED",
		2: "PERIODIC_TUNING_ENABLED",
	}
	Model_PeriodicTuningState_value = map[string]int32{
		"PERIODIC_TUNING_STATE_UNSPECIFIED": 0,
		"PERIODIC_TUNING_DISABLED":          1,
		"ALL_TUNING_DISABLED":               3,
		"PERIODIC_TUNING_ENABLED":           2,
	}
)

func (x Model_PeriodicTuningState) Enum() *Model_PeriodicTuningState {
	p := new(Model_PeriodicTuningState)
	*p = x
	return p
}

func (x Model_PeriodicTuningState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Model_PeriodicTuningState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_retail_v2beta_model_proto_enumTypes[2].Descriptor()
}

func (Model_PeriodicTuningState) Type() protoreflect.EnumType {
	return &file_google_cloud_retail_v2beta_model_proto_enumTypes[2]
}

func (x Model_PeriodicTuningState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Model_PeriodicTuningState.Descriptor instead.
func (Model_PeriodicTuningState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2beta_model_proto_rawDescGZIP(), []int{0, 2}
}

// Describes whether this model have sufficient training data
// to be continuously trained.
type Model_DataState int32

const (
	// Unspecified default value, should never be explicitly set.
	Model_DATA_STATE_UNSPECIFIED Model_DataState = 0
	// The model has sufficient training data.
	Model_DATA_OK Model_DataState = 1
	// The model does not have sufficient training data. Error
	// messages can be queried via Stackdriver.
	Model_DATA_ERROR Model_DataState = 2
)

// Enum value maps for Model_DataState.
var (
	Model_DataState_name = map[int32]string{
		0: "DATA_STATE_UNSPECIFIED",
		1: "DATA_OK",
		2: "DATA_ERROR",
	}
	Model_DataState_value = map[string]int32{
		"DATA_STATE_UNSPECIFIED": 0,
		"DATA_OK":                1,
		"DATA_ERROR":             2,
	}
)

func (x Model_DataState) Enum() *Model_DataState {
	p := new(Model_DataState)
	*p = x
	return p
}

func (x Model_DataState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Model_DataState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_retail_v2beta_model_proto_enumTypes[3].Descriptor()
}

func (Model_DataState) Type() protoreflect.EnumType {
	return &file_google_cloud_retail_v2beta_model_proto_enumTypes[3]
}

func (x Model_DataState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Model_DataState.Descriptor instead.
func (Model_DataState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2beta_model_proto_rawDescGZIP(), []int{0, 3}
}

// Metadata that describes the training and serving parameters of a
// [Model][google.cloud.retail.v2beta.Model]. A
// [Model][google.cloud.retail.v2beta.Model] can be associated with a
// [ServingConfig][google.cloud.retail.v2beta.ServingConfig] and then queried
// through the Predict API.
type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The fully qualified resource name of the model.
	//
	// Format:
	// `projects/{project_number}/locations/{location_id}/catalogs/{catalog_id}/models/{model_id}`
	// catalog_id has char limit of 50.
	// recommendation_model_id has char limit of 40.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The display name of the model.
	//
	// Should be human readable, used to display Recommendation Models in the
	// Retail Cloud Console Dashboard. UTF-8 encoded string with limit of 1024
	// characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. The training state that the model is in (e.g.
	// `TRAINING` or `PAUSED`).
	//
	// Since part of the cost of running the service
	// is frequency of training - this can be used to determine when to train
	// model in order to control cost. If not specified: the default value for
	// `CreateModel` method is `TRAINING`. The default value for
	// `UpdateModel` method is to keep the state the same as before.
	TrainingState Model_TrainingState `protobuf:"varint,3,opt,name=training_state,json=trainingState,proto3,enum=google.cloud.retail.v2beta.Model_TrainingState" json:"training_state,omitempty"`
	// Output only. The serving state of the model: `ACTIVE`, `NOT_ACTIVE`.
	ServingState Model_ServingState `protobuf:"varint,4,opt,name=serving_state,json=servingState,proto3,enum=google.cloud.retail.v2beta.Model_ServingState" json:"serving_state,omitempty"`
	// Output only. Timestamp the Recommendation Model was created at.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp the Recommendation Model was last updated. E.g.
	// if a Recommendation Model was paused - this would be the time the pause was
	// initiated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Required. The type of model e.g. `home-page`.
	//
	// Currently supported values: `recommended-for-you`, `others-you-may-like`,
	// `frequently-bought-together`, `page-optimization`, `similar-items`,
	// `buy-it-again`, and `recently-viewed`(readonly value).
	//
	// This field together with
	// [optimization_objective][google.cloud.retail.v2beta.Model.optimization_objective]
	// describe model metadata to use to control model training and serving.
	// See https://cloud.google.com/retail/docs/models
	// for more details on what the model metadata control and which combination
	// of parameters are valid. For invalid combinations of parameters (e.g. type
	// = `frequently-bought-together` and optimization_objective = `ctr`), you
	// receive an error 400 if you try to create/update a recommendation with
	// this set of knobs.
	Type string `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	// Optional. The optimization objective e.g. `cvr`.
	//
	// Currently supported
	// values: `ctr`, `cvr`, `revenue-per-order`.
	//
	//	If not specified, we choose default based on model type.
	//
	// Default depends on type of recommendation:
	//
	// `recommended-for-you` => `ctr`
	//
	// `others-you-may-like` => `ctr`
	//
	// `frequently-bought-together` => `revenue_per_order`
	//
	// This field together with
	// [optimization_objective][google.cloud.retail.v2beta.Model.type]
	// describe model metadata to use to control model training and serving.
	// See https://cloud.google.com/retail/docs/models
	// for more details on what the model metadata control and which combination
	// of parameters are valid. For invalid combinations of parameters (e.g. type
	// = `frequently-bought-together` and optimization_objective = `ctr`), you
	// receive an error 400 if you try to create/update a recommendation with
	// this set of knobs.
	OptimizationObjective string `protobuf:"bytes,8,opt,name=optimization_objective,json=optimizationObjective,proto3" json:"optimization_objective,omitempty"`
	// Optional. The state of periodic tuning.
	//
	// The period we use is 3 months - to do a
	// one-off tune earlier use the `TuneModel` method. Default value
	// is `PERIODIC_TUNING_ENABLED`.
	PeriodicTuningState Model_PeriodicTuningState `protobuf:"varint,11,opt,name=periodic_tuning_state,json=periodicTuningState,proto3,enum=google.cloud.retail.v2beta.Model_PeriodicTuningState" json:"periodic_tuning_state,omitempty"`
	// Output only. The timestamp when the latest successful tune finished.
	LastTuneTime *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=last_tune_time,json=lastTuneTime,proto3" json:"last_tune_time,omitempty"`
	// Output only. The tune operation associated with the model.
	//
	// Can be used to determine if there is an ongoing tune for this
	// recommendation. Empty field implies no tune is goig on.
	TuningOperation string `protobuf:"bytes,15,opt,name=tuning_operation,json=tuningOperation,proto3" json:"tuning_operation,omitempty"`
	// Output only. The state of data requirements for this model: `DATA_OK` and
	// `DATA_ERROR`.
	//
	// Recommendation model cannot be trained if the data is in
	// `DATA_ERROR` state. Recommendation model can have `DATA_ERROR` state even
	// if serving state is `ACTIVE`: models were trained successfully before, but
	// cannot be refreshed because model no longer has sufficient
	// data for training.
	DataState Model_DataState `protobuf:"varint,16,opt,name=data_state,json=dataState,proto3,enum=google.cloud.retail.v2beta.Model_DataState" json:"data_state,omitempty"`
	// Optional. If `RECOMMENDATIONS_FILTERING_ENABLED`, recommendation filtering
	// by attributes is enabled for the model.
	FilteringOption RecommendationsFilteringOption `protobuf:"varint,18,opt,name=filtering_option,json=filteringOption,proto3,enum=google.cloud.retail.v2beta.RecommendationsFilteringOption" json:"filtering_option,omitempty"`
	// Output only. The list of valid serving configs associated with the
	// PageOptimizationConfig.
	ServingConfigLists []*Model_ServingConfigList `protobuf:"bytes,19,rep,name=serving_config_lists,json=servingConfigLists,proto3" json:"serving_config_lists,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2beta_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2beta_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2beta_model_proto_rawDescGZIP(), []int{0}
}

func (x *Model) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Model) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Model) GetTrainingState() Model_TrainingState {
	if x != nil {
		return x.TrainingState
	}
	return Model_TRAINING_STATE_UNSPECIFIED
}

func (x *Model) GetServingState() Model_ServingState {
	if x != nil {
		return x.ServingState
	}
	return Model_SERVING_STATE_UNSPECIFIED
}

func (x *Model) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Model) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Model) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Model) GetOptimizationObjective() string {
	if x != nil {
		return x.OptimizationObjective
	}
	return ""
}

func (x *Model) GetPeriodicTuningState() Model_PeriodicTuningState {
	if x != nil {
		return x.PeriodicTuningState
	}
	return Model_PERIODIC_TUNING_STATE_UNSPECIFIED
}

func (x *Model) GetLastTuneTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastTuneTime
	}
	return nil
}

func (x *Model) GetTuningOperation() string {
	if x != nil {
		return x.TuningOperation
	}
	return ""
}

func (x *Model) GetDataState() Model_DataState {
	if x != nil {
		return x.DataState
	}
	return Model_DATA_STATE_UNSPECIFIED
}

func (x *Model) GetFilteringOption() RecommendationsFilteringOption {
	if x != nil {
		return x.FilteringOption
	}
	return RecommendationsFilteringOption_RECOMMENDATIONS_FILTERING_OPTION_UNSPECIFIED
}

func (x *Model) GetServingConfigLists() []*Model_ServingConfigList {
	if x != nil {
		return x.ServingConfigLists
	}
	return nil
}

// Represents an ordered combination of valid serving configs, which
// can be used for `PAGE_OPTIMIZATION` recommendations.
type Model_ServingConfigList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. A set of valid serving configs that may be used for
	// `PAGE_OPTIMIZATION`.
	ServingConfigIds []string `protobuf:"bytes,1,rep,name=serving_config_ids,json=servingConfigIds,proto3" json:"serving_config_ids,omitempty"`
}

func (x *Model_ServingConfigList) Reset() {
	*x = Model_ServingConfigList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2beta_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model_ServingConfigList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model_ServingConfigList) ProtoMessage() {}

func (x *Model_ServingConfigList) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2beta_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model_ServingConfigList.ProtoReflect.Descriptor instead.
func (*Model_ServingConfigList) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2beta_model_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Model_ServingConfigList) GetServingConfigIds() []string {
	if x != nil {
		return x.ServingConfigIds
	}
	return nil
}

var File_google_cloud_retail_v2beta_model_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2beta_model_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x0c, 0x0a, 0x05, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x17, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x16, 0x6f, 0x70, 0x74,
	0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x15,
	0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x6e, 0x0a, 0x15, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69,
	0x63, 0x5f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63,
	0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x52, 0x13, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x75,
	0x6e, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x54, 0x75, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x10,
	0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x74, 0x75, 0x6e,
	0x69, 0x6e, 0x67, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x0a,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x6a, 0x0a,
	0x10, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6a, 0x0a, 0x14, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x73, 0x1a, 0x46, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x12, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x10, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x73, 0x22, 0x52, 0x0a,
	0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a,
	0x19, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x55, 0x4e, 0x45, 0x44, 0x10,
	0x03, 0x22, 0x49, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x22, 0x90, 0x01, 0x0a,
	0x13, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x21, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x49, 0x43,
	0x5f, 0x54, 0x55, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x50,
	0x45, 0x52, 0x49, 0x4f, 0x44, 0x49, 0x43, 0x5f, 0x54, 0x55, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x44,
	0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x4c, 0x4c,
	0x5f, 0x54, 0x55, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44,
	0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x49, 0x43, 0x5f, 0x54,
	0x55, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x22,
	0x44, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x16,
	0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x41, 0x54, 0x41,
	0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x02, 0x3a, 0x6b, 0xea, 0x41, 0x68, 0x0a, 0x1b, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x49, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x7d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x7b, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x7d, 0x42, 0xd3, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x3b, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0xa2, 0x02, 0x06, 0x52, 0x45, 0x54, 0x41, 0x49, 0x4c, 0xaa, 0x02,
	0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x56, 0x32, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x1a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_retail_v2beta_model_proto_rawDescOnce sync.Once
	file_google_cloud_retail_v2beta_model_proto_rawDescData = file_google_cloud_retail_v2beta_model_proto_rawDesc
)

func file_google_cloud_retail_v2beta_model_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_v2beta_model_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_v2beta_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_v2beta_model_proto_rawDescData)
	})
	return file_google_cloud_retail_v2beta_model_proto_rawDescData
}

var file_google_cloud_retail_v2beta_model_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_google_cloud_retail_v2beta_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_retail_v2beta_model_proto_goTypes = []interface{}{
	(Model_ServingState)(0),             // 0: google.cloud.retail.v2beta.Model.ServingState
	(Model_TrainingState)(0),            // 1: google.cloud.retail.v2beta.Model.TrainingState
	(Model_PeriodicTuningState)(0),      // 2: google.cloud.retail.v2beta.Model.PeriodicTuningState
	(Model_DataState)(0),                // 3: google.cloud.retail.v2beta.Model.DataState
	(*Model)(nil),                       // 4: google.cloud.retail.v2beta.Model
	(*Model_ServingConfigList)(nil),     // 5: google.cloud.retail.v2beta.Model.ServingConfigList
	(*timestamppb.Timestamp)(nil),       // 6: google.protobuf.Timestamp
	(RecommendationsFilteringOption)(0), // 7: google.cloud.retail.v2beta.RecommendationsFilteringOption
}
var file_google_cloud_retail_v2beta_model_proto_depIdxs = []int32{
	1, // 0: google.cloud.retail.v2beta.Model.training_state:type_name -> google.cloud.retail.v2beta.Model.TrainingState
	0, // 1: google.cloud.retail.v2beta.Model.serving_state:type_name -> google.cloud.retail.v2beta.Model.ServingState
	6, // 2: google.cloud.retail.v2beta.Model.create_time:type_name -> google.protobuf.Timestamp
	6, // 3: google.cloud.retail.v2beta.Model.update_time:type_name -> google.protobuf.Timestamp
	2, // 4: google.cloud.retail.v2beta.Model.periodic_tuning_state:type_name -> google.cloud.retail.v2beta.Model.PeriodicTuningState
	6, // 5: google.cloud.retail.v2beta.Model.last_tune_time:type_name -> google.protobuf.Timestamp
	3, // 6: google.cloud.retail.v2beta.Model.data_state:type_name -> google.cloud.retail.v2beta.Model.DataState
	7, // 7: google.cloud.retail.v2beta.Model.filtering_option:type_name -> google.cloud.retail.v2beta.RecommendationsFilteringOption
	5, // 8: google.cloud.retail.v2beta.Model.serving_config_lists:type_name -> google.cloud.retail.v2beta.Model.ServingConfigList
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2beta_model_proto_init() }
func file_google_cloud_retail_v2beta_model_proto_init() {
	if File_google_cloud_retail_v2beta_model_proto != nil {
		return
	}
	file_google_cloud_retail_v2beta_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_retail_v2beta_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model); i {
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
		file_google_cloud_retail_v2beta_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model_ServingConfigList); i {
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
			RawDescriptor: file_google_cloud_retail_v2beta_model_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_retail_v2beta_model_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2beta_model_proto_depIdxs,
		EnumInfos:         file_google_cloud_retail_v2beta_model_proto_enumTypes,
		MessageInfos:      file_google_cloud_retail_v2beta_model_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_v2beta_model_proto = out.File
	file_google_cloud_retail_v2beta_model_proto_rawDesc = nil
	file_google_cloud_retail_v2beta_model_proto_goTypes = nil
	file_google_cloud_retail_v2beta_model_proto_depIdxs = nil
}
