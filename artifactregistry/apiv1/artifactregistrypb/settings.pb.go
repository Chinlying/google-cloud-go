// Copyright 2025 Google LLC
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
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.7
// source: google/devtools/artifactregistry/v1/settings.proto

package artifactregistrypb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The possible redirection states for legacy repositories.
type ProjectSettings_RedirectionState int32

const (
	// No redirection status has been set.
	ProjectSettings_REDIRECTION_STATE_UNSPECIFIED ProjectSettings_RedirectionState = 0
	// Redirection is disabled.
	ProjectSettings_REDIRECTION_FROM_GCR_IO_DISABLED ProjectSettings_RedirectionState = 1
	// Redirection is enabled.
	ProjectSettings_REDIRECTION_FROM_GCR_IO_ENABLED ProjectSettings_RedirectionState = 2
	// Redirection is enabled, and has been finalized so cannot be reverted.
	//
	// Deprecated: Marked as deprecated in google/devtools/artifactregistry/v1/settings.proto.
	ProjectSettings_REDIRECTION_FROM_GCR_IO_FINALIZED ProjectSettings_RedirectionState = 3
	// Redirection is enabled and missing images are copied from GCR
	ProjectSettings_REDIRECTION_FROM_GCR_IO_ENABLED_AND_COPYING ProjectSettings_RedirectionState = 5
	// Redirection is partially enabled and missing images are copied from GCR
	ProjectSettings_REDIRECTION_FROM_GCR_IO_PARTIAL_AND_COPYING ProjectSettings_RedirectionState = 6
)

// Enum value maps for ProjectSettings_RedirectionState.
var (
	ProjectSettings_RedirectionState_name = map[int32]string{
		0: "REDIRECTION_STATE_UNSPECIFIED",
		1: "REDIRECTION_FROM_GCR_IO_DISABLED",
		2: "REDIRECTION_FROM_GCR_IO_ENABLED",
		3: "REDIRECTION_FROM_GCR_IO_FINALIZED",
		5: "REDIRECTION_FROM_GCR_IO_ENABLED_AND_COPYING",
		6: "REDIRECTION_FROM_GCR_IO_PARTIAL_AND_COPYING",
	}
	ProjectSettings_RedirectionState_value = map[string]int32{
		"REDIRECTION_STATE_UNSPECIFIED":               0,
		"REDIRECTION_FROM_GCR_IO_DISABLED":            1,
		"REDIRECTION_FROM_GCR_IO_ENABLED":             2,
		"REDIRECTION_FROM_GCR_IO_FINALIZED":           3,
		"REDIRECTION_FROM_GCR_IO_ENABLED_AND_COPYING": 5,
		"REDIRECTION_FROM_GCR_IO_PARTIAL_AND_COPYING": 6,
	}
)

func (x ProjectSettings_RedirectionState) Enum() *ProjectSettings_RedirectionState {
	p := new(ProjectSettings_RedirectionState)
	*p = x
	return p
}

func (x ProjectSettings_RedirectionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectSettings_RedirectionState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_artifactregistry_v1_settings_proto_enumTypes[0].Descriptor()
}

func (ProjectSettings_RedirectionState) Type() protoreflect.EnumType {
	return &file_google_devtools_artifactregistry_v1_settings_proto_enumTypes[0]
}

func (x ProjectSettings_RedirectionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectSettings_RedirectionState.Descriptor instead.
func (ProjectSettings_RedirectionState) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1_settings_proto_rawDescGZIP(), []int{0, 0}
}

// The Artifact Registry settings that apply to a Project.
type ProjectSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the project's settings.
	//
	// Always of the form:
	// projects/{project-id}/projectSettings
	//
	// In update request: never set
	// In response: always set
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The redirection state of the legacy repositories in this project.
	LegacyRedirectionState ProjectSettings_RedirectionState `protobuf:"varint,2,opt,name=legacy_redirection_state,json=legacyRedirectionState,proto3,enum=google.devtools.artifactregistry.v1.ProjectSettings_RedirectionState" json:"legacy_redirection_state,omitempty"`
	// The percentage of pull traffic to redirect from GCR to AR when using
	// partial redirection.
	PullPercent int32 `protobuf:"varint,3,opt,name=pull_percent,json=pullPercent,proto3" json:"pull_percent,omitempty"`
}

func (x *ProjectSettings) Reset() {
	*x = ProjectSettings{}
	mi := &file_google_devtools_artifactregistry_v1_settings_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectSettings) ProtoMessage() {}

func (x *ProjectSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1_settings_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectSettings.ProtoReflect.Descriptor instead.
func (*ProjectSettings) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1_settings_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectSettings) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectSettings) GetLegacyRedirectionState() ProjectSettings_RedirectionState {
	if x != nil {
		return x.LegacyRedirectionState
	}
	return ProjectSettings_REDIRECTION_STATE_UNSPECIFIED
}

func (x *ProjectSettings) GetPullPercent() int32 {
	if x != nil {
		return x.PullPercent
	}
	return 0
}

// Gets the redirection status for a project.
type GetProjectSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the projectSettings resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetProjectSettingsRequest) Reset() {
	*x = GetProjectSettingsRequest{}
	mi := &file_google_devtools_artifactregistry_v1_settings_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProjectSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectSettingsRequest) ProtoMessage() {}

