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
// source: google/cloud/alloydb/v1beta/csql_resources.proto

package alloydbpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The source CloudSQL backup resource.
type CloudSQLBackupRunSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The project ID of the source CloudSQL instance. This should be the same as
	// the AlloyDB cluster's project.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Required. The CloudSQL instance ID.
	InstanceId string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// Required. The CloudSQL backup run ID.
	BackupRunId int64 `protobuf:"varint,3,opt,name=backup_run_id,json=backupRunId,proto3" json:"backup_run_id,omitempty"`
}

func (x *CloudSQLBackupRunSource) Reset() {
	*x = CloudSQLBackupRunSource{}
	mi := &file_google_cloud_alloydb_v1beta_csql_resources_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CloudSQLBackupRunSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudSQLBackupRunSource) ProtoMessage() {}

func (x *CloudSQLBackupRunSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_alloydb_v1beta_csql_resources_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudSQLBackupRunSource.ProtoReflect.Descriptor instead.
func (*CloudSQLBackupRunSource) Descriptor() ([]byte, []int) {
	return file_google_cloud_alloydb_v1beta_csql_resources_proto_rawDescGZIP(), []int{0}
}

func (x *CloudSQLBackupRunSource) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *CloudSQLBackupRunSource) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *CloudSQLBackupRunSource) GetBackupRunId() int64 {
	if x != nil {
		return x.BackupRunId
	}
	return 0
}

var File_google_cloud_alloydb_v1beta_csql_resources_proto protoreflect.FileDescriptor

var file_google_cloud_alloydb_v1beta_csql_resources_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x6c, 0x6c, 0x6f, 0x79, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x73,
	0x71, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x79, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x82, 0x01, 0x0a, 0x17, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x51, 0x4c, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x52, 0x75, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0d,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x52, 0x75, 0x6e, 0x49, 0x64, 0x42, 0xcf, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x79,
	0x64, 0x62, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x12, 0x43, 0x73, 0x71, 0x6c, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x39, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x79, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x79, 0x64, 0x62, 0x70, 0x62,
	0x3b, 0x61, 0x6c, 0x6c, 0x6f, 0x79, 0x64, 0x62, 0x70, 0x62, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x79, 0x44,
	0x62, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x6c, 0x6c, 0x6f, 0x79, 0x44, 0x62, 0x5c,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x6c, 0x6c, 0x6f, 0x79, 0x44, 0x42, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_alloydb_v1beta_csql_resources_proto_rawDescOnce sync.Once
	file_google_cloud_alloydb_v1beta_csql_resources_proto_rawDescData = file_google_cloud_alloydb_v1beta_csql_resources_proto_rawDesc
)

func file_google_cloud_alloydb_v1beta_csql_resources_proto_rawDescGZIP() []byte {
	file_google_cloud_alloydb_v1beta_csql_resources_proto_rawDescOnce.Do(func() {
		file_google_cloud_alloydb_v1beta_csql_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_alloydb_v1beta_csql_resources_proto_rawDescData)
	})
	return file_google_cloud_alloydb_v1beta_csql_resources_proto_rawDescData
}

var file_google_cloud_alloydb_v1beta_csql_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_alloydb_v1beta_csql_resources_proto_goTypes = []any{
	(*CloudSQLBackupRunSource)(nil), // 0: google.cloud.alloydb.v1beta.CloudSQLBackupRunSource
}
var file_google_cloud_alloydb_v1beta_csql_resources_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_alloydb_v1beta_csql_resources_proto_init() }
func file_google_cloud_alloydb_v1beta_csql_resources_proto_init() {
	if File_google_cloud_alloydb_v1beta_csql_resources_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_alloydb_v1beta_csql_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_alloydb_v1beta_csql_resources_proto_goTypes,
		DependencyIndexes: file_google_cloud_alloydb_v1beta_csql_resources_proto_depIdxs,
		MessageInfos:      file_google_cloud_alloydb_v1beta_csql_resources_proto_msgTypes,
	}.Build()
	File_google_cloud_alloydb_v1beta_csql_resources_proto = out.File
	file_google_cloud_alloydb_v1beta_csql_resources_proto_rawDesc = nil
	file_google_cloud_alloydb_v1beta_csql_resources_proto_goTypes = nil
	file_google_cloud_alloydb_v1beta_csql_resources_proto_depIdxs = nil
}
