// Code generated by protoc-gen-grpc-gin. DO NOT EDIT.
// source: content/diary.service.proto

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

func request_DiaryService_DiaryBook_0(ctx *gin.Context, client DiaryServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DiaryBookReq
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

	msg, err := client.DiaryBook(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DiaryService_DiaryBook_0(server DiaryServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq DiaryBookReq

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

	return server.DiaryBook(ctx.Request.Context(), &protoReq)

}

func request_DiaryService_DiaryBookList_0(ctx *gin.Context, client DiaryServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DiaryBookListReq
	var metadata runtime.ServerMetadata

	msg, err := client.DiaryBookList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DiaryService_DiaryBookList_0(server DiaryServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq DiaryBookListReq

	return server.DiaryBookList(ctx.Request.Context(), &protoReq)

}

func request_DiaryService_AddDiaryBook_0(ctx *gin.Context, client DiaryServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddDiaryBookReq
	var metadata runtime.ServerMetadata

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, metadata, err
	}

	msg, err := client.AddDiaryBook(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DiaryService_AddDiaryBook_0(server DiaryServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq AddDiaryBookReq

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, err
	}

	return server.AddDiaryBook(ctx.Request.Context(), &protoReq)

}

func request_DiaryService_EditDiaryBook_0(ctx *gin.Context, client DiaryServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddDiaryBookReq
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

	msg, err := client.EditDiaryBook(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DiaryService_EditDiaryBook_0(server DiaryServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq AddDiaryBookReq

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

	return server.EditDiaryBook(ctx.Request.Context(), &protoReq)

}

func request_DiaryService_Info_0(ctx *gin.Context, client DiaryServiceClient) (proto.Message, runtime.ServerMetadata, error) {
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

	msg, err := client.Info(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DiaryService_Info_0(server DiaryServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq request_0.Object

	var (
		err error
		_   = err
	)

	protoReq.Id, err = runtime.Uint64(ctx.Param("id"))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", request.Error(err))
	}

	return server.Info(ctx.Request.Context(), &protoReq)

}

func request_DiaryService_Add_0(ctx *gin.Context, client DiaryServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddDiaryReq
	var metadata runtime.ServerMetadata

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, metadata, err
	}

	msg, err := client.Add(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DiaryService_Add_0(server DiaryServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq AddDiaryReq

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, err
	}

	return server.Add(ctx.Request.Context(), &protoReq)

}

func request_DiaryService_Edit_0(ctx *gin.Context, client DiaryServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddDiaryReq
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

	msg, err := client.Edit(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DiaryService_Edit_0(server DiaryServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq AddDiaryReq

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

	return server.Edit(ctx.Request.Context(), &protoReq)

}

func request_DiaryService_List_0(ctx *gin.Context, client DiaryServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DiaryListReq
	var metadata runtime.ServerMetadata

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, metadata, err
	}

	msg, err := client.List(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DiaryService_List_0(server DiaryServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq DiaryListReq

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, err
	}

	return server.List(ctx.Request.Context(), &protoReq)

}

func request_DiaryService_Delete_0(ctx *gin.Context, client DiaryServiceClient) (proto.Message, runtime.ServerMetadata, error) {
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

	msg, err := client.Delete(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DiaryService_Delete_0(server DiaryServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq request_0.Object

	var (
		err error
		_   = err
	)

	protoReq.Id, err = runtime.Uint64(ctx.Param("id"))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", request.Error(err))
	}

	return server.Delete(ctx.Request.Context(), &protoReq)

}

// RegisterDiaryServiceHandlerServer registers the http handlers for service DiaryService to "mux".
// UnaryRPC     :call DiaryServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterDiaryServiceHandlerFromEndpoint instead.
func RegisterDiaryServiceHandlerServer(mux *gin.Engine, server DiaryServiceServer) error {

	mux.Handle("GET", "/api/v1/diaryBook/:id", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_DiaryService_DiaryBook_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/diaryBook", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_DiaryService_DiaryBookList_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("POST", "/api/v1/diaryBook", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_DiaryService_AddDiaryBook_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("PUT", "/api/v1/diaryBook/:id", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_DiaryService_EditDiaryBook_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/diary/:id", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_DiaryService_Info_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("POST", "/api/v1/diary", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_DiaryService_Add_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("PUT", "/api/v1/diary/:id", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_DiaryService_Edit_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/diary", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_DiaryService_List_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("DELETE", "/api/v1/diary/:id", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_DiaryService_Delete_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	return nil
}

// RegisterDiaryServiceHandlerFromEndpoint is same as RegisterDiaryServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterDiaryServiceHandlerFromEndpoint(ctx context.Context, mux *gin.Engine, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterDiaryServiceHandler(ctx, mux, conn)
}

// RegisterDiaryServiceHandler registers the http handlers for service DiaryService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterDiaryServiceHandler(ctx context.Context, mux *gin.Engine, conn *grpc.ClientConn) error {
	return RegisterDiaryServiceHandlerClient(ctx, mux, NewDiaryServiceClient(conn))
}

// RegisterDiaryServiceHandlerClient registers the http handlers for service DiaryService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "DiaryServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "DiaryServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "DiaryServiceClient" to call the correct interceptors.
func RegisterDiaryServiceHandlerClient(ctx context.Context, mux *gin.Engine, client DiaryServiceClient) error {

	mux.Handle("GET", "/api/v1/diaryBook/:id", func(ctx *gin.Context) {
		resp, md, err := request_DiaryService_DiaryBook_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/diaryBook", func(ctx *gin.Context) {
		resp, md, err := request_DiaryService_DiaryBookList_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("POST", "/api/v1/diaryBook", func(ctx *gin.Context) {
		resp, md, err := request_DiaryService_AddDiaryBook_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("PUT", "/api/v1/diaryBook/:id", func(ctx *gin.Context) {
		resp, md, err := request_DiaryService_EditDiaryBook_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/diary/:id", func(ctx *gin.Context) {
		resp, md, err := request_DiaryService_Info_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("POST", "/api/v1/diary", func(ctx *gin.Context) {
		resp, md, err := request_DiaryService_Add_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("PUT", "/api/v1/diary/:id", func(ctx *gin.Context) {
		resp, md, err := request_DiaryService_Edit_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("GET", "/api/v1/diary", func(ctx *gin.Context) {
		resp, md, err := request_DiaryService_List_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	mux.Handle("DELETE", "/api/v1/diary/:id", func(ctx *gin.Context) {
		resp, md, err := request_DiaryService_Delete_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	return nil
}