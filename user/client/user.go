package client

import (
	model "github.com/liov/tiga/protobuf/user"
	"github.com/liov/tiga/utils/log"
	"google.golang.org/grpc"
)

func GetUserClient() (model.UserServiceClient, *grpc.ClientConn) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return model.NewUserServiceClient(conn), conn
}
