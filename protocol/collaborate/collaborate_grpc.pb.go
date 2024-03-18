// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: protocol/collaborate/collaborate.proto

package collaborate

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

// CollaborateServiceClient is the client API for CollaborateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollaborateServiceClient interface {
	// 添加好友
	AddFriend(ctx context.Context, in *AddFriendReq, opts ...grpc.CallOption) (*AddFriendResp, error)
	// 删除好友
	DeleteFriend(ctx context.Context, in *DeleteFriendReq, opts ...grpc.CallOption) (*DeleteFriendResp, error)
	// 修改好友信息
	UpdateFriend(ctx context.Context, in *UpdateFriendReq, opts ...grpc.CallOption) (*UpdateFriendResp, error)
}

type collaborateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollaborateServiceClient(cc grpc.ClientConnInterface) CollaborateServiceClient {
	return &collaborateServiceClient{cc}
}

func (c *collaborateServiceClient) AddFriend(ctx context.Context, in *AddFriendReq, opts ...grpc.CallOption) (*AddFriendResp, error) {
	out := new(AddFriendResp)
	err := c.cc.Invoke(ctx, "/collaborate.CollaborateService/AddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborateServiceClient) DeleteFriend(ctx context.Context, in *DeleteFriendReq, opts ...grpc.CallOption) (*DeleteFriendResp, error) {
	out := new(DeleteFriendResp)
	err := c.cc.Invoke(ctx, "/collaborate.CollaborateService/DeleteFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborateServiceClient) UpdateFriend(ctx context.Context, in *UpdateFriendReq, opts ...grpc.CallOption) (*UpdateFriendResp, error) {
	out := new(UpdateFriendResp)
	err := c.cc.Invoke(ctx, "/collaborate.CollaborateService/UpdateFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollaborateServiceServer is the server API for CollaborateService service.
// All implementations must embed UnimplementedCollaborateServiceServer
// for forward compatibility
type CollaborateServiceServer interface {
	// 添加好友
	AddFriend(context.Context, *AddFriendReq) (*AddFriendResp, error)
	// 删除好友
	DeleteFriend(context.Context, *DeleteFriendReq) (*DeleteFriendResp, error)
	// 修改好友信息
	UpdateFriend(context.Context, *UpdateFriendReq) (*UpdateFriendResp, error)
	mustEmbedUnimplementedCollaborateServiceServer()
}

// UnimplementedCollaborateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCollaborateServiceServer struct {
}

func (UnimplementedCollaborateServiceServer) AddFriend(context.Context, *AddFriendReq) (*AddFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedCollaborateServiceServer) DeleteFriend(context.Context, *DeleteFriendReq) (*DeleteFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriend not implemented")
}
func (UnimplementedCollaborateServiceServer) UpdateFriend(context.Context, *UpdateFriendReq) (*UpdateFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFriend not implemented")
}
func (UnimplementedCollaborateServiceServer) mustEmbedUnimplementedCollaborateServiceServer() {}

// UnsafeCollaborateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollaborateServiceServer will
// result in compilation errors.
type UnsafeCollaborateServiceServer interface {
	mustEmbedUnimplementedCollaborateServiceServer()
}

func RegisterCollaborateServiceServer(s grpc.ServiceRegistrar, srv CollaborateServiceServer) {
	s.RegisterService(&CollaborateService_ServiceDesc, srv)
}

func _CollaborateService_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborateServiceServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collaborate.CollaborateService/AddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborateServiceServer).AddFriend(ctx, req.(*AddFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaborateService_DeleteFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborateServiceServer).DeleteFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collaborate.CollaborateService/DeleteFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborateServiceServer).DeleteFriend(ctx, req.(*DeleteFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaborateService_UpdateFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborateServiceServer).UpdateFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collaborate.CollaborateService/UpdateFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborateServiceServer).UpdateFriend(ctx, req.(*UpdateFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CollaborateService_ServiceDesc is the grpc.ServiceDesc for CollaborateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CollaborateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "collaborate.CollaborateService",
	HandlerType: (*CollaborateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFriend",
			Handler:    _CollaborateService_AddFriend_Handler,
		},
		{
			MethodName: "DeleteFriend",
			Handler:    _CollaborateService_DeleteFriend_Handler,
		},
		{
			MethodName: "UpdateFriend",
			Handler:    _CollaborateService_UpdateFriend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol/collaborate/collaborate.proto",
}
