//
// Copyright 2024 The Chainloop Authors.
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
// - protoc             (unknown)
// source: controlplane/v1/user.proto

package v1

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
	UserService_ListMemberships_FullMethodName      = "/controlplane.v1.UserService/ListMemberships"
	UserService_DeleteMembership_FullMethodName     = "/controlplane.v1.UserService/DeleteMembership"
	UserService_SetCurrentMembership_FullMethodName = "/controlplane.v1.UserService/SetCurrentMembership"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// User Membership management
	ListMemberships(ctx context.Context, in *UserServiceListMembershipsRequest, opts ...grpc.CallOption) (*UserServiceListMembershipsResponse, error)
	DeleteMembership(ctx context.Context, in *DeleteMembershipRequest, opts ...grpc.CallOption) (*DeleteMembershipResponse, error)
	SetCurrentMembership(ctx context.Context, in *SetCurrentMembershipRequest, opts ...grpc.CallOption) (*SetCurrentMembershipResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) ListMemberships(ctx context.Context, in *UserServiceListMembershipsRequest, opts ...grpc.CallOption) (*UserServiceListMembershipsResponse, error) {
	out := new(UserServiceListMembershipsResponse)
	err := c.cc.Invoke(ctx, UserService_ListMemberships_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteMembership(ctx context.Context, in *DeleteMembershipRequest, opts ...grpc.CallOption) (*DeleteMembershipResponse, error) {
	out := new(DeleteMembershipResponse)
	err := c.cc.Invoke(ctx, UserService_DeleteMembership_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetCurrentMembership(ctx context.Context, in *SetCurrentMembershipRequest, opts ...grpc.CallOption) (*SetCurrentMembershipResponse, error) {
	out := new(SetCurrentMembershipResponse)
	err := c.cc.Invoke(ctx, UserService_SetCurrentMembership_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// User Membership management
	ListMemberships(context.Context, *UserServiceListMembershipsRequest) (*UserServiceListMembershipsResponse, error)
	DeleteMembership(context.Context, *DeleteMembershipRequest) (*DeleteMembershipResponse, error)
	SetCurrentMembership(context.Context, *SetCurrentMembershipRequest) (*SetCurrentMembershipResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) ListMemberships(context.Context, *UserServiceListMembershipsRequest) (*UserServiceListMembershipsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMemberships not implemented")
}
func (UnimplementedUserServiceServer) DeleteMembership(context.Context, *DeleteMembershipRequest) (*DeleteMembershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMembership not implemented")
}
func (UnimplementedUserServiceServer) SetCurrentMembership(context.Context, *SetCurrentMembershipRequest) (*SetCurrentMembershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCurrentMembership not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_ListMemberships_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserServiceListMembershipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListMemberships(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ListMemberships_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListMemberships(ctx, req.(*UserServiceListMembershipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMembershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteMembership_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteMembership(ctx, req.(*DeleteMembershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetCurrentMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCurrentMembershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetCurrentMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SetCurrentMembership_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetCurrentMembership(ctx, req.(*SetCurrentMembershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "controlplane.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMemberships",
			Handler:    _UserService_ListMemberships_Handler,
		},
		{
			MethodName: "DeleteMembership",
			Handler:    _UserService_DeleteMembership_Handler,
		},
		{
			MethodName: "SetCurrentMembership",
			Handler:    _UserService_SetCurrentMembership_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controlplane/v1/user.proto",
}