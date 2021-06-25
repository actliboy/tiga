// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

import (
	context "context"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	empty "github.com/liov/tiga/protobuf/utils/empty"
	oauth "github.com/liov/tiga/protobuf/utils/oauth"
	request "github.com/liov/tiga/protobuf/utils/request"
	response "github.com/liov/tiga/protobuf/utils/response"
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
	// 验证码
	VerifyCode(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*wrappers.StringValue, error)
	// 验证码
	SendVerifyCode(ctx context.Context, in *SendVerifyCodeReq, opts ...grpc.CallOption) (*empty.Empty, error)
	//注册验证
	SignupVerify(ctx context.Context, in *SingUpVerifyReq, opts ...grpc.CallOption) (*empty.Empty, error)
	//注册
	Signup(ctx context.Context, in *SignupReq, opts ...grpc.CallOption) (*empty.Empty, error)
	//注册
	EasySignup(ctx context.Context, in *SignupReq, opts ...grpc.CallOption) (*LoginRep, error)
	//激活
	Active(ctx context.Context, in *ActiveReq, opts ...grpc.CallOption) (*LoginRep, error)
	//编辑
	Edit(ctx context.Context, in *EditReq, opts ...grpc.CallOption) (*empty.Empty, error)
	//登录
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRep, error)
	//退出
	Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	//鉴权
	AuthInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserAuthInfo, error)
	//重置密码
	ForgetPassword(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*response.TinyRep, error)
	ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*response.TinyRep, error)
	//获取用户信息
	Info(ctx context.Context, in *request.Object, opts ...grpc.CallOption) (*UserRep, error)
	ActionLogList(ctx context.Context, in *ActionLogListReq, opts ...grpc.CallOption) (*ActionLogListRep, error)
	//获取基础用户信息
	BaseList(ctx context.Context, in *BaseListReq, opts ...grpc.CallOption) (*BaseListRep, error)
	// 关注
	Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*empty.Empty, error)
	// 取消关注
	DelFollow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*BaseListRep, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) VerifyCode(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*wrappers.StringValue, error) {
	out := new(wrappers.StringValue)
	err := c.cc.Invoke(ctx, "/user.UserService/VerifyCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SendVerifyCode(ctx context.Context, in *SendVerifyCodeReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/SendVerifyCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SignupVerify(ctx context.Context, in *SingUpVerifyReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/SignupVerify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Signup(ctx context.Context, in *SignupReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EasySignup(ctx context.Context, in *SignupReq, opts ...grpc.CallOption) (*LoginRep, error) {
	out := new(LoginRep)
	err := c.cc.Invoke(ctx, "/user.UserService/EasySignup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Active(ctx context.Context, in *ActiveReq, opts ...grpc.CallOption) (*LoginRep, error) {
	out := new(LoginRep)
	err := c.cc.Invoke(ctx, "/user.UserService/Active", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Edit(ctx context.Context, in *EditReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/Edit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRep, error) {
	out := new(LoginRep)
	err := c.cc.Invoke(ctx, "/user.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AuthInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserAuthInfo, error) {
	out := new(UserAuthInfo)
	err := c.cc.Invoke(ctx, "/user.UserService/AuthInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ForgetPassword(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*response.TinyRep, error) {
	out := new(response.TinyRep)
	err := c.cc.Invoke(ctx, "/user.UserService/ForgetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*response.TinyRep, error) {
	out := new(response.TinyRep)
	err := c.cc.Invoke(ctx, "/user.UserService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Info(ctx context.Context, in *request.Object, opts ...grpc.CallOption) (*UserRep, error) {
	out := new(UserRep)
	err := c.cc.Invoke(ctx, "/user.UserService/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ActionLogList(ctx context.Context, in *ActionLogListReq, opts ...grpc.CallOption) (*ActionLogListRep, error) {
	out := new(ActionLogListRep)
	err := c.cc.Invoke(ctx, "/user.UserService/ActionLogList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) BaseList(ctx context.Context, in *BaseListReq, opts ...grpc.CallOption) (*BaseListRep, error) {
	out := new(BaseListRep)
	err := c.cc.Invoke(ctx, "/user.UserService/BaseList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/Follow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DelFollow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*BaseListRep, error) {
	out := new(BaseListRep)
	err := c.cc.Invoke(ctx, "/user.UserService/delFollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// 验证码
	VerifyCode(context.Context, *empty.Empty) (*wrappers.StringValue, error)
	// 验证码
	SendVerifyCode(context.Context, *SendVerifyCodeReq) (*empty.Empty, error)
	//注册验证
	SignupVerify(context.Context, *SingUpVerifyReq) (*empty.Empty, error)
	//注册
	Signup(context.Context, *SignupReq) (*empty.Empty, error)
	//注册
	EasySignup(context.Context, *SignupReq) (*LoginRep, error)
	//激活
	Active(context.Context, *ActiveReq) (*LoginRep, error)
	//编辑
	Edit(context.Context, *EditReq) (*empty.Empty, error)
	//登录
	Login(context.Context, *LoginReq) (*LoginRep, error)
	//退出
	Logout(context.Context, *empty.Empty) (*empty.Empty, error)
	//鉴权
	AuthInfo(context.Context, *empty.Empty) (*UserAuthInfo, error)
	//重置密码
	ForgetPassword(context.Context, *LoginReq) (*response.TinyRep, error)
	ResetPassword(context.Context, *ResetPasswordReq) (*response.TinyRep, error)
	//获取用户信息
	Info(context.Context, *request.Object) (*UserRep, error)
	ActionLogList(context.Context, *ActionLogListReq) (*ActionLogListRep, error)
	//获取基础用户信息
	BaseList(context.Context, *BaseListReq) (*BaseListRep, error)
	// 关注
	Follow(context.Context, *FollowReq) (*empty.Empty, error)
	// 取消关注
	DelFollow(context.Context, *FollowReq) (*BaseListRep, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) VerifyCode(context.Context, *empty.Empty) (*wrappers.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCode not implemented")
}
func (UnimplementedUserServiceServer) SendVerifyCode(context.Context, *SendVerifyCodeReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerifyCode not implemented")
}
func (UnimplementedUserServiceServer) SignupVerify(context.Context, *SingUpVerifyReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignupVerify not implemented")
}
func (UnimplementedUserServiceServer) Signup(context.Context, *SignupReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedUserServiceServer) EasySignup(context.Context, *SignupReq) (*LoginRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EasySignup not implemented")
}
func (UnimplementedUserServiceServer) Active(context.Context, *ActiveReq) (*LoginRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Active not implemented")
}
func (UnimplementedUserServiceServer) Edit(context.Context, *EditReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginReq) (*LoginRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) Logout(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserServiceServer) AuthInfo(context.Context, *empty.Empty) (*UserAuthInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthInfo not implemented")
}
func (UnimplementedUserServiceServer) ForgetPassword(context.Context, *LoginReq) (*response.TinyRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgetPassword not implemented")
}
func (UnimplementedUserServiceServer) ResetPassword(context.Context, *ResetPasswordReq) (*response.TinyRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedUserServiceServer) Info(context.Context, *request.Object) (*UserRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedUserServiceServer) ActionLogList(context.Context, *ActionLogListReq) (*ActionLogListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionLogList not implemented")
}
func (UnimplementedUserServiceServer) BaseList(context.Context, *BaseListReq) (*BaseListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BaseList not implemented")
}
func (UnimplementedUserServiceServer) Follow(context.Context, *FollowReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedUserServiceServer) DelFollow(context.Context, *FollowReq) (*BaseListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelFollow not implemented")
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

func _UserService_VerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).VerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/VerifyCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).VerifyCode(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SendVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVerifyCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SendVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/SendVerifyCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SendVerifyCode(ctx, req.(*SendVerifyCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SignupVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingUpVerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignupVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/SignupVerify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignupVerify(ctx, req.(*SingUpVerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Signup(ctx, req.(*SignupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EasySignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EasySignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/EasySignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EasySignup(ctx, req.(*SignupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Active_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Active(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Active",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Active(ctx, req.(*ActiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Edit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Edit(ctx, req.(*EditReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Logout(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AuthInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AuthInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/AuthInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AuthInfo(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ForgetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ForgetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ForgetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ForgetPassword(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ResetPassword(ctx, req.(*ResetPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Info(ctx, req.(*request.Object))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ActionLogList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionLogListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ActionLogList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ActionLogList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ActionLogList(ctx, req.(*ActionLogListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_BaseList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).BaseList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/BaseList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).BaseList(ctx, req.(*BaseListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Follow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Follow(ctx, req.(*FollowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DelFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DelFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/delFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DelFollow(ctx, req.(*FollowReq))
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
			MethodName: "VerifyCode",
			Handler:    _UserService_VerifyCode_Handler,
		},
		{
			MethodName: "SendVerifyCode",
			Handler:    _UserService_SendVerifyCode_Handler,
		},
		{
			MethodName: "SignupVerify",
			Handler:    _UserService_SignupVerify_Handler,
		},
		{
			MethodName: "Signup",
			Handler:    _UserService_Signup_Handler,
		},
		{
			MethodName: "EasySignup",
			Handler:    _UserService_EasySignup_Handler,
		},
		{
			MethodName: "Active",
			Handler:    _UserService_Active_Handler,
		},
		{
			MethodName: "Edit",
			Handler:    _UserService_Edit_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserService_Logout_Handler,
		},
		{
			MethodName: "AuthInfo",
			Handler:    _UserService_AuthInfo_Handler,
		},
		{
			MethodName: "ForgetPassword",
			Handler:    _UserService_ForgetPassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _UserService_ResetPassword_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _UserService_Info_Handler,
		},
		{
			MethodName: "ActionLogList",
			Handler:    _UserService_ActionLogList_Handler,
		},
		{
			MethodName: "BaseList",
			Handler:    _UserService_BaseList_Handler,
		},
		{
			MethodName: "Follow",
			Handler:    _UserService_Follow_Handler,
		},
		{
			MethodName: "delFollow",
			Handler:    _UserService_DelFollow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.service.proto",
}

// OauthServiceClient is the client API for OauthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OauthServiceClient interface {
	OauthAuthorize(ctx context.Context, in *oauth.OauthReq, opts ...grpc.CallOption) (*response.HttpResponse, error)
	OauthToken(ctx context.Context, in *oauth.OauthReq, opts ...grpc.CallOption) (*response.HttpResponse, error)
}

type oauthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOauthServiceClient(cc grpc.ClientConnInterface) OauthServiceClient {
	return &oauthServiceClient{cc}
}

func (c *oauthServiceClient) OauthAuthorize(ctx context.Context, in *oauth.OauthReq, opts ...grpc.CallOption) (*response.HttpResponse, error) {
	out := new(response.HttpResponse)
	err := c.cc.Invoke(ctx, "/user.OauthService/OauthAuthorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauthServiceClient) OauthToken(ctx context.Context, in *oauth.OauthReq, opts ...grpc.CallOption) (*response.HttpResponse, error) {
	out := new(response.HttpResponse)
	err := c.cc.Invoke(ctx, "/user.OauthService/OauthToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OauthServiceServer is the server API for OauthService service.
// All implementations must embed UnimplementedOauthServiceServer
// for forward compatibility
type OauthServiceServer interface {
	OauthAuthorize(context.Context, *oauth.OauthReq) (*response.HttpResponse, error)
	OauthToken(context.Context, *oauth.OauthReq) (*response.HttpResponse, error)
	mustEmbedUnimplementedOauthServiceServer()
}

// UnimplementedOauthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOauthServiceServer struct {
}

func (UnimplementedOauthServiceServer) OauthAuthorize(context.Context, *oauth.OauthReq) (*response.HttpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OauthAuthorize not implemented")
}
func (UnimplementedOauthServiceServer) OauthToken(context.Context, *oauth.OauthReq) (*response.HttpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OauthToken not implemented")
}
func (UnimplementedOauthServiceServer) mustEmbedUnimplementedOauthServiceServer() {}

// UnsafeOauthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OauthServiceServer will
// result in compilation errors.
type UnsafeOauthServiceServer interface {
	mustEmbedUnimplementedOauthServiceServer()
}

func RegisterOauthServiceServer(s grpc.ServiceRegistrar, srv OauthServiceServer) {
	s.RegisterService(&OauthService_ServiceDesc, srv)
}

func _OauthService_OauthAuthorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(oauth.OauthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OauthServiceServer).OauthAuthorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.OauthService/OauthAuthorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OauthServiceServer).OauthAuthorize(ctx, req.(*oauth.OauthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OauthService_OauthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(oauth.OauthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OauthServiceServer).OauthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.OauthService/OauthToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OauthServiceServer).OauthToken(ctx, req.(*oauth.OauthReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OauthService_ServiceDesc is the grpc.ServiceDesc for OauthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OauthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.OauthService",
	HandlerType: (*OauthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OauthAuthorize",
			Handler:    _OauthService_OauthAuthorize_Handler,
		},
		{
			MethodName: "OauthToken",
			Handler:    _OauthService_OauthToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.service.proto",
}