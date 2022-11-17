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
// source: google/cloud/notebooks/v1/schedule.proto

package notebookspb

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

// State of the job.
type Schedule_State int32

const (
	// Unspecified state.
	Schedule_STATE_UNSPECIFIED Schedule_State = 0
	// The job is executing normally.
	Schedule_ENABLED Schedule_State = 1
	// The job is paused by the user. It will not execute. A user can
	// intentionally pause the job using
	// [PauseJobRequest][].
	Schedule_PAUSED Schedule_State = 2
	// The job is disabled by the system due to error. The user
	// cannot directly set a job to be disabled.
	Schedule_DISABLED Schedule_State = 3
	// The job state resulting from a failed [CloudScheduler.UpdateJob][]
	// operation. To recover a job from this state, retry
	// [CloudScheduler.UpdateJob][] until a successful response is received.
	Schedule_UPDATE_FAILED Schedule_State = 4
	// The schedule resource is being created.
	Schedule_INITIALIZING Schedule_State = 5
	// The schedule resource is being deleted.
	Schedule_DELETING Schedule_State = 6
)

// Enum value maps for Schedule_State.
var (
	Schedule_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ENABLED",
		2: "PAUSED",
		3: "DISABLED",
		4: "UPDATE_FAILED",
		5: "INITIALIZING",
		6: "DELETING",
	}
	Schedule_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ENABLED":           1,
		"PAUSED":            2,
		"DISABLED":          3,
		"UPDATE_FAILED":     4,
		"INITIALIZING":      5,
		"DELETING":          6,
	}
)

func (x Schedule_State) Enum() *Schedule_State {
	p := new(Schedule_State)
	*p = x
	return p
}

func (x Schedule_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Schedule_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_notebooks_v1_schedule_proto_enumTypes[0].Descriptor()
}

func (Schedule_State) Type() protoreflect.EnumType {
	return &file_google_cloud_notebooks_v1_schedule_proto_enumTypes[0]
}

func (x Schedule_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Schedule_State.Descriptor instead.
func (Schedule_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_notebooks_v1_schedule_proto_rawDescGZIP(), []int{0, 0}
}

// The definition of a schedule.
type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The name of this schedule. Format:
	// `projects/{project_id}/locations/{location}/schedules/{schedule_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Display name used for UI purposes.
	// Name can only contain alphanumeric characters, hyphens '-',
	// and underscores '_'.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// A brief description of this environment.
	Description string         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	State       Schedule_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.cloud.notebooks.v1.Schedule_State" json:"state,omitempty"`
	// Cron-tab formatted schedule by which the job will execute.
	// Format: minute, hour, day of month, month, day of week,
	// e.g. 0 0 * * WED = every Wednesday
	// More examples: https://crontab.guru/examples.html
	CronSchedule string `protobuf:"bytes,5,opt,name=cron_schedule,json=cronSchedule,proto3" json:"cron_schedule,omitempty"`
	// Timezone on which the cron_schedule.
	// The value of this field must be a time zone name from the tz database.
	// TZ Database: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	//
	// Note that some time zones include a provision for daylight savings time.
	// The rules for daylight saving time are determined by the chosen tz.
	// For UTC use the string "utc". If a time zone is not specified,
	// the default will be in UTC (also known as GMT).
	TimeZone string `protobuf:"bytes,6,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	// Output only. Time the schedule was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Time the schedule was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Notebook Execution Template corresponding to this schedule.
	ExecutionTemplate *ExecutionTemplate `protobuf:"bytes,9,opt,name=execution_template,json=executionTemplate,proto3" json:"execution_template,omitempty"`
	// Output only. The most recent execution names triggered from this schedule and their
	// corresponding states.
	RecentExecutions []*Execution `protobuf:"bytes,10,rep,name=recent_executions,json=recentExecutions,proto3" json:"recent_executions,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_notebooks_v1_schedule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_notebooks_v1_schedule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_google_cloud_notebooks_v1_schedule_proto_rawDescGZIP(), []int{0}
}

func (x *Schedule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Schedule) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Schedule) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Schedule) GetState() Schedule_State {
	if x != nil {
		return x.State
	}
	return Schedule_STATE_UNSPECIFIED
}

