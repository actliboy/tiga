// Code generated by protoc-gen-grpc-gin. DO NOT EDIT.
// source: content/content.service.proto

/*
Package content is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package content

import (
	"context"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	request_0 "github.com/liov/tiga/protobuf/utils/request"
	gin_0 "github.com/liov/tiga/utils/net/http/gin"
	"github.com/liov/tiga/utils/net/http/grpc/gateway"
	"github.com/liov/tiga/utils/net/http/request"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join
var _ = request.Error

func request_ContentService_TagInfo_0(ctx *gin.Context, client ContentServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq request_0.Object
	var metadata runtime.ServerMetadata

	var (
		err error
		_   = err
	)

	protoReq.Id, err = runtime.Uint64(ctx.Param("id"))
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", request.Error(err))
	}

	msg, err := client.TagInfo(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ContentService_TagInfo_0(server ContentServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq request_0.Object

	var (
		err error
		_   = err
	)

	protoReq.Id, err = runtime.Uint64(ctx.Param("id"))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", request.Error(err))
	}

	return server.TagInfo(ctx.Request.Context(), &protoReq)

}

func request_ContentService_AddTag_0(ctx *gin.Context, client ContentServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddTagReq
	var metadata runtime.ServerMetadata

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, metadata, err
	}

	msg, err := client.AddTag(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ContentService_AddTag_0(server ContentServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq AddTagReq

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, err
	}

	return server.AddTag(ctx.Request.Context(), &protoReq)

}

func request_ContentService_EditTag_0(ctx *gin.Context, client ContentServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq EditTagReq
	var metadata runtime.ServerMetadata

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, metadata, err
	}

	var (
		err error
		_   = err
	)

	protoReq.Id, err = runtime.Uint64(ctx.Param("id"))
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", request.Error(err))
	}

	msg, err := client.EditTag(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ContentService_EditTag_0(server ContentServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq EditTagReq

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, err
	}

	var (
		err error
		_   = err
	)

	protoReq.Id, err = runtime.Uint64(ctx.Param("id"))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", request.Error(err))
	}

	return server.EditTag(ctx.Request.Context(), &protoReq)

}

func request_ContentService_TagList_0(ctx *gin.Context, client ContentServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TagListReq
	var metadata runtime.ServerMetadata

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, metadata, err
	}

	msg, err := client.TagList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ContentService_TagList_0(server ContentServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq TagListReq

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, err
	}

	return server.TagList(ctx.Request.Context(), &protoReq)

}

func request_ContentService_TagGroupList_0(ctx *gin.Context, client ContentServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TagGroupListReq
	var metadata runtime.ServerMetadata

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, metadata, err
	}

	msg, err := client.TagGroupList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ContentService_TagGroupList_0(server ContentServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq TagGroupListReq

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, err
	}

	return server.TagGroupList(ctx.Request.Context(), &protoReq)

}

func request_ContentService_AttrInfo_0(ctx *gin.Context, client ContentServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq request_0.Object
	var metadata runtime.ServerMetadata

	var (
		err error
		_   = err
	)

	protoReq.Id, err = runtime.Uint64(ctx.Param("id"))
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", request.Error(err))
	}

	msg, err := client.AttrInfo(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ContentService_AttrInfo_0(server ContentServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq request_0.Object

	var (
		err error
		_   = err
	)

	protoReq.Id, err = runtime.Uint64(ctx.Param("id"))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", request.Error(err))
	}

	return server.AttrInfo(ctx.Request.Context(), &protoReq)

}

func request_ContentService_AddAttr_0(ctx *gin.Context, client ContentServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddAttrReq
	var metadata runtime.ServerMetadata

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, metadata, err
	}

	msg, err := client.AddAttr(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ContentService_AddAttr_0(server ContentServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq AddAttrReq

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, err
	}

	return server.AddAttr(ctx.Request.Context(), &protoReq)

}

func request_ContentService_EditAttr_0(ctx *gin.Context, client ContentServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq EditAttrReq
	var metadata runtime.ServerMetadata

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, metadata, err
	}

	var (
		err error
		_   = err
	)

	protoReq.Id, err = runtime.Uint64(ctx.Param("id"))
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", request.Error(err))
	}

	msg, err := client.EditAttr(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ContentService_EditAttr_0(server ContentServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq EditAttrReq

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, err
	}

	var (
		err error
		_   = err
	)

	protoReq.Id, err = runtime.Uint64(ctx.Param("id"))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", request.Error(err))
	}

	return server.EditAttr(ctx.Request.Context(), &protoReq)

}

func request_ContentService_AttrList_0(ctx *gin.Context, client ContentServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AttrListReq
	var metadata runtime.ServerMetadata

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, metadata, err
	}

	msg, err := client.AttrList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ContentService_AttrList_0(server ContentServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq AttrListReq

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, err
	}

	return server.AttrList(ctx.Request.Context(), &protoReq)

}

func request_ContentService_FavList_0(ctx *gin.Context, client ContentServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FavListReq
	var metadata runtime.ServerMetadata

	var (
		err error
		_   = err
	)

	protoReq.UserId, err = runtime.Uint32(ctx.Param("userId"))
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "userId", request.Error(err))
	}

	msg, err := client.FavList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ContentService_FavList_0(server ContentServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq FavListReq

	var (
		err error
		_   = err
	)

	protoReq.UserId, err = runtime.Uint32(ctx.Param("userId"))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "userId", request.Error(err))
	}

	return server.FavList(ctx.Request.Context(), &protoReq)

}

func request_ContentService_TinyFavList_0(ctx *gin.Context, client ContentServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FavListReq
	var metadata runtime.ServerMetadata

	var (
		err error
		_   = err
	)

	protoReq.UserId, err = runtime.Uint32(ctx.Param("userId"))
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "userId", request.Error(err))
	}

	msg, err := client.TinyFavList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ContentService_TinyFavList_0(server ContentServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq FavListReq

	var (
		err error
		_   = err
	)

	protoReq.UserId, err = runtime.Uint32(ctx.Param("userId"))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "userId", request.Error(err))
	}

	return server.TinyFavList(ctx.Request.Context(), &protoReq)

}

func request_ContentService_AddFav_0(ctx *gin.Context, client ContentServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddFavReq
	var metadata runtime.ServerMetadata

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, metadata, err
	}

	msg, err := client.AddFav(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ContentService_AddFav_0(server ContentServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq AddFavReq

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, err
	}

	return server.AddFav(ctx.Request.Context(), &protoReq)

}

func request_ContentService_EditFav_0(ctx *gin.Context, client ContentServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddFavReq
	var metadata runtime.ServerMetadata

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, metadata, err
	}

	var (
		err error
		_   = err
	)

	protoReq.Id, err = runtime.Uint64(ctx.Param("id"))
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", request.Error(err))
	}

	msg, err := client.EditFav(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ContentService_EditFav_0(server ContentServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq AddFavReq

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, err
	}

	var (
		err error
		_   = err
	)

	protoReq.Id, err = runtime.Uint64(ctx.Param("id"))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", request.Error(err))
	}

	return server.EditFav(ctx.Request.Context(), &protoReq)

}

func request_ContentService_AddContainer_0(ctx *gin.Context, client ContentServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddContainerReq
	var metadata runtime.ServerMetadata

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, metadata, err
	}

	msg, err := client.AddContainer(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ContentService_AddContainer_0(server ContentServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq AddContainerReq

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, err
	}

	return server.AddContainer(ctx.Request.Context(), &protoReq)

}

func request_ContentService_EditDiaryContainer_0(ctx *gin.Context, client ContentServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddContainerReq
	var metadata runtime.ServerMetadata

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, metadata, err
	}

	var (
		err error
		_   = err
	)

	protoReq.Id, err = runtime.Uint64(ctx.Param("id"))
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", request.Error(err))
	}

	msg, err := client.EditDiaryContainer(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ContentService_EditDiaryContainer_0(server ContentServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq AddContainerReq

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, err
	}

	var (
		err error
		_   = err
	)

	protoReq.Id, err = runtime.Uint64(ctx.Param("id"))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", request.Error(err))
	}

	return server.EditDiaryContainer(ctx.Request.Context(), &protoReq)

}

func request_ContentService_UserContentCount_0(ctx *gin.Context, client ContentServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq request_0.Object
	var metadata runtime.ServerMetadata

	var (
		err error
		_   = err
	)

	protoReq.Id, err = runtime.Uint64(ctx.Param("id"))
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", request.Error(err))
	}

	msg, err := client.UserContentCount(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ContentService_UserContentCount_0(server ContentServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq request_0.Object

	var (
		err error
		_   = err
	)

	protoReq.Id, err = runtime.Uint64(ctx.Param("id"))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", request.Error(err))
	}

	return server.UserContentCount(ctx.Request.Context(), &protoReq)

}

// RegisterContentServiceHandlerServer registers the http handlers for service ContentService to "mux".
// UnaryRPC     :call ContentServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterContentServiceHandlerFromEndpoint instead.
func RegisterContentServiceHandlerServer(mux *gin.Engine, server ContentServiceServer) error {

	mux.Handle("GET", "/api/v1/content/tag/:id", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_ContentService_TagInfo_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("POST", "/api/v1/content/tag", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_ContentService_AddTag_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("PUT", "/api/v1/content/tag/:id", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_ContentService_EditTag_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/content/tag", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_ContentService_TagList_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/content/tagGroup", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_ContentService_TagGroupList_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/content/attr/:id", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_ContentService_AttrInfo_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("POST", "/api/v1/content/attr", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_ContentService_AddAttr_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("PUT", "/api/v1/content/attr/:id", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_ContentService_EditAttr_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/content/attr", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_ContentService_AttrList_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/content/fav/:userId", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_ContentService_FavList_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/content/tinyFav/:userId", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_ContentService_TinyFavList_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("POST", "/api/v1/content/fav", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_ContentService_AddFav_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("PUT", "/api/v1/content/fav/:id", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_ContentService_EditFav_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("POST", "/api/v1/content/container", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_ContentService_AddContainer_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("PUT", "/api/v1/content/container/:id", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_ContentService_EditDiaryContainer_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/content/count/:id", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_ContentService_UserContentCount_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	return nil
}

// RegisterContentServiceHandlerFromEndpoint is same as RegisterContentServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterContentServiceHandlerFromEndpoint(ctx context.Context, mux *gin.Engine, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterContentServiceHandler(ctx, mux, conn)
}

// RegisterContentServiceHandler registers the http handlers for service ContentService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterContentServiceHandler(ctx context.Context, mux *gin.Engine, conn *grpc.ClientConn) error {
	return RegisterContentServiceHandlerClient(ctx, mux, NewContentServiceClient(conn))
}

// RegisterContentServiceHandlerClient registers the http handlers for service ContentService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ContentServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ContentServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ContentServiceClient" to call the correct interceptors.
func RegisterContentServiceHandlerClient(ctx context.Context, mux *gin.Engine, client ContentServiceClient) error {

	mux.Handle("GET", "/api/v1/content/tag/:id", func(ctx *gin.Context) {
		resp, md, err := request_ContentService_TagInfo_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("POST", "/api/v1/content/tag", func(ctx *gin.Context) {
		resp, md, err := request_ContentService_AddTag_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("PUT", "/api/v1/content/tag/:id", func(ctx *gin.Context) {
		resp, md, err := request_ContentService_EditTag_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/content/tag", func(ctx *gin.Context) {
		resp, md, err := request_ContentService_TagList_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/content/tagGroup", func(ctx *gin.Context) {
		resp, md, err := request_ContentService_TagGroupList_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/content/attr/:id", func(ctx *gin.Context) {
		resp, md, err := request_ContentService_AttrInfo_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("POST", "/api/v1/content/attr", func(ctx *gin.Context) {
		resp, md, err := request_ContentService_AddAttr_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("PUT", "/api/v1/content/attr/:id", func(ctx *gin.Context) {
		resp, md, err := request_ContentService_EditAttr_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/content/attr", func(ctx *gin.Context) {
		resp, md, err := request_ContentService_AttrList_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/content/fav/:userId", func(ctx *gin.Context) {
		resp, md, err := request_ContentService_FavList_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/content/tinyFav/:userId", func(ctx *gin.Context) {
		resp, md, err := request_ContentService_TinyFavList_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("POST", "/api/v1/content/fav", func(ctx *gin.Context) {
		resp, md, err := request_ContentService_AddFav_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("PUT", "/api/v1/content/fav/:id", func(ctx *gin.Context) {
		resp, md, err := request_ContentService_EditFav_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("POST", "/api/v1/content/container", func(ctx *gin.Context) {
		resp, md, err := request_ContentService_AddContainer_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("PUT", "/api/v1/content/container/:id", func(ctx *gin.Context) {
		resp, md, err := request_ContentService_EditDiaryContainer_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/content/count/:id", func(ctx *gin.Context) {
		resp, md, err := request_ContentService_UserContentCount_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	return nil
}
