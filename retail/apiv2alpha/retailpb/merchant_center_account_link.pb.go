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
// source: google/cloud/retail/v2alpha/merchant_center_account_link.proto

package retailpb

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

// The state of the link.
type MerchantCenterAccountLink_State int32

const (
	// Default value.
	MerchantCenterAccountLink_STATE_UNSPECIFIED MerchantCenterAccountLink_State = 0
	// Link is created and LRO is not complete.
	MerchantCenterAccountLink_PENDING MerchantCenterAccountLink_State = 1
	// Link is active.
	MerchantCenterAccountLink_ACTIVE MerchantCenterAccountLink_State = 2
	// Link creation failed.
	MerchantCenterAccountLink_FAILED MerchantCenterAccountLink_State = 3
)

// Enum value maps for MerchantCenterAccountLink_State.
var (
	MerchantCenterAccountLink_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "PENDING",
		2: "ACTIVE",
		3: "FAILED",
	}
	MerchantCenterAccountLink_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"PENDING":           1,
		"ACTIVE":            2,
		"FAILED":            3,
	}
)

func (x MerchantCenterAccountLink_State) Enum() *MerchantCenterAccountLink_State {
	p := new(MerchantCenterAccountLink_State)
	*p = x
	return p
}

func (x MerchantCenterAccountLink_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MerchantCenterAccountLink_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_enumTypes[0].Descriptor()
}

func (MerchantCenterAccountLink_State) Type() protoreflect.EnumType {
	return &file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_enumTypes[0]
}