func (x *Schedule) GetCronSchedule() string {
	if x != nil {
		return x.CronSchedule
	}
	return ""
}

func (x *Schedule) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

func (x *Schedule) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Schedule) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Schedule) GetExecutionTemplate() *ExecutionTemplate {
	if x != nil {
		return x.ExecutionTemplate
	}
	return nil
}

func (x *Schedule) GetRecentExecutions() []*Execution {
	if x != nil {
		return x.RecentExecutions
	}
	return nil
}

var File_google_cloud_notebooks_v1_schedule_proto protoreflect.FileDescriptor

var file_google_cloud_notebooks_v1_schedule_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e,
	0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x06,
	0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x6f, 0x74, 0x65,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65,
	0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x5b, 0x0a, 0x12, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x11,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x56, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x6f, 0x74, 0x65,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x78, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41,
	0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x11, 0x0a, 0x0d, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a,
	0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x06, 0x3a, 0x63, 0xea, 0x41, 0x60, 0x0a, 0x21, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x3b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x7d, 0x42, 0x74, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_notebooks_v1_schedule_proto_rawDescOnce sync.Once
	file_google_cloud_notebooks_v1_schedule_proto_rawDescData = file_google_cloud_notebooks_v1_schedule_proto_rawDesc
)

func file_google_cloud_notebooks_v1_schedule_proto_rawDescGZIP() []byte {
	file_google_cloud_notebooks_v1_schedule_proto_rawDescOnce.Do(func() {
		file_google_cloud_notebooks_v1_schedule_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_notebooks_v1_schedule_proto_rawDescData)
	})
	return file_google_cloud_notebooks_v1_schedule_proto_rawDescData
}

var file_google_cloud_notebooks_v1_schedule_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_notebooks_v1_schedule_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_notebooks_v1_schedule_proto_goTypes = []interface{}{
	(Schedule_State)(0),           // 0: google.cloud.notebooks.v1.Schedule.State
	(*Schedule)(nil),              // 1: google.cloud.notebooks.v1.Schedule
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*ExecutionTemplate)(nil),     // 3: google.cloud.notebooks.v1.ExecutionTemplate
	(*Execution)(nil),             // 4: google.cloud.notebooks.v1.Execution
}
var file_google_cloud_notebooks_v1_schedule_proto_depIdxs = []int32{
	0, // 0: google.cloud.notebooks.v1.Schedule.state:type_name -> google.cloud.notebooks.v1.Schedule.State
	2, // 1: google.cloud.notebooks.v1.Schedule.create_time:type_name -> google.protobuf.Timestamp
	2, // 2: google.cloud.notebooks.v1.Schedule.update_time:type_name -> google.protobuf.Timestamp
	3, // 3: google.cloud.notebooks.v1.Schedule.execution_template:type_name -> google.cloud.notebooks.v1.ExecutionTemplate
	4, // 4: google.cloud.notebooks.v1.Schedule.recent_executions:type_name -> google.cloud.notebooks.v1.Execution
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_notebooks_v1_schedule_proto_init() }
func file_google_cloud_notebooks_v1_schedule_proto_init() {
	if File_google_cloud_notebooks_v1_schedule_proto != nil {
		return
	}
	file_google_cloud_notebooks_v1_execution_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_notebooks_v1_schedule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
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
			RawDescriptor: file_google_cloud_notebooks_v1_schedule_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_notebooks_v1_schedule_proto_goTypes,
		DependencyIndexes: file_google_cloud_notebooks_v1_schedule_proto_depIdxs,
		EnumInfos:         file_google_cloud_notebooks_v1_schedule_proto_enumTypes,
		MessageInfos:      file_google_cloud_notebooks_v1_schedule_proto_msgTypes,
	}.Build()
	File_google_cloud_notebooks_v1_schedule_proto = out.File
	file_google_cloud_notebooks_v1_schedule_proto_rawDesc = nil
	file_google_cloud_notebooks_v1_schedule_proto_goTypes = nil
	file_google_cloud_notebooks_v1_schedule_proto_depIdxs = nil
}
