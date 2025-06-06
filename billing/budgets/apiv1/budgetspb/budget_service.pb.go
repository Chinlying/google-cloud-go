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
// source: google/cloud/billing/budgets/v1/budget_service.proto

package budgetspb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// Request for CreateBudget
type CreateBudgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the billing account to create the budget in. Values
	// are of the form `billingAccounts/{billingAccountId}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. Budget to create.
	Budget *Budget `protobuf:"bytes,2,opt,name=budget,proto3" json:"budget,omitempty"`
}

func (x *CreateBudgetRequest) Reset() {
	*x = CreateBudgetRequest{}
	mi := &file_google_cloud_billing_budgets_v1_budget_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBudgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBudgetRequest) ProtoMessage() {}

func (x *CreateBudgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_billing_budgets_v1_budget_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBudgetRequest.ProtoReflect.Descriptor instead.
func (*CreateBudgetRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_billing_budgets_v1_budget_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBudgetRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateBudgetRequest) GetBudget() *Budget {
	if x != nil {
		return x.Budget
	}
	return nil
}

// Request for UpdateBudget
type UpdateBudgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The updated budget object.
	// The budget to update is specified by the budget name in the budget.
	Budget *Budget `protobuf:"bytes,1,opt,name=budget,proto3" json:"budget,omitempty"`
	// Optional. Indicates which fields in the provided budget to update.
	// Read-only fields (such as `name`) cannot be changed. If this is not
	// provided, then only fields with non-default values from the request are
	// updated. See
	// https://developers.google.com/protocol-buffers/docs/proto3#default for more
	// details about default values.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateBudgetRequest) Reset() {
	*x = UpdateBudgetRequest{}
	mi := &file_google_cloud_billing_budgets_v1_budget_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBudgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBudgetRequest) ProtoMessage() {}

