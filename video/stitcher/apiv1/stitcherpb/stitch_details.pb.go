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
// source: google/cloud/video/stitcher/v1/stitch_details.proto

package stitcherpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Information related to the interstitial of a VOD session. This resource is
// only available for VOD sessions that do not implement Google Ad Manager ad
// insertion.
type VodStitchDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the stitch detail in the specified VOD session, in the form of
	// `projects/{project}/locations/{location}/vodSessions/{vod_session_id}/vodStitchDetails/{id}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of ad processing details for the fetched ad playlist.
	AdStitchDetails []*AdStitchDetail `protobuf:"bytes,3,rep,name=ad_stitch_details,json=adStitchDetails,proto3" json:"ad_stitch_details,omitempty"`
}

func (x *VodStitchDetail) Reset() {
	*x = VodStitchDetail{}
	mi := &file_google_cloud_video_stitcher_v1_stitch_details_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VodStitchDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodStitchDetail) ProtoMessage() {}

func (x *VodStitchDetail) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_stitch_details_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodStitchDetail.ProtoReflect.Descriptor instead.
func (*VodStitchDetail) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_stitch_details_proto_rawDescGZIP(), []int{0}
}

func (x *VodStitchDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VodStitchDetail) GetAdStitchDetails() []*AdStitchDetail {
	if x != nil {
		return x.AdStitchDetails
	}
	return nil
}

