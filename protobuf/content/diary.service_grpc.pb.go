// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package content

import (
	context "context"
	empty "github.com/liov/tiga/protobuf/utils/empty"
	request "github.com/liov/tiga/protobuf/utils/request"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DiaryServiceClient is the client API for DiaryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiaryServiceClient interface {
	// 日记本
	DiaryBook(ctx context.Context, in *DiaryBookReq, opts ...grpc.CallOption) (*DiaryBookRep, error)
	// 日记本列表
	DiaryBookList(ctx context.Context, in *DiaryBookListReq, opts ...grpc.CallOption) (*DiaryBookListRep, error)
	// 创建日记本
	AddDiaryBook(ctx context.Context, in *AddDiaryBookReq, opts ...grpc.CallOption) (*request.Object, error)
	// 修改日记本
	EditDiaryBook(ctx context.Context, in *AddDiaryBookReq, opts ...grpc.CallOption) (*empty.Empty, error)
	// 详情
	Info(ctx context.Context, in *request.Object, opts ...grpc.CallOption) (*Diary, error)
	// 新建
	Add(ctx context.Context, in *AddDiaryReq, opts ...grpc.CallOption) (*request.Object, error)
	// 修改
	Edit(ctx context.Context, in *AddDiaryReq, opts ...grpc.CallOption) (*empty.Empty, error)
	// 列表
	List(ctx context.Context, in *DiaryListReq, opts ...grpc.CallOption) (*DiaryListRep, error)
	// 删除
	Delete(ctx context.Context, in *request.Object, opts ...grpc.CallOption) (*empty.Empty, error)
}

type diaryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiaryServiceClient(cc grpc.ClientConnInterface) DiaryServiceClient {
	return &diaryServiceClient{cc}
}

func (c *diaryServiceClient) DiaryBook(ctx context.Context, in *DiaryBookReq, opts ...grpc.CallOption) (*DiaryBookRep, error) {
	out := new(DiaryBookRep)
	err := c.cc.Invoke(ctx, "/content.DiaryService/DiaryBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diaryServiceClient) DiaryBookList(ctx context.Context, in *DiaryBookListReq, opts ...grpc.CallOption) (*DiaryBookListRep, error) {
	out := new(DiaryBookListRep)
	err := c.cc.Invoke(ctx, "/content.DiaryService/DiaryBookList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diaryServiceClient) AddDiaryBook(ctx context.Context, in *AddDiaryBookReq, opts ...grpc.CallOption) (*request.Object, error) {
	out := new(request.Object)
	err := c.cc.Invoke(ctx, "/content.DiaryService/AddDiaryBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diaryServiceClient) EditDiaryBook(ctx context.Context, in *AddDiaryBookReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.DiaryService/EditDiaryBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diaryServiceClient) Info(ctx context.Context, in *request.Object, opts ...grpc.CallOption) (*Diary, error) {
	out := new(Diary)
	err := c.cc.Invoke(ctx, "/content.DiaryService/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diaryServiceClient) Add(ctx context.Context, in *AddDiaryReq, opts ...grpc.CallOption) (*request.Object, error) {
	out := new(request.Object)
	err := c.cc.Invoke(ctx, "/content.DiaryService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diaryServiceClient) Edit(ctx context.Context, in *AddDiaryReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.DiaryService/Edit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diaryServiceClient) List(ctx context.Context, in *DiaryListReq, opts ...grpc.CallOption) (*DiaryListRep, error) {
	out := new(DiaryListRep)
	err := c.cc.Invoke(ctx, "/content.DiaryService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diaryServiceClient) Delete(ctx context.Context, in *request.Object, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.DiaryService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiaryServiceServer is the server API for DiaryService service.
// All implementations must embed UnimplementedDiaryServiceServer
// for forward compatibility
type DiaryServiceServer interface {
	// 日记本
	DiaryBook(context.Context, *DiaryBookReq) (*DiaryBookRep, error)
	// 日记本列表
	DiaryBookList(context.Context, *DiaryBookListReq) (*DiaryBookListRep, error)
	// 创建日记本
	AddDiaryBook(context.Context, *AddDiaryBookReq) (*request.Object, error)
	// 修改日记本
	EditDiaryBook(context.Context, *AddDiaryBookReq) (*empty.Empty, error)
	// 详情
	Info(context.Context, *request.Object) (*Diary, error)
	// 新建
	Add(context.Context, *AddDiaryReq) (*request.Object, error)
	// 修改
	Edit(context.Context, *AddDiaryReq) (*empty.Empty, error)
	// 列表
	List(context.Context, *DiaryListReq) (*DiaryListRep, error)
	// 删除
	Delete(context.Context, *request.Object) (*empty.Empty, error)
	mustEmbedUnimplementedDiaryServiceServer()
}

// UnimplementedDiaryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDiaryServiceServer struct {
}

func (UnimplementedDiaryServiceServer) DiaryBook(context.Context, *DiaryBookReq) (*DiaryBookRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiaryBook not implemented")
}
func (UnimplementedDiaryServiceServer) DiaryBookList(context.Context, *DiaryBookListReq) (*DiaryBookListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiaryBookList not implemented")
}
func (UnimplementedDiaryServiceServer) AddDiaryBook(context.Context, *AddDiaryBookReq) (*request.Object, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDiaryBook not implemented")
}
func (UnimplementedDiaryServiceServer) EditDiaryBook(context.Context, *AddDiaryBookReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditDiaryBook not implemented")
}
func (UnimplementedDiaryServiceServer) Info(context.Context, *request.Object) (*Diary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedDiaryServiceServer) Add(context.Context, *AddDiaryReq) (*request.Object, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedDiaryServiceServer) Edit(context.Context, *AddDiaryReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}
func (UnimplementedDiaryServiceServer) List(context.Context, *DiaryListReq) (*DiaryListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDiaryServiceServer) Delete(context.Context, *request.Object) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDiaryServiceServer) mustEmbedUnimplementedDiaryServiceServer() {}

// UnsafeDiaryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiaryServiceServer will
// result in compilation errors.
type UnsafeDiaryServiceServer interface {
	mustEmbedUnimplementedDiaryServiceServer()
}

func RegisterDiaryServiceServer(s grpc.ServiceRegistrar, srv DiaryServiceServer) {
	s.RegisterService(&DiaryService_ServiceDesc, srv)
}

func _DiaryService_DiaryBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiaryBookReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiaryServiceServer).DiaryBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.DiaryService/DiaryBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiaryServiceServer).DiaryBook(ctx, req.(*DiaryBookReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiaryService_DiaryBookList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiaryBookListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiaryServiceServer).DiaryBookList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.DiaryService/DiaryBookList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiaryServiceServer).DiaryBookList(ctx, req.(*DiaryBookListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiaryService_AddDiaryBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDiaryBookReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiaryServiceServer).AddDiaryBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.DiaryService/AddDiaryBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiaryServiceServer).AddDiaryBook(ctx, req.(*AddDiaryBookReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiaryService_EditDiaryBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDiaryBookReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiaryServiceServer).EditDiaryBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.DiaryService/EditDiaryBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiaryServiceServer).EditDiaryBook(ctx, req.(*AddDiaryBookReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiaryService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiaryServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.DiaryService/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiaryServiceServer).Info(ctx, req.(*request.Object))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiaryService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDiaryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiaryServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.DiaryService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiaryServiceServer).Add(ctx, req.(*AddDiaryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiaryService_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDiaryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiaryServiceServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.DiaryService/Edit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiaryServiceServer).Edit(ctx, req.(*AddDiaryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiaryService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiaryListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiaryServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.DiaryService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiaryServiceServer).List(ctx, req.(*DiaryListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiaryService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiaryServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.DiaryService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiaryServiceServer).Delete(ctx, req.(*request.Object))
	}
	return interceptor(ctx, in, info, handler)
}

// DiaryService_ServiceDesc is the grpc.ServiceDesc for DiaryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiaryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content.DiaryService",
	HandlerType: (*DiaryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DiaryBook",
			Handler:    _DiaryService_DiaryBook_Handler,
		},
		{
			MethodName: "DiaryBookList",
			Handler:    _DiaryService_DiaryBookList_Handler,
		},
		{
			MethodName: "AddDiaryBook",
			Handler:    _DiaryService_AddDiaryBook_Handler,
		},
		{
			MethodName: "EditDiaryBook",
			Handler:    _DiaryService_EditDiaryBook_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _DiaryService_Info_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _DiaryService_Add_Handler,
		},
		{
			MethodName: "Edit",
			Handler:    _DiaryService_Edit_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DiaryService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DiaryService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content/diary.service.proto",
}
