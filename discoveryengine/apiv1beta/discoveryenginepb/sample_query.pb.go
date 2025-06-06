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
// source: google/cloud/discoveryengine/v1beta/sample_query.proto

package discoveryenginepb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Sample Query captures metadata to be used for evaluation.
type SampleQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The content of the sample query.
	//
	// Types that are assignable to Content:
	//
	//	*SampleQuery_QueryEntry_
	Content isSampleQuery_Content `protobuf_oneof:"content"`
	// Identifier. The full resource name of the sample query, in the format of
	// `projects/{project}/locations/{location}/sampleQuerySets/{sample_query_set}/sampleQueries/{sample_query}`.
	//
	// This field must be a UTF-8 encoded string with a length limit of 1024
	// characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Timestamp the
	// [SampleQuery][google.cloud.discoveryengine.v1beta.SampleQuery] was created
	// at.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *SampleQuery) Reset() {
	*x = SampleQuery{}
	mi := &file_google_cloud_discoveryengine_v1beta_sample_query_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SampleQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleQuery) ProtoMessage() {}

func (x *SampleQuery) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_sample_query_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleQuery.ProtoReflect.Descriptor instead.
func (*SampleQuery) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_sample_query_proto_rawDescGZIP(), []int{0}
}

func (m *SampleQuery) GetContent() isSampleQuery_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *SampleQuery) GetQueryEntry() *SampleQuery_QueryEntry {
	if x, ok := x.GetContent().(*SampleQuery_QueryEntry_); ok {
		return x.QueryEntry
	}
	return nil
}

func (x *SampleQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SampleQuery) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

type isSampleQuery_Content interface {
	isSampleQuery_Content()
}

type SampleQuery_QueryEntry_ struct {
	// The query entry.
	QueryEntry *SampleQuery_QueryEntry `protobuf:"bytes,2,opt,name=query_entry,json=queryEntry,proto3,oneof"`
}

func (*SampleQuery_QueryEntry_) isSampleQuery_Content() {}

// Query Entry captures metadata to be used for search evaluation.
type SampleQuery_QueryEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The query.
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// List of targets for the query.
	Targets []*SampleQuery_QueryEntry_Target `protobuf:"bytes,3,rep,name=targets,proto3" json:"targets,omitempty"`
}

func (x *SampleQuery_QueryEntry) Reset() {
	*x = SampleQuery_QueryEntry{}
	mi := &file_google_cloud_discoveryengine_v1beta_sample_query_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SampleQuery_QueryEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleQuery_QueryEntry) ProtoMessage() {}

func (x *SampleQuery_QueryEntry) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_sample_query_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleQuery_QueryEntry.ProtoReflect.Descriptor instead.
func (*SampleQuery_QueryEntry) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_sample_query_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SampleQuery_QueryEntry) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SampleQuery_QueryEntry) GetTargets() []*SampleQuery_QueryEntry_Target {
	if x != nil {
		return x.Targets
	}
	return nil
}

// Defines the parameters of the query's expected outcome.
type SampleQuery_QueryEntry_Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Expected uri of the target.
	//
	// This field must be a UTF-8 encoded string with a length limit of 2048
	// characters.
	//
	// Example of valid uris: `https://example.com/abc`,
	// `gcs://example/example.pdf`.
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Expected page numbers of the target.
	//
	// Each page number must be non negative.
	PageNumbers []int32 `protobuf:"varint,2,rep,packed,name=page_numbers,json=pageNumbers,proto3" json:"page_numbers,omitempty"`
	// Relevance score of the target.
	Score *float64 `protobuf:"fixed64,3,opt,name=score,proto3,oneof" json:"score,omitempty"`
}

func (x *SampleQuery_QueryEntry_Target) Reset() {
	*x = SampleQuery_QueryEntry_Target{}
	mi := &file_google_cloud_discoveryengine_v1beta_sample_query_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SampleQuery_QueryEntry_Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleQuery_QueryEntry_Target) ProtoMessage() {}

func (x *SampleQuery_QueryEntry_Target) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_sample_query_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleQuery_QueryEntry_Target.ProtoReflect.Descriptor instead.
func (*SampleQuery_QueryEntry_Target) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_sample_query_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *SampleQuery_QueryEntry_Target) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *SampleQuery_QueryEntry_Target) GetPageNumbers() []int32 {
	if x != nil {
		return x.PageNumbers
	}
	return nil
}

func (x *SampleQuery_QueryEntry_Target) GetScore() float64 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

var File_google_cloud_discoveryengine_v1beta_sample_query_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1beta_sample_query_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x04, 0x0a, 0x0b, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x5e, 0x0a, 0x0b, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x48, 0x00, 0x52, 0x0a,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0xe9, 0x01, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x5c, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x1a, 0x62, 0x0a,
	0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0b, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x19, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x3a, 0x99, 0x01, 0xea, 0x41, 0x95, 0x01, 0x0a, 0x2a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x67, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x7d,
	0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x7d, 0x42, 0x09, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x97, 0x02, 0x0a, 0x27, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x42, 0x10, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0xa2, 0x02, 0x0f, 0x44, 0x49,
	0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0xaa, 0x02, 0x23,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0x42,
	0x65, 0x74, 0x61, 0xca, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_discoveryengine_v1beta_sample_query_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1beta_sample_query_proto_rawDescData = file_google_cloud_discoveryengine_v1beta_sample_query_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1beta_sample_query_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1beta_sample_query_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1beta_sample_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1beta_sample_query_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1beta_sample_query_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1beta_sample_query_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_discoveryengine_v1beta_sample_query_proto_goTypes = []any{
	(*SampleQuery)(nil),                   // 0: google.cloud.discoveryengine.v1beta.SampleQuery
	(*SampleQuery_QueryEntry)(nil),        // 1: google.cloud.discoveryengine.v1beta.SampleQuery.QueryEntry
	(*SampleQuery_QueryEntry_Target)(nil), // 2: google.cloud.discoveryengine.v1beta.SampleQuery.QueryEntry.Target
	(*timestamppb.Timestamp)(nil),         // 3: google.protobuf.Timestamp
}
var file_google_cloud_discoveryengine_v1beta_sample_query_proto_depIdxs = []int32{
	1, // 0: google.cloud.discoveryengine.v1beta.SampleQuery.query_entry:type_name -> google.cloud.discoveryengine.v1beta.SampleQuery.QueryEntry
	3, // 1: google.cloud.discoveryengine.v1beta.SampleQuery.create_time:type_name -> google.protobuf.Timestamp
	2, // 2: google.cloud.discoveryengine.v1beta.SampleQuery.QueryEntry.targets:type_name -> google.cloud.discoveryengine.v1beta.SampleQuery.QueryEntry.Target
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1beta_sample_query_proto_init() }
func file_google_cloud_discoveryengine_v1beta_sample_query_proto_init() {
	if File_google_cloud_discoveryengine_v1beta_sample_query_proto != nil {
		return
	}
	file_google_cloud_discoveryengine_v1beta_sample_query_proto_msgTypes[0].OneofWrappers = []any{
		(*SampleQuery_QueryEntry_)(nil),
	}
	file_google_cloud_discoveryengine_v1beta_sample_query_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_discoveryengine_v1beta_sample_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1beta_sample_query_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1beta_sample_query_proto_depIdxs,
		MessageInfos:      file_google_cloud_discoveryengine_v1beta_sample_query_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1beta_sample_query_proto = out.File
	file_google_cloud_discoveryengine_v1beta_sample_query_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1beta_sample_query_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1beta_sample_query_proto_depIdxs = nil
}