func (x *UpdateBudgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_billing_budgets_v1_budget_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBudgetRequest.ProtoReflect.Descriptor instead.
func (*UpdateBudgetRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_billing_budgets_v1_budget_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateBudgetRequest) GetBudget() *Budget {
	if x != nil {
		return x.Budget
	}
	return nil
}

func (x *UpdateBudgetRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// Request for GetBudget
type GetBudgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of budget to get. Values are of the form
	// `billingAccounts/{billingAccountId}/budgets/{budgetId}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetBudgetRequest) Reset() {
	*x = GetBudgetRequest{}
	mi := &file_google_cloud_billing_budgets_v1_budget_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBudgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBudgetRequest) ProtoMessage() {}

func (x *GetBudgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_billing_budgets_v1_budget_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBudgetRequest.ProtoReflect.Descriptor instead.
func (*GetBudgetRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_billing_budgets_v1_budget_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetBudgetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request for ListBudgets
type ListBudgetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of billing account to list budgets under. Values
	// are of the form `billingAccounts/{billingAccountId}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. Set the scope of the budgets to be returned, in the format of the
	// resource name. The scope of a budget is the cost that it tracks, such as
	// costs for a single project, or the costs for all projects in a folder. Only
	// project scope (in the format of "projects/project-id" or "projects/123") is
	// supported in this field. When this field is set to a project's resource
	// name, the budgets returned are tracking the costs for that project.
	Scope string `protobuf:"bytes,4,opt,name=scope,proto3" json:"scope,omitempty"`
	// Optional. The maximum number of budgets to return per page.
	// The default and maximum value are 100.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The value returned by the last `ListBudgetsResponse` which
	// indicates that this is a continuation of a prior `ListBudgets` call,
	// and that the system should return the next page of data.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListBudgetsRequest) Reset() {
	*x = ListBudgetsRequest{}
	mi := &file_google_cloud_billing_budgets_v1_budget_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBudgetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBudgetsRequest) ProtoMessage() {}

func (x *ListBudgetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_billing_budgets_v1_budget_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBudgetsRequest.ProtoReflect.Descriptor instead.
func (*ListBudgetsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_billing_budgets_v1_budget_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListBudgetsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListBudgetsRequest) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *ListBudgetsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListBudgetsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response for ListBudgets
type ListBudgetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of the budgets owned by the requested billing account.
	Budgets []*Budget `protobuf:"bytes,1,rep,name=budgets,proto3" json:"budgets,omitempty"`
	// If not empty, indicates that there may be more budgets that match the
	// request; this value should be passed in a new `ListBudgetsRequest`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListBudgetsResponse) Reset() {
	*x = ListBudgetsResponse{}
	mi := &file_google_cloud_billing_budgets_v1_budget_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBudgetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBudgetsResponse) ProtoMessage() {}

func (x *ListBudgetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_billing_budgets_v1_budget_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBudgetsResponse.ProtoReflect.Descriptor instead.
func (*ListBudgetsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_billing_budgets_v1_budget_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListBudgetsResponse) GetBudgets() []*Budget {
	if x != nil {
		return x.Budgets
	}
	return nil
}

func (x *ListBudgetsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request for DeleteBudget
type DeleteBudgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the budget to delete. Values are of the form
	// `billingAccounts/{billingAccountId}/budgets/{budgetId}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteBudgetRequest) Reset() {
	*x = DeleteBudgetRequest{}
	mi := &file_google_cloud_billing_budgets_v1_budget_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBudgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBudgetRequest) ProtoMessage() {}

func (x *DeleteBudgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_billing_budgets_v1_budget_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBudgetRequest.ProtoReflect.Descriptor instead.
func (*DeleteBudgetRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_billing_budgets_v1_budget_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteBudgetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_cloud_billing_budgets_v1_budget_service_proto protoreflect.FileDescriptor

var file_google_cloud_billing_budgets_v1_budget_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x26, 0x12, 0x24, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x06, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x22, 0x9d, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x06, 0x62, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12,
	0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73,
	0x6b, 0x22, 0x54, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xbb, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c,
	0xe0, 0x41, 0x02, 0xfa, 0x41, 0x26, 0x12, 0x24, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x80, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a,
	0x07, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x07, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x57, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x40, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0,
	0x41, 0x02, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x62, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x32, 0xf7, 0x07, 0x0a, 0x0d, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xb5, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x12, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x22, 0x46, 0xda, 0x41, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x3a, 0x06, 0x62, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x22, 0x26, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x3d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x2f, 0x2a, 0x7d, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x12, 0xc1, 0x01, 0x0a, 0x0c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x34, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x22, 0x52, 0xda, 0x41, 0x12,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x3a, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x32, 0x2d, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x12,
	0x9e, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x31, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x22, 0x35, 0xda, 0x41, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x3d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x2a, 0x7d,
	0x12, 0xb1, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73,
	0x12, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0xda, 0x41, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26, 0x2f, 0x76,
	0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x62, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x73, 0x12, 0x93, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x35, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x28, 0x2a, 0x26, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x1a, 0x7f, 0xca, 0x41, 0x1d, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x5c, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2c, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x7a, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0x12, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x70, 0x62, 0x3b, 0x62, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_billing_budgets_v1_budget_service_proto_rawDescOnce sync.Once
	file_google_cloud_billing_budgets_v1_budget_service_proto_rawDescData = file_google_cloud_billing_budgets_v1_budget_service_proto_rawDesc
)

func file_google_cloud_billing_budgets_v1_budget_service_proto_rawDescGZIP() []byte {
	file_google_cloud_billing_budgets_v1_budget_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_billing_budgets_v1_budget_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_billing_budgets_v1_budget_service_proto_rawDescData)
	})
	return file_google_cloud_billing_budgets_v1_budget_service_proto_rawDescData
}

var file_google_cloud_billing_budgets_v1_budget_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_billing_budgets_v1_budget_service_proto_goTypes = []any{
	(*CreateBudgetRequest)(nil),   // 0: google.cloud.billing.budgets.v1.CreateBudgetRequest
	(*UpdateBudgetRequest)(nil),   // 1: google.cloud.billing.budgets.v1.UpdateBudgetRequest
	(*GetBudgetRequest)(nil),      // 2: google.cloud.billing.budgets.v1.GetBudgetRequest
	(*ListBudgetsRequest)(nil),    // 3: google.cloud.billing.budgets.v1.ListBudgetsRequest
	(*ListBudgetsResponse)(nil),   // 4: google.cloud.billing.budgets.v1.ListBudgetsResponse
	(*DeleteBudgetRequest)(nil),   // 5: google.cloud.billing.budgets.v1.DeleteBudgetRequest
	(*Budget)(nil),                // 6: google.cloud.billing.budgets.v1.Budget
	(*fieldmaskpb.FieldMask)(nil), // 7: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_google_cloud_billing_budgets_v1_budget_service_proto_depIdxs = []int32{
	6, // 0: google.cloud.billing.budgets.v1.CreateBudgetRequest.budget:type_name -> google.cloud.billing.budgets.v1.Budget
	6, // 1: google.cloud.billing.budgets.v1.UpdateBudgetRequest.budget:type_name -> google.cloud.billing.budgets.v1.Budget
	7, // 2: google.cloud.billing.budgets.v1.UpdateBudgetRequest.update_mask:type_name -> google.protobuf.FieldMask
	6, // 3: google.cloud.billing.budgets.v1.ListBudgetsResponse.budgets:type_name -> google.cloud.billing.budgets.v1.Budget
	0, // 4: google.cloud.billing.budgets.v1.BudgetService.CreateBudget:input_type -> google.cloud.billing.budgets.v1.CreateBudgetRequest
	1, // 5: google.cloud.billing.budgets.v1.BudgetService.UpdateBudget:input_type -> google.cloud.billing.budgets.v1.UpdateBudgetRequest
	2, // 6: google.cloud.billing.budgets.v1.BudgetService.GetBudget:input_type -> google.cloud.billing.budgets.v1.GetBudgetRequest
	3, // 7: google.cloud.billing.budgets.v1.BudgetService.ListBudgets:input_type -> google.cloud.billing.budgets.v1.ListBudgetsRequest
	5, // 8: google.cloud.billing.budgets.v1.BudgetService.DeleteBudget:input_type -> google.cloud.billing.budgets.v1.DeleteBudgetRequest
	6, // 9: google.cloud.billing.budgets.v1.BudgetService.CreateBudget:output_type -> google.cloud.billing.budgets.v1.Budget
	6, // 10: google.cloud.billing.budgets.v1.BudgetService.UpdateBudget:output_type -> google.cloud.billing.budgets.v1.Budget
	6, // 11: google.cloud.billing.budgets.v1.BudgetService.GetBudget:output_type -> google.cloud.billing.budgets.v1.Budget
	4, // 12: google.cloud.billing.budgets.v1.BudgetService.ListBudgets:output_type -> google.cloud.billing.budgets.v1.ListBudgetsResponse
	8, // 13: google.cloud.billing.budgets.v1.BudgetService.DeleteBudget:output_type -> google.protobuf.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_billing_budgets_v1_budget_service_proto_init() }
func file_google_cloud_billing_budgets_v1_budget_service_proto_init() {
	if File_google_cloud_billing_budgets_v1_budget_service_proto != nil {
		return
	}
	file_google_cloud_billing_budgets_v1_budget_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_billing_budgets_v1_budget_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_billing_budgets_v1_budget_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_billing_budgets_v1_budget_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_billing_budgets_v1_budget_service_proto_msgTypes,
	}.Build()
	File_google_cloud_billing_budgets_v1_budget_service_proto = out.File
	file_google_cloud_billing_budgets_v1_budget_service_proto_rawDesc = nil
	file_google_cloud_billing_budgets_v1_budget_service_proto_goTypes = nil
	file_google_cloud_billing_budgets_v1_budget_service_proto_depIdxs = nil
}