func (x MerchantCenterAccountLink_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MerchantCenterAccountLink_State.Descriptor instead.
func (MerchantCenterAccountLink_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_rawDescGZIP(), []int{0, 0}
}

// Represents a link between a Merchant Center account and a branch.
// After a link is established, products from the linked Merchant Center account
// are streamed to the linked branch.
type MerchantCenterAccountLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Immutable. Full resource name of the Merchant Center Account
	// Link, such as
	// `projects/*/locations/global/catalogs/default_catalog/merchantCenterAccountLinks/merchant_center_account_link`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Immutable.
	// [MerchantCenterAccountLink][google.cloud.retail.v2alpha.MerchantCenterAccountLink]
	// identifier, which is the final component of
	// [name][google.cloud.retail.v2alpha.MerchantCenterAccountLink.name]. This
	// field is auto generated and follows the convention:
	// `BranchId_MerchantCenterAccountId`.
	// `projects/*/locations/global/catalogs/default_catalog/merchantCenterAccountLinks/id_1`.
	Id string `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
	// Required. The linked [Merchant center account
	// id](https://developers.google.com/shopping-content/guides/accountstatuses).
	// The account must be a standalone account or a sub-account of a MCA.
	MerchantCenterAccountId int64 `protobuf:"varint,2,opt,name=merchant_center_account_id,json=merchantCenterAccountId,proto3" json:"merchant_center_account_id,omitempty"`
	// Required. The branch ID (e.g. 0/1/2) within the catalog that products from
	// merchant_center_account_id are streamed to. When updating this field, an
	// empty value will use the currently configured default branch. However,
	// changing the default branch later on won't change the linked branch here.
	//
	// A single branch ID can only have one linked Merchant Center account ID.
	BranchId string `protobuf:"bytes,3,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	// The FeedLabel used to perform filtering.
	// Note: this replaces
	// [region_id](https://developers.google.com/shopping-content/reference/rest/v2.1/products#Product.FIELDS.feed_label).
	//
	// Example value: `US`.
	// Example value: `FeedLabel1`.
	FeedLabel string `protobuf:"bytes,4,opt,name=feed_label,json=feedLabel,proto3" json:"feed_label,omitempty"`
	// Language of the title/description and other string attributes. Use language
	// tags defined by [BCP 47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt).
	// ISO 639-1.
	//
	// This specifies the language of offers in Merchant Center that will be
	// accepted. If empty, no language filtering will be performed.
	//
	// Example value: `en`.
	LanguageCode string `protobuf:"bytes,5,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Criteria for the Merchant Center feeds to be ingested via the link.
	// All offers will be ingested if the list is empty.
	// Otherwise the offers will be ingested from selected feeds.
	FeedFilters []*MerchantCenterAccountLink_MerchantCenterFeedFilter `protobuf:"bytes,6,rep,name=feed_filters,json=feedFilters,proto3" json:"feed_filters,omitempty"`
	// Output only. Represents the state of the link.
	State MerchantCenterAccountLink_State `protobuf:"varint,7,opt,name=state,proto3,enum=google.cloud.retail.v2alpha.MerchantCenterAccountLink_State" json:"state,omitempty"`
	// Output only. Google Cloud project ID.
	ProjectId string `protobuf:"bytes,9,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Optional. An optional arbitrary string that could be used as a tag for
	// tracking link source.
	Source string `protobuf:"bytes,10,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *MerchantCenterAccountLink) Reset() {
	*x = MerchantCenterAccountLink{}
	mi := &file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MerchantCenterAccountLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantCenterAccountLink) ProtoMessage() {}

func (x *MerchantCenterAccountLink) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantCenterAccountLink.ProtoReflect.Descriptor instead.
func (*MerchantCenterAccountLink) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_rawDescGZIP(), []int{0}
}

func (x *MerchantCenterAccountLink) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MerchantCenterAccountLink) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MerchantCenterAccountLink) GetMerchantCenterAccountId() int64 {
	if x != nil {
		return x.MerchantCenterAccountId
	}
	return 0
}

func (x *MerchantCenterAccountLink) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *MerchantCenterAccountLink) GetFeedLabel() string {
	if x != nil {
		return x.FeedLabel
	}
	return ""
}

func (x *MerchantCenterAccountLink) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *MerchantCenterAccountLink) GetFeedFilters() []*MerchantCenterAccountLink_MerchantCenterFeedFilter {
	if x != nil {
		return x.FeedFilters
	}
	return nil
}

func (x *MerchantCenterAccountLink) GetState() MerchantCenterAccountLink_State {
	if x != nil {
		return x.State
	}
	return MerchantCenterAccountLink_STATE_UNSPECIFIED
}

func (x *MerchantCenterAccountLink) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *MerchantCenterAccountLink) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

// Common metadata related to the progress of the operations.
type CreateMerchantCenterAccountLinkMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Operation create time.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Operation last update time. If the operation is done, this is also the
	// finish time.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *CreateMerchantCenterAccountLinkMetadata) Reset() {
	*x = CreateMerchantCenterAccountLinkMetadata{}
	mi := &file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMerchantCenterAccountLinkMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMerchantCenterAccountLinkMetadata) ProtoMessage() {}

func (x *CreateMerchantCenterAccountLinkMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMerchantCenterAccountLinkMetadata.ProtoReflect.Descriptor instead.
func (*CreateMerchantCenterAccountLinkMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMerchantCenterAccountLinkMetadata) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *CreateMerchantCenterAccountLinkMetadata) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// Merchant Center Feed filter criterion.
type MerchantCenterAccountLink_MerchantCenterFeedFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Merchant Center primary feed ID.
	// Deprecated: use data_source_id instead.
	//
	// Deprecated: Marked as deprecated in google/cloud/retail/v2alpha/merchant_center_account_link.proto.
	PrimaryFeedId int64 `protobuf:"varint,1,opt,name=primary_feed_id,json=primaryFeedId,proto3" json:"primary_feed_id,omitempty"`
	// AFM data source ID.
	DataSourceId int64 `protobuf:"varint,3,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	// Merchant Center primary feed name. The name is used for the display
	// purposes only.
	PrimaryFeedName string `protobuf:"bytes,2,opt,name=primary_feed_name,json=primaryFeedName,proto3" json:"primary_feed_name,omitempty"`
}

func (x *MerchantCenterAccountLink_MerchantCenterFeedFilter) Reset() {
	*x = MerchantCenterAccountLink_MerchantCenterFeedFilter{}
	mi := &file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MerchantCenterAccountLink_MerchantCenterFeedFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantCenterAccountLink_MerchantCenterFeedFilter) ProtoMessage() {}