func (x *GetProjectSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1_settings_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectSettingsRequest.ProtoReflect.Descriptor instead.
func (*GetProjectSettingsRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1_settings_proto_rawDescGZIP(), []int{1}
}

func (x *GetProjectSettingsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Sets the settings of the project.
type UpdateProjectSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The project settings.
	ProjectSettings *ProjectSettings `protobuf:"bytes,2,opt,name=project_settings,json=projectSettings,proto3" json:"project_settings,omitempty"`
	// Field mask to support partial updates.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateProjectSettingsRequest) Reset() {
	*x = UpdateProjectSettingsRequest{}
	mi := &file_google_devtools_artifactregistry_v1_settings_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProjectSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProjectSettingsRequest) ProtoMessage() {}

func (x *UpdateProjectSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1_settings_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProjectSettingsRequest.ProtoReflect.Descriptor instead.
func (*UpdateProjectSettingsRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1_settings_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateProjectSettingsRequest) GetProjectSettings() *ProjectSettings {
	if x != nil {
		return x.ProjectSettings
	}
	return nil
}

func (x *UpdateProjectSettingsRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_google_devtools_artifactregistry_v1_settings_proto protoreflect.FileDescriptor

var file_google_devtools_artifactregistry_v1_settings_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x04, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x7f, 0x0a, 0x18, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x16, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79,
	0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x75, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x22, 0x8d, 0x02, 0x0a, 0x10, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x45, 0x44, 0x49,
	0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x52,
	0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f,
	0x47, 0x43, 0x52, 0x5f, 0x49, 0x4f, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x23, 0x0a, 0x1f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x47, 0x43, 0x52, 0x5f, 0x49, 0x4f, 0x5f, 0x45, 0x4e, 0x41,
	0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x29, 0x0a, 0x21, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x47, 0x43, 0x52, 0x5f, 0x49,
	0x4f, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x03, 0x1a, 0x02, 0x08,
	0x01, 0x12, 0x2f, 0x0a, 0x2b, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x47, 0x43, 0x52, 0x5f, 0x49, 0x4f, 0x5f, 0x45, 0x4e, 0x41,
	0x42, 0x4c, 0x45, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x43, 0x4f, 0x50, 0x59, 0x49, 0x4e, 0x47,
	0x10, 0x05, 0x12, 0x2f, 0x0a, 0x2b, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x47, 0x43, 0x52, 0x5f, 0x49, 0x4f, 0x5f, 0x50, 0x41,
	0x52, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x43, 0x4f, 0x50, 0x59, 0x49, 0x4e,
	0x47, 0x10, 0x06, 0x3a, 0x58, 0xea, 0x41, 0x55, 0x0a, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x68, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x37, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x31,
	0x0a, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xbc, 0x01, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5f, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0xf8, 0x01, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x42, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x50, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x70, 0x62, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x70, 0x62, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x23, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_artifactregistry_v1_settings_proto_rawDescOnce sync.Once
	file_google_devtools_artifactregistry_v1_settings_proto_rawDescData = file_google_devtools_artifactregistry_v1_settings_proto_rawDesc
)

func file_google_devtools_artifactregistry_v1_settings_proto_rawDescGZIP() []byte {
	file_google_devtools_artifactregistry_v1_settings_proto_rawDescOnce.Do(func() {
		file_google_devtools_artifactregistry_v1_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_artifactregistry_v1_settings_proto_rawDescData)
	})
	return file_google_devtools_artifactregistry_v1_settings_proto_rawDescData
}

var file_google_devtools_artifactregistry_v1_settings_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_devtools_artifactregistry_v1_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_devtools_artifactregistry_v1_settings_proto_goTypes = []any{
	(ProjectSettings_RedirectionState)(0), // 0: google.devtools.artifactregistry.v1.ProjectSettings.RedirectionState
	(*ProjectSettings)(nil),               // 1: google.devtools.artifactregistry.v1.ProjectSettings
	(*GetProjectSettingsRequest)(nil),     // 2: google.devtools.artifactregistry.v1.GetProjectSettingsRequest
	(*UpdateProjectSettingsRequest)(nil),  // 3: google.devtools.artifactregistry.v1.UpdateProjectSettingsRequest
	(*fieldmaskpb.FieldMask)(nil),         // 4: google.protobuf.FieldMask
}
var file_google_devtools_artifactregistry_v1_settings_proto_depIdxs = []int32{
	0, // 0: google.devtools.artifactregistry.v1.ProjectSettings.legacy_redirection_state:type_name -> google.devtools.artifactregistry.v1.ProjectSettings.RedirectionState
	1, // 1: google.devtools.artifactregistry.v1.UpdateProjectSettingsRequest.project_settings:type_name -> google.devtools.artifactregistry.v1.ProjectSettings
	4, // 2: google.devtools.artifactregistry.v1.UpdateProjectSettingsRequest.update_mask:type_name -> google.protobuf.FieldMask
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_devtools_artifactregistry_v1_settings_proto_init() }
func file_google_devtools_artifactregistry_v1_settings_proto_init() {
	if File_google_devtools_artifactregistry_v1_settings_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_devtools_artifactregistry_v1_settings_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_artifactregistry_v1_settings_proto_goTypes,
		DependencyIndexes: file_google_devtools_artifactregistry_v1_settings_proto_depIdxs,
		EnumInfos:         file_google_devtools_artifactregistry_v1_settings_proto_enumTypes,
		MessageInfos:      file_google_devtools_artifactregistry_v1_settings_proto_msgTypes,
	}.Build()
	File_google_devtools_artifactregistry_v1_settings_proto = out.File
	file_google_devtools_artifactregistry_v1_settings_proto_rawDesc = nil
	file_google_devtools_artifactregistry_v1_settings_proto_goTypes = nil
	file_google_devtools_artifactregistry_v1_settings_proto_depIdxs = nil
}
