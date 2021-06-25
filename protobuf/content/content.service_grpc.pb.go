// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package content

import (
	context "context"
	user "github.com/liov/tiga/protobuf/user"
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

// ContentServiceClient is the client API for ContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentServiceClient interface {
	// 详情
	TagInfo(ctx context.Context, in *request.Object, opts ...grpc.CallOption) (*Tag, error)
	// 新建
	AddTag(ctx context.Context, in *AddTagReq, opts ...grpc.CallOption) (*empty.Empty, error)
	// 修改
	EditTag(ctx context.Context, in *EditTagReq, opts ...grpc.CallOption) (*empty.Empty, error)
	// 列表
	TagList(ctx context.Context, in *TagListReq, opts ...grpc.CallOption) (*TagListRep, error)
	// 列表
	TagGroupList(ctx context.Context, in *TagGroupListReq, opts ...grpc.CallOption) (*TagGroupListRep, error)
	// 详情
	AttrInfo(ctx context.Context, in *request.Object, opts ...grpc.CallOption) (*Attributes, error)
	// 新建
	AddAttr(ctx context.Context, in *AddAttrReq, opts ...grpc.CallOption) (*request.Object, error)
	// 修改
	EditAttr(ctx context.Context, in *EditAttrReq, opts ...grpc.CallOption) (*empty.Empty, error)
	// 列表
	AttrList(ctx context.Context, in *AttrListReq, opts ...grpc.CallOption) (*AttrListRep, error)
	//收藏夹列表
	FavList(ctx context.Context, in *FavListReq, opts ...grpc.CallOption) (*FavListRep, error)
	//收藏夹列表
	TinyFavList(ctx context.Context, in *FavListReq, opts ...grpc.CallOption) (*TinyFavListRep, error)
	// 创建收藏夹
	AddFav(ctx context.Context, in *AddFavReq, opts ...grpc.CallOption) (*request.Object, error)
	// 修改收藏夹
	EditFav(ctx context.Context, in *AddFavReq, opts ...grpc.CallOption) (*empty.Empty, error)
	// 创建合集
	AddContainer(ctx context.Context, in *AddContainerReq, opts ...grpc.CallOption) (*empty.Empty, error)
	// 修改合集
	EditDiaryContainer(ctx context.Context, in *AddContainerReq, opts ...grpc.CallOption) (*empty.Empty, error)
	// 用户内容数量
	UserContentCount(ctx context.Context, in *request.Object, opts ...grpc.CallOption) (*user.UserContent, error)
}

type contentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContentServiceClient(cc grpc.ClientConnInterface) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) TagInfo(ctx context.Context, in *request.Object, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, "/content.ContentService/TagInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) AddTag(ctx context.Context, in *AddTagReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.ContentService/AddTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) EditTag(ctx context.Context, in *EditTagReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.ContentService/EditTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) TagList(ctx context.Context, in *TagListReq, opts ...grpc.CallOption) (*TagListRep, error) {
	out := new(TagListRep)
	err := c.cc.Invoke(ctx, "/content.ContentService/TagList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) TagGroupList(ctx context.Context, in *TagGroupListReq, opts ...grpc.CallOption) (*TagGroupListRep, error) {
	out := new(TagGroupListRep)
	err := c.cc.Invoke(ctx, "/content.ContentService/TagGroupList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) AttrInfo(ctx context.Context, in *request.Object, opts ...grpc.CallOption) (*Attributes, error) {
	out := new(Attributes)
	err := c.cc.Invoke(ctx, "/content.ContentService/AttrInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) AddAttr(ctx context.Context, in *AddAttrReq, opts ...grpc.CallOption) (*request.Object, error) {
	out := new(request.Object)
	err := c.cc.Invoke(ctx, "/content.ContentService/AddAttr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) EditAttr(ctx context.Context, in *EditAttrReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.ContentService/EditAttr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) AttrList(ctx context.Context, in *AttrListReq, opts ...grpc.CallOption) (*AttrListRep, error) {
	out := new(AttrListRep)
	err := c.cc.Invoke(ctx, "/content.ContentService/AttrList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) FavList(ctx context.Context, in *FavListReq, opts ...grpc.CallOption) (*FavListRep, error) {
	out := new(FavListRep)
	err := c.cc.Invoke(ctx, "/content.ContentService/FavList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) TinyFavList(ctx context.Context, in *FavListReq, opts ...grpc.CallOption) (*TinyFavListRep, error) {
	out := new(TinyFavListRep)
	err := c.cc.Invoke(ctx, "/content.ContentService/TinyFavList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) AddFav(ctx context.Context, in *AddFavReq, opts ...grpc.CallOption) (*request.Object, error) {
	out := new(request.Object)
	err := c.cc.Invoke(ctx, "/content.ContentService/AddFav", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) EditFav(ctx context.Context, in *AddFavReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.ContentService/EditFav", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) AddContainer(ctx context.Context, in *AddContainerReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.ContentService/AddContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) EditDiaryContainer(ctx context.Context, in *AddContainerReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.ContentService/EditDiaryContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) UserContentCount(ctx context.Context, in *request.Object, opts ...grpc.CallOption) (*user.UserContent, error) {
	out := new(user.UserContent)
	err := c.cc.Invoke(ctx, "/content.ContentService/UserContentCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServiceServer is the server API for ContentService service.
// All implementations must embed UnimplementedContentServiceServer
// for forward compatibility
type ContentServiceServer interface {
	// 详情
	TagInfo(context.Context, *request.Object) (*Tag, error)
	// 新建
	AddTag(context.Context, *AddTagReq) (*empty.Empty, error)
	// 修改
	EditTag(context.Context, *EditTagReq) (*empty.Empty, error)
	// 列表
	TagList(context.Context, *TagListReq) (*TagListRep, error)
	// 列表
	TagGroupList(context.Context, *TagGroupListReq) (*TagGroupListRep, error)
	// 详情
	AttrInfo(context.Context, *request.Object) (*Attributes, error)
	// 新建
	AddAttr(context.Context, *AddAttrReq) (*request.Object, error)
	// 修改
	EditAttr(context.Context, *EditAttrReq) (*empty.Empty, error)
	// 列表
	AttrList(context.Context, *AttrListReq) (*AttrListRep, error)
	//收藏夹列表
	FavList(context.Context, *FavListReq) (*FavListRep, error)
	//收藏夹列表
	TinyFavList(context.Context, *FavListReq) (*TinyFavListRep, error)
	// 创建收藏夹
	AddFav(context.Context, *AddFavReq) (*request.Object, error)
	// 修改收藏夹
	EditFav(context.Context, *AddFavReq) (*empty.Empty, error)
	// 创建合集
	AddContainer(context.Context, *AddContainerReq) (*empty.Empty, error)
	// 修改合集
	EditDiaryContainer(context.Context, *AddContainerReq) (*empty.Empty, error)
	// 用户内容数量
	UserContentCount(context.Context, *request.Object) (*user.UserContent, error)
	mustEmbedUnimplementedContentServiceServer()
}

// UnimplementedContentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContentServiceServer struct {
}

func (UnimplementedContentServiceServer) TagInfo(context.Context, *request.Object) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TagInfo not implemented")
}
func (UnimplementedContentServiceServer) AddTag(context.Context, *AddTagReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTag not implemented")
}
func (UnimplementedContentServiceServer) EditTag(context.Context, *EditTagReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditTag not implemented")
}
func (UnimplementedContentServiceServer) TagList(context.Context, *TagListReq) (*TagListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TagList not implemented")
}
func (UnimplementedContentServiceServer) TagGroupList(context.Context, *TagGroupListReq) (*TagGroupListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TagGroupList not implemented")
}
func (UnimplementedContentServiceServer) AttrInfo(context.Context, *request.Object) (*Attributes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttrInfo not implemented")
}
func (UnimplementedContentServiceServer) AddAttr(context.Context, *AddAttrReq) (*request.Object, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAttr not implemented")
}
func (UnimplementedContentServiceServer) EditAttr(context.Context, *EditAttrReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditAttr not implemented")
}
func (UnimplementedContentServiceServer) AttrList(context.Context, *AttrListReq) (*AttrListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttrList not implemented")
}
func (UnimplementedContentServiceServer) FavList(context.Context, *FavListReq) (*FavListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavList not implemented")
}
func (UnimplementedContentServiceServer) TinyFavList(context.Context, *FavListReq) (*TinyFavListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TinyFavList not implemented")
}
func (UnimplementedContentServiceServer) AddFav(context.Context, *AddFavReq) (*request.Object, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFav not implemented")
}
func (UnimplementedContentServiceServer) EditFav(context.Context, *AddFavReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditFav not implemented")
}
func (UnimplementedContentServiceServer) AddContainer(context.Context, *AddContainerReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddContainer not implemented")
}
func (UnimplementedContentServiceServer) EditDiaryContainer(context.Context, *AddContainerReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditDiaryContainer not implemented")
}
func (UnimplementedContentServiceServer) UserContentCount(context.Context, *request.Object) (*user.UserContent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserContentCount not implemented")
}
func (UnimplementedContentServiceServer) mustEmbedUnimplementedContentServiceServer() {}

// UnsafeContentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServiceServer will
// result in compilation errors.
type UnsafeContentServiceServer interface {
	mustEmbedUnimplementedContentServiceServer()
}

func RegisterContentServiceServer(s grpc.ServiceRegistrar, srv ContentServiceServer) {
	s.RegisterService(&ContentService_ServiceDesc, srv)
}

func _ContentService_TagInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).TagInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/TagInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).TagInfo(ctx, req.(*request.Object))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_AddTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).AddTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/AddTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).AddTag(ctx, req.(*AddTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_EditTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).EditTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/EditTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).EditTag(ctx, req.(*EditTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_TagList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).TagList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/TagList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).TagList(ctx, req.(*TagListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_TagGroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagGroupListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).TagGroupList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/TagGroupList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).TagGroupList(ctx, req.(*TagGroupListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_AttrInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).AttrInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/AttrInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).AttrInfo(ctx, req.(*request.Object))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_AddAttr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAttrReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).AddAttr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/AddAttr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).AddAttr(ctx, req.(*AddAttrReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_EditAttr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditAttrReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).EditAttr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/EditAttr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).EditAttr(ctx, req.(*EditAttrReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_AttrList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttrListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).AttrList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/AttrList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).AttrList(ctx, req.(*AttrListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_FavList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).FavList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/FavList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).FavList(ctx, req.(*FavListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_TinyFavList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).TinyFavList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/TinyFavList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).TinyFavList(ctx, req.(*FavListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_AddFav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFavReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).AddFav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/AddFav",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).AddFav(ctx, req.(*AddFavReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_EditFav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFavReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).EditFav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/EditFav",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).EditFav(ctx, req.(*AddFavReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_AddContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddContainerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).AddContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/AddContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).AddContainer(ctx, req.(*AddContainerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_EditDiaryContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddContainerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).EditDiaryContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/EditDiaryContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).EditDiaryContainer(ctx, req.(*AddContainerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_UserContentCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).UserContentCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/UserContentCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).UserContentCount(ctx, req.(*request.Object))
	}
	return interceptor(ctx, in, info, handler)
}