func (x *MerchantCenterAccountLink_MerchantCenterFeedFilter) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantCenterAccountLink_MerchantCenterFeedFilter.ProtoReflect.Descriptor instead.
func (*MerchantCenterAccountLink_MerchantCenterFeedFilter) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_rawDescGZIP(), []int{0, 0}
}

// Deprecated: Marked as deprecated in google/cloud/retail/v2alpha/merchant_center_account_link.proto.
func (x *MerchantCenterAccountLink_MerchantCenterFeedFilter) GetPrimaryFeedId() int64 {
	if x != nil {
		return x.PrimaryFeedId
	}
	return 0
}

func (x *MerchantCenterAccountLink_MerchantCenterFeedFilter) GetDataSourceId() int64 {
	if x != nil {
		return x.DataSourceId
	}
	return 0
}

func (x *MerchantCenterAccountLink_MerchantCenterFeedFilter) GetPrimaryFeedName() string {
	if x != nil {
		return x.PrimaryFeedName
	}
	return ""
}

var File_google_cloud_retail_v2alpha_merchant_center_account_link_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x07, 0x0a, 0x19, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x05, 0xe0, 0x41, 0x03, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xe0, 0x41, 0x05, 0xe0, 0x41, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x1a,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x17, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x65, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x72, 0x0a, 0x0c, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x46, 0x65, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x66, 0x65, 0x65,
	0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x57, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x1a, 0x98, 0x01, 0x0a, 0x18, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x2a, 0x0a, 0x0f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0d, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x46, 0x65, 0x65, 0x64, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x66, 0x65, 0x65,
	0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x46, 0x65, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x43, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x03, 0x3a, 0xab, 0x01, 0xea, 0x41, 0xa7, 0x01, 0x0a, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x74, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x7d, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x7d,
	0x22, 0xa3, 0x01, 0x0a, 0x27, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0xe2, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x1e, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x70, 0x62, 0x3b, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x70, 0x62, 0xa2, 0x02, 0x06, 0x52, 0x45, 0x54, 0x41, 0x49, 0x4c, 0xaa, 0x02,
	0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x56, 0x32, 0x41, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x1b, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x52, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_rawDescOnce sync.Once
	file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_rawDescData = file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_rawDesc
)

func file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_rawDescData)
	})
	return file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_rawDescData
}

var file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_goTypes = []any{
	(MerchantCenterAccountLink_State)(0),                       // 0: google.cloud.retail.v2alpha.MerchantCenterAccountLink.State
	(*MerchantCenterAccountLink)(nil),                          // 1: google.cloud.retail.v2alpha.MerchantCenterAccountLink
	(*CreateMerchantCenterAccountLinkMetadata)(nil),            // 2: google.cloud.retail.v2alpha.CreateMerchantCenterAccountLinkMetadata
	(*MerchantCenterAccountLink_MerchantCenterFeedFilter)(nil), // 3: google.cloud.retail.v2alpha.MerchantCenterAccountLink.MerchantCenterFeedFilter
	(*timestamppb.Timestamp)(nil),                              // 4: google.protobuf.Timestamp
}
var file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_depIdxs = []int32{
	3, // 0: google.cloud.retail.v2alpha.MerchantCenterAccountLink.feed_filters:type_name -> google.cloud.retail.v2alpha.MerchantCenterAccountLink.MerchantCenterFeedFilter
	0, // 1: google.cloud.retail.v2alpha.MerchantCenterAccountLink.state:type_name -> google.cloud.retail.v2alpha.MerchantCenterAccountLink.State
	4, // 2: google.cloud.retail.v2alpha.CreateMerchantCenterAccountLinkMetadata.create_time:type_name -> google.protobuf.Timestamp
	4, // 3: google.cloud.retail.v2alpha.CreateMerchantCenterAccountLinkMetadata.update_time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_init() }
func file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_init() {
	if File_google_cloud_retail_v2alpha_merchant_center_account_link_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_depIdxs,
		EnumInfos:         file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_enumTypes,
		MessageInfos:      file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_v2alpha_merchant_center_account_link_proto = out.File
	file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_rawDesc = nil
	file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_goTypes = nil
	file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_depIdxs = nil
}
