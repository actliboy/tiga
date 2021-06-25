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

// ActionServiceClient is the client API for ActionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActionServiceClient interface {
	// 动作  Like Unlike
	Like(ctx context.Context, in *LikeReq, opts ...grpc.CallOption) (*request.Object, error)
	// 动作  Like Unlike
	DelLike(ctx context.Context, in *request.Object, opts ...grpc.CallOption) (*empty.Empty, error)
	// 评论
	Comment(ctx context.Context, in *CommentReq, opts ...grpc.CallOption) (*request.Object, error)
	// 评论列表
	CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListRep, error)
	// 评论
	DelComment(ctx context.Context, in *request.Object, opts ...grpc.CallOption) (*empty.Empty, error)
	// 收藏
	Collect(ctx context.Context, in *CollectReq, opts ...grpc.CallOption) (*empty.Empty, error)
	// 举报
	Report(ctx context.Context, in *ReportReq, opts ...grpc.CallOption) (*empty.Empty, error)
}

type actionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActionServiceClient(cc grpc.ClientConnInterface) ActionServiceClient {
	return &actionServiceClient{cc}
}

func (c *actionServiceClient) Like(ctx context.Context, in *LikeReq, opts ...grpc.CallOption) (*request.Object, error) {
	out := new(request.Object)
	err := c.cc.Invoke(ctx, "/content.ActionService/Like", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) DelLike(ctx context.Context, in *request.Object, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.ActionService/DelLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) Comment(ctx context.Context, in *CommentReq, opts ...grpc.CallOption) (*request.Object, error) {
	out := new(request.Object)
	err := c.cc.Invoke(ctx, "/content.ActionService/Comment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListRep, error) {
	out := new(CommentListRep)
	err := c.cc.Invoke(ctx, "/content.ActionService/CommentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) DelComment(ctx context.Context, in *request.Object, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.ActionService/DelComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) Collect(ctx context.Context, in *CollectReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.ActionService/Collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) Report(ctx context.Context, in *ReportReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.ActionService/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionServiceServer is the server API for ActionService service.
// All implementations must embed UnimplementedActionServiceServer
// for forward compatibility
type ActionServiceServer interface {
	// 动作  Like Unlike
	Like(context.Context, *LikeReq) (*request.Object, error)
	// 动作  Like Unlike
	DelLike(context.Context, *request.Object) (*empty.Empty, error)
	// 评论
	Comment(context.Context, *CommentReq) (*request.Object, error)
	// 评论列表
	CommentList(context.Context, *CommentListReq) (*CommentListRep, error)
	// 评论
	DelComment(context.Context, *request.Object) (*empty.Empty, error)
	// 收藏
	Collect(context.Context, *CollectReq) (*empty.Empty, error)
	// 举报
	Report(context.Context, *ReportReq) (*empty.Empty, error)
	mustEmbedUnimplementedActionServiceServer()
}

// UnimplementedActionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedActionServiceServer struct {
}

func (UnimplementedActionServiceServer) Like(context.Context, *LikeReq) (*request.Object, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (UnimplementedActionServiceServer) DelLike(context.Context, *request.Object) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelLike not implemented")
}
func (UnimplementedActionServiceServer) Comment(context.Context, *CommentReq) (*request.Object, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comment not implemented")
}
func (UnimplementedActionServiceServer) CommentList(context.Context, *CommentListReq) (*CommentListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentList not implemented")
}
func (UnimplementedActionServiceServer) DelComment(context.Context, *request.Object) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelComment not implemented")
}
func (UnimplementedActionServiceServer) Collect(context.Context, *CollectReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}
func (UnimplementedActionServiceServer) Report(context.Context, *ReportReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}
func (UnimplementedActionServiceServer) mustEmbedUnimplementedActionServiceServer() {}

// UnsafeActionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActionServiceServer will
// result in compilation errors.
type UnsafeActionServiceServer interface {
	mustEmbedUnimplementedActionServiceServer()
}

func RegisterActionServiceServer(s grpc.ServiceRegistrar, srv ActionServiceServer) {
	s.RegisterService(&ActionService_ServiceDesc, srv)
}

func _ActionService_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ActionService/Like",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).Like(ctx, req.(*LikeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_DelLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).DelLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ActionService/DelLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).DelLike(ctx, req.(*request.Object))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_Comment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).Comment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ActionService/Comment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).Comment(ctx, req.(*CommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_CommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).CommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ActionService/CommentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).CommentList(ctx, req.(*CommentListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_DelComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).DelComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ActionService/DelComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).DelComment(ctx, req.(*request.Object))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ActionService/Collect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).Collect(ctx, req.(*CollectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ActionService/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).Report(ctx, req.(*ReportReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ActionService_ServiceDesc is the grpc.ServiceDesc for ActionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content.ActionService",
	HandlerType: (*ActionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Like",
			Handler:    _ActionService_Like_Handler,
		},
		{
			MethodName: "DelLike",
			Handler:    _ActionService_DelLike_Handler,
		},
		{
			MethodName: "Comment",
			Handler:    _ActionService_Comment_Handler,
		},
		{
			MethodName: "CommentList",
			Handler:    _ActionService_CommentList_Handler,
		},
		{
			MethodName: "DelComment",
			Handler:    _ActionService_DelComment_Handler,
		},
		{
			MethodName: "Collect",
			Handler:    _ActionService_Collect_Handler,
		},
		{
			MethodName: "Report",
			Handler:    _ActionService_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content/action.service.proto",
}