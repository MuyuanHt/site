// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: protocol/user/user.proto

package user

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// 查询单个用户
	FindOneUser(ctx context.Context, in *FindOneUserReq, opts ...grpc.CallOption) (*FindOneUserResp, error)
	// 查询单个账户
	FindAccount(ctx context.Context, in *FindAccountReq, opts ...grpc.CallOption) (*FindAccountResp, error)
	// 新增用户
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error)
	// 修改密码
	UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordResp, error)
	// 修改用户信息
	UpdateUserInfo(ctx context.Context, in *ChangeUserInfoReq, opts ...grpc.CallOption) (*ChangeUserInfoResp, error)
	// 模糊查询用户
	FuzzyQueryUsers(ctx context.Context, in *FuzzyQueryUsersReq, opts ...grpc.CallOption) (*FuzzyQueryUsersResp, error)
	// 修改最后登录时间
	UpdateLastLoginTime(ctx context.Context, in *GeneralReq, opts ...grpc.CallOption) (*GeneralResp, error)
	// 修改用户隐私权限
	UpdateUserLimit(ctx context.Context, in *UpdateUserLimitReq, opts ...grpc.CallOption) (*UpdateUserLimitResp, error)
	// 修改用户好友关系数量
	UpdateUserRelation(ctx context.Context, in *UpdateUserRelationReq, opts ...grpc.CallOption) (*UpdateUserRelationResp, error)
	// 详细的好友关系操作
	// 添加或删除好友
	UpdateUserFriendNum(ctx context.Context, in *UpdateUserFriendNumReq, opts ...grpc.CallOption) (*UpdateUserFriendNumResp, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) FindOneUser(ctx context.Context, in *FindOneUserReq, opts ...grpc.CallOption) (*FindOneUserResp, error) {
	out := new(FindOneUserResp)
	err := c.cc.Invoke(ctx, "/user.UserService/FindOneUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FindAccount(ctx context.Context, in *FindAccountReq, opts ...grpc.CallOption) (*FindAccountResp, error) {
	out := new(FindAccountResp)
	err := c.cc.Invoke(ctx, "/user.UserService/FindAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error) {
	out := new(CreateUserResp)
	err := c.cc.Invoke(ctx, "/user.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordResp, error) {
	out := new(UpdatePasswordResp)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserInfo(ctx context.Context, in *ChangeUserInfoReq, opts ...grpc.CallOption) (*ChangeUserInfoResp, error) {
	out := new(ChangeUserInfoResp)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FuzzyQueryUsers(ctx context.Context, in *FuzzyQueryUsersReq, opts ...grpc.CallOption) (*FuzzyQueryUsersResp, error) {
	out := new(FuzzyQueryUsersResp)
	err := c.cc.Invoke(ctx, "/user.UserService/FuzzyQueryUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateLastLoginTime(ctx context.Context, in *GeneralReq, opts ...grpc.CallOption) (*GeneralResp, error) {
	out := new(GeneralResp)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateLastLoginTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserLimit(ctx context.Context, in *UpdateUserLimitReq, opts ...grpc.CallOption) (*UpdateUserLimitResp, error) {
	out := new(UpdateUserLimitResp)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUserLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserRelation(ctx context.Context, in *UpdateUserRelationReq, opts ...grpc.CallOption) (*UpdateUserRelationResp, error) {
	out := new(UpdateUserRelationResp)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUserRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserFriendNum(ctx context.Context, in *UpdateUserFriendNumReq, opts ...grpc.CallOption) (*UpdateUserFriendNumResp, error) {
	out := new(UpdateUserFriendNumResp)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUserFriendNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// 查询单个用户
	FindOneUser(context.Context, *FindOneUserReq) (*FindOneUserResp, error)
	// 查询单个账户
	FindAccount(context.Context, *FindAccountReq) (*FindAccountResp, error)
	// 新增用户
	CreateUser(context.Context, *CreateUserReq) (*CreateUserResp, error)
	// 修改密码
	UpdatePassword(context.Context, *UpdatePasswordReq) (*UpdatePasswordResp, error)
	// 修改用户信息
	UpdateUserInfo(context.Context, *ChangeUserInfoReq) (*ChangeUserInfoResp, error)
	// 模糊查询用户
	FuzzyQueryUsers(context.Context, *FuzzyQueryUsersReq) (*FuzzyQueryUsersResp, error)
	// 修改最后登录时间
	UpdateLastLoginTime(context.Context, *GeneralReq) (*GeneralResp, error)
	// 修改用户隐私权限
	UpdateUserLimit(context.Context, *UpdateUserLimitReq) (*UpdateUserLimitResp, error)
	// 修改用户好友关系数量
	UpdateUserRelation(context.Context, *UpdateUserRelationReq) (*UpdateUserRelationResp, error)
	// 详细的好友关系操作
	// 添加或删除好友
	UpdateUserFriendNum(context.Context, *UpdateUserFriendNumReq) (*UpdateUserFriendNumResp, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) FindOneUser(context.Context, *FindOneUserReq) (*FindOneUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneUser not implemented")
}
func (UnimplementedUserServiceServer) FindAccount(context.Context, *FindAccountReq) (*FindAccountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAccount not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserReq) (*CreateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) UpdatePassword(context.Context, *UpdatePasswordReq) (*UpdatePasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserInfo(context.Context, *ChangeUserInfoReq) (*ChangeUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedUserServiceServer) FuzzyQueryUsers(context.Context, *FuzzyQueryUsersReq) (*FuzzyQueryUsersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FuzzyQueryUsers not implemented")
}
func (UnimplementedUserServiceServer) UpdateLastLoginTime(context.Context, *GeneralReq) (*GeneralResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLastLoginTime not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserLimit(context.Context, *UpdateUserLimitReq) (*UpdateUserLimitResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserLimit not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserRelation(context.Context, *UpdateUserRelationReq) (*UpdateUserRelationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserRelation not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserFriendNum(context.Context, *UpdateUserFriendNumReq) (*UpdateUserFriendNumResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserFriendNum not implemented")
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

func _UserService_FindOneUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindOneUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FindOneUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindOneUser(ctx, req.(*FindOneUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FindAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FindAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindAccount(ctx, req.(*FindAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePassword(ctx, req.(*UpdatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserInfo(ctx, req.(*ChangeUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FuzzyQueryUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuzzyQueryUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FuzzyQueryUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FuzzyQueryUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FuzzyQueryUsers(ctx, req.(*FuzzyQueryUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateLastLoginTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneralReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateLastLoginTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateLastLoginTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateLastLoginTime(ctx, req.(*GeneralReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserLimitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUserLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserLimit(ctx, req.(*UpdateUserLimitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRelationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUserRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserRelation(ctx, req.(*UpdateUserRelationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserFriendNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserFriendNumReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserFriendNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUserFriendNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserFriendNum(ctx, req.(*UpdateUserFriendNumReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindOneUser",
			Handler:    _UserService_FindOneUser_Handler,
		},
		{
			MethodName: "FindAccount",
			Handler:    _UserService_FindAccount_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserService_UpdatePassword_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _UserService_UpdateUserInfo_Handler,
		},
		{
			MethodName: "FuzzyQueryUsers",
			Handler:    _UserService_FuzzyQueryUsers_Handler,
		},
		{
			MethodName: "UpdateLastLoginTime",
			Handler:    _UserService_UpdateLastLoginTime_Handler,
		},
		{
			MethodName: "UpdateUserLimit",
			Handler:    _UserService_UpdateUserLimit_Handler,
		},
		{
			MethodName: "UpdateUserRelation",
			Handler:    _UserService_UpdateUserRelation_Handler,
		},
		{
			MethodName: "UpdateUserFriendNum",
			Handler:    _UserService_UpdateUserFriendNum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol/user/user.proto",
}
