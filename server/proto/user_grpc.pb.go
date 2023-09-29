// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: proto/user.proto

package proto

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
	UserManager_Fetch_FullMethodName = "/user.UserManager/fetch"
)

// UserManagerClient is the client API for UserManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserManagerClient interface {
	Fetch(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagerClient(cc grpc.ClientConnInterface) UserManagerClient {
	return &userManagerClient{cc}
}

func (c *userManagerClient) Fetch(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserManager_Fetch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagerServer is the server API for UserManager service.
// All implementations should embed UnimplementedUserManagerServer
// for forward compatibility
type UserManagerServer interface {
	Fetch(context.Context, *UserRequest) (*UserResponse, error)
}

// UnimplementedUserManagerServer should be embedded to have forward compatible implementations.
type UnimplementedUserManagerServer struct {
}

func (UnimplementedUserManagerServer) Fetch(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}

// UnsafeUserManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserManagerServer will
// result in compilation errors.
type UnsafeUserManagerServer interface {
	mustEmbedUnimplementedUserManagerServer()
}

func RegisterUserManagerServer(s grpc.ServiceRegistrar, srv UserManagerServer) {
	s.RegisterService(&UserManager_ServiceDesc, srv)
}

func _UserManager_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManager_Fetch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).Fetch(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserManager_ServiceDesc is the grpc.ServiceDesc for UserManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserManager",
	HandlerType: (*UserManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "fetch",
			Handler:    _UserManager_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}