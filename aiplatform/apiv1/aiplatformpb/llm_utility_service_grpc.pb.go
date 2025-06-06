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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.7
// source: google/cloud/aiplatform/v1/llm_utility_service.proto

package aiplatformpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	LlmUtilityService_CountTokens_FullMethodName   = "/google.cloud.aiplatform.v1.LlmUtilityService/CountTokens"
	LlmUtilityService_ComputeTokens_FullMethodName = "/google.cloud.aiplatform.v1.LlmUtilityService/ComputeTokens"
)

// LlmUtilityServiceClient is the client API for LlmUtilityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LlmUtilityServiceClient interface {
	// Perform a token counting.
	CountTokens(ctx context.Context, in *CountTokensRequest, opts ...grpc.CallOption) (*CountTokensResponse, error)
	// Return a list of tokens based on the input text.
	ComputeTokens(ctx context.Context, in *ComputeTokensRequest, opts ...grpc.CallOption) (*ComputeTokensResponse, error)
}

type llmUtilityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLlmUtilityServiceClient(cc grpc.ClientConnInterface) LlmUtilityServiceClient {
	return &llmUtilityServiceClient{cc}
}

func (c *llmUtilityServiceClient) CountTokens(ctx context.Context, in *CountTokensRequest, opts ...grpc.CallOption) (*CountTokensResponse, error) {
	out := new(CountTokensResponse)
	err := c.cc.Invoke(ctx, LlmUtilityService_CountTokens_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *llmUtilityServiceClient) ComputeTokens(ctx context.Context, in *ComputeTokensRequest, opts ...grpc.CallOption) (*ComputeTokensResponse, error) {
	out := new(ComputeTokensResponse)
	err := c.cc.Invoke(ctx, LlmUtilityService_ComputeTokens_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LlmUtilityServiceServer is the server API for LlmUtilityService service.
// All implementations should embed UnimplementedLlmUtilityServiceServer
// for forward compatibility
type LlmUtilityServiceServer interface {
	// Perform a token counting.
	CountTokens(context.Context, *CountTokensRequest) (*CountTokensResponse, error)
	// Return a list of tokens based on the input text.
	ComputeTokens(context.Context, *ComputeTokensRequest) (*ComputeTokensResponse, error)
}

// UnimplementedLlmUtilityServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLlmUtilityServiceServer struct {
}

func (UnimplementedLlmUtilityServiceServer) CountTokens(context.Context, *CountTokensRequest) (*CountTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountTokens not implemented")
}
func (UnimplementedLlmUtilityServiceServer) ComputeTokens(context.Context, *ComputeTokensRequest) (*ComputeTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeTokens not implemented")
}

// UnsafeLlmUtilityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LlmUtilityServiceServer will
// result in compilation errors.
type UnsafeLlmUtilityServiceServer interface {
	mustEmbedUnimplementedLlmUtilityServiceServer()
}

func RegisterLlmUtilityServiceServer(s grpc.ServiceRegistrar, srv LlmUtilityServiceServer) {
	s.RegisterService(&LlmUtilityService_ServiceDesc, srv)
}

func _LlmUtilityService_CountTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LlmUtilityServiceServer).CountTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LlmUtilityService_CountTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LlmUtilityServiceServer).CountTokens(ctx, req.(*CountTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LlmUtilityService_ComputeTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputeTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LlmUtilityServiceServer).ComputeTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LlmUtilityService_ComputeTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LlmUtilityServiceServer).ComputeTokens(ctx, req.(*ComputeTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LlmUtilityService_ServiceDesc is the grpc.ServiceDesc for LlmUtilityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LlmUtilityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.aiplatform.v1.LlmUtilityService",
	HandlerType: (*LlmUtilityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CountTokens",
			Handler:    _LlmUtilityService_CountTokens_Handler,
		},
		{
			MethodName: "ComputeTokens",
			Handler:    _LlmUtilityService_ComputeTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/aiplatform/v1/llm_utility_service.proto",
}