// Metadata for a stitched ad.
type AdStitchDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ad break ID of the processed ad.
	AdBreakId string `protobuf:"bytes,1,opt,name=ad_break_id,json=adBreakId,proto3" json:"ad_break_id,omitempty"`
	// Required. The ad ID of the processed ad.
	AdId string `protobuf:"bytes,2,opt,name=ad_id,json=adId,proto3" json:"ad_id,omitempty"`
	// Required. The time offset of the processed ad.
	AdTimeOffset *durationpb.Duration `protobuf:"bytes,3,opt,name=ad_time_offset,json=adTimeOffset,proto3" json:"ad_time_offset,omitempty"`
	// Optional. Indicates the reason why the ad has been skipped.
	SkipReason string `protobuf:"bytes,4,opt,name=skip_reason,json=skipReason,proto3" json:"skip_reason,omitempty"`
	// Optional. The metadata of the chosen media file for the ad.
	Media map[string]*structpb.Value `protobuf:"bytes,5,rep,name=media,proto3" json:"media,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AdStitchDetail) Reset() {
	*x = AdStitchDetail{}
	mi := &file_google_cloud_video_stitcher_v1_stitch_details_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdStitchDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdStitchDetail) ProtoMessage() {}

func (x *AdStitchDetail) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_stitch_details_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdStitchDetail.ProtoReflect.Descriptor instead.
func (*AdStitchDetail) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_stitch_details_proto_rawDescGZIP(), []int{1}
}

func (x *AdStitchDetail) GetAdBreakId() string {
	if x != nil {
		return x.AdBreakId
	}
	return ""
}

func (x *AdStitchDetail) GetAdId() string {
	if x != nil {
		return x.AdId
	}
	return ""
}

func (x *AdStitchDetail) GetAdTimeOffset() *durationpb.Duration {
	if x != nil {
		return x.AdTimeOffset
	}
	return nil
}

func (x *AdStitchDetail) GetSkipReason() string {
	if x != nil {
		return x.SkipReason
	}
	return ""
}

func (x *AdStitchDetail) GetMedia() map[string]*structpb.Value {
	if x != nil {
		return x.Media
	}
	return nil
}

var File_google_cloud_video_stitcher_v1_stitch_details_proto protoreflect.FileDescriptor

var file_google_cloud_video_stitcher_v1_stitch_details_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9e, 0x02, 0x0a, 0x0f, 0x56, 0x6f, 0x64, 0x53, 0x74, 0x69, 0x74, 0x63, 0x68, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x5a, 0x0a, 0x11, 0x61, 0x64, 0x5f, 0x73, 0x74,
	0x69, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x53, 0x74, 0x69, 0x74, 0x63, 0x68, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x0f, 0x61, 0x64, 0x53, 0x74, 0x69, 0x74, 0x63, 0x68, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x3a, 0x9a, 0x01, 0xea, 0x41, 0x96, 0x01, 0x0a, 0x2c, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x6f, 0x64, 0x53, 0x74, 0x69, 0x74,
	0x63, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x66, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x7d, 0x2f, 0x76, 0x6f, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x76,
	0x6f, 0x64, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x76, 0x6f, 0x64, 0x53,
	0x74, 0x69, 0x74, 0x63, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x7b, 0x76, 0x6f,
	0x64, 0x5f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x7d,
	0x22, 0xe3, 0x02, 0x0a, 0x0e, 0x41, 0x64, 0x53, 0x74, 0x69, 0x74, 0x63, 0x68, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0b, 0x61, 0x64, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x61,
	0x64, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x05, 0x61, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x61, 0x64,
	0x49, 0x64, 0x12, 0x44, 0x0a, 0x0e, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0c, 0x61, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x73, 0x6b, 0x69, 0x70,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x0a, 0x73, 0x6b, 0x69, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x54,
	0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x53, 0x74, 0x69, 0x74, 0x63, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x05, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x1a, 0x50, 0x0a, 0x0a, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x7a, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x53, 0x74,
	0x69, 0x74, 0x63, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x73, 0x74,
	0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x69,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_video_stitcher_v1_stitch_details_proto_rawDescOnce sync.Once
	file_google_cloud_video_stitcher_v1_stitch_details_proto_rawDescData = file_google_cloud_video_stitcher_v1_stitch_details_proto_rawDesc
)

func file_google_cloud_video_stitcher_v1_stitch_details_proto_rawDescGZIP() []byte {
	file_google_cloud_video_stitcher_v1_stitch_details_proto_rawDescOnce.Do(func() {
		file_google_cloud_video_stitcher_v1_stitch_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_video_stitcher_v1_stitch_details_proto_rawDescData)
	})
	return file_google_cloud_video_stitcher_v1_stitch_details_proto_rawDescData
}

var file_google_cloud_video_stitcher_v1_stitch_details_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_video_stitcher_v1_stitch_details_proto_goTypes = []any{
	(*VodStitchDetail)(nil),     // 0: google.cloud.video.stitcher.v1.VodStitchDetail
	(*AdStitchDetail)(nil),      // 1: google.cloud.video.stitcher.v1.AdStitchDetail
	nil,                         // 2: google.cloud.video.stitcher.v1.AdStitchDetail.MediaEntry
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
	(*structpb.Value)(nil),      // 4: google.protobuf.Value
}
var file_google_cloud_video_stitcher_v1_stitch_details_proto_depIdxs = []int32{
	1, // 0: google.cloud.video.stitcher.v1.VodStitchDetail.ad_stitch_details:type_name -> google.cloud.video.stitcher.v1.AdStitchDetail
	3, // 1: google.cloud.video.stitcher.v1.AdStitchDetail.ad_time_offset:type_name -> google.protobuf.Duration
	2, // 2: google.cloud.video.stitcher.v1.AdStitchDetail.media:type_name -> google.cloud.video.stitcher.v1.AdStitchDetail.MediaEntry
	4, // 3: google.cloud.video.stitcher.v1.AdStitchDetail.MediaEntry.value:type_name -> google.protobuf.Value
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_video_stitcher_v1_stitch_details_proto_init() }
func file_google_cloud_video_stitcher_v1_stitch_details_proto_init() {
	if File_google_cloud_video_stitcher_v1_stitch_details_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_video_stitcher_v1_stitch_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_video_stitcher_v1_stitch_details_proto_goTypes,
		DependencyIndexes: file_google_cloud_video_stitcher_v1_stitch_details_proto_depIdxs,
		MessageInfos:      file_google_cloud_video_stitcher_v1_stitch_details_proto_msgTypes,
	}.Build()
	File_google_cloud_video_stitcher_v1_stitch_details_proto = out.File
	file_google_cloud_video_stitcher_v1_stitch_details_proto_rawDesc = nil
	file_google_cloud_video_stitcher_v1_stitch_details_proto_goTypes = nil
	file_google_cloud_video_stitcher_v1_stitch_details_proto_depIdxs = nil
}