// ContentService_ServiceDesc is the grpc.ServiceDesc for ContentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content.ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TagInfo",
			Handler:    _ContentService_TagInfo_Handler,
		},
		{
			MethodName: "AddTag",
			Handler:    _ContentService_AddTag_Handler,
		},
		{
			MethodName: "EditTag",
			Handler:    _ContentService_EditTag_Handler,
		},
		{
			MethodName: "TagList",
			Handler:    _ContentService_TagList_Handler,
		},
		{
			MethodName: "TagGroupList",
			Handler:    _ContentService_TagGroupList_Handler,
		},
		{
			MethodName: "AttrInfo",
			Handler:    _ContentService_AttrInfo_Handler,
		},
		{
			MethodName: "AddAttr",
			Handler:    _ContentService_AddAttr_Handler,
		},
		{
			MethodName: "EditAttr",
			Handler:    _ContentService_EditAttr_Handler,
		},
		{
			MethodName: "AttrList",
			Handler:    _ContentService_AttrList_Handler,
		},
		{
			MethodName: "FavList",
			Handler:    _ContentService_FavList_Handler,
		},
		{
			MethodName: "TinyFavList",
			Handler:    _ContentService_TinyFavList_Handler,
		},
		{
			MethodName: "AddFav",
			Handler:    _ContentService_AddFav_Handler,
		},
		{
			MethodName: "EditFav",
			Handler:    _ContentService_EditFav_Handler,
		},
		{
			MethodName: "AddContainer",
			Handler:    _ContentService_AddContainer_Handler,
		},
		{
			MethodName: "EditDiaryContainer",
			Handler:    _ContentService_EditDiaryContainer_Handler,
		},
		{
			MethodName: "UserContentCount",
			Handler:    _ContentService_UserContentCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content/content.service.proto",
}