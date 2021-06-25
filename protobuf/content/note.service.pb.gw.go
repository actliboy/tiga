// Code generated by protoc-gen-grpc-gin. DO NOT EDIT.
// source: content/note.service.proto

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

func request_NoteService_Create_0(ctx *gin.Context, client NoteServiceClient) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Note
	var metadata runtime.ServerMetadata

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, metadata, err
	}

	msg, err := client.Create(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_NoteService_Create_0(server NoteServiceServer, ctx *gin.Context) (proto.Message, error) {
	var protoReq Note

	if err := gin_0.Bind(ctx, &protoReq); err != nil {
		return nil, err
	}

	return server.Create(ctx.Request.Context(), &protoReq)

}

// RegisterNoteServiceHandlerServer registers the http handlers for service NoteService to "mux".
// UnaryRPC     :call NoteServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterNoteServiceHandlerFromEndpoint instead.
func RegisterNoteServiceHandlerServer(mux *gin.Engine, server NoteServiceServer) error {

	mux.Handle("POST", "/api/v1/note", func(ctx *gin.Context) {
		var md runtime.ServerMetadata
		resp, err := local_request_NoteService_Create_0(server, ctx)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	return nil
}

// RegisterNoteServiceHandlerFromEndpoint is same as RegisterNoteServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterNoteServiceHandlerFromEndpoint(ctx context.Context, mux *gin.Engine, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterNoteServiceHandler(ctx, mux, conn)
}

// RegisterNoteServiceHandler registers the http handlers for service NoteService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterNoteServiceHandler(ctx context.Context, mux *gin.Engine, conn *grpc.ClientConn) error {
	return RegisterNoteServiceHandlerClient(ctx, mux, NewNoteServiceClient(conn))
}

// RegisterNoteServiceHandlerClient registers the http handlers for service NoteService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "NoteServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "NoteServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "NoteServiceClient" to call the correct interceptors.
func RegisterNoteServiceHandlerClient(ctx context.Context, mux *gin.Engine, client NoteServiceClient) error {

	mux.Handle("POST", "/api/v1/note", func(ctx *gin.Context) {
		resp, md, err := request_NoteService_Create_0(ctx, client)
		if err != nil {
			gateway.HTTPError(ctx, err)
			return
		}

		gateway.ForwardResponseMessage(ctx, md, resp)

	})

	return nil
}