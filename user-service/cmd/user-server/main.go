package main

import (
	"fmt"
	api "github.com/matf-pp/2022_MATDAQ/api/user-service"
	"github.com/matf-pp/2022_MATDAQ/user-service/internal"
	"google.golang.org/grpc"
	"log"
	"net"
)

const PORT int = 9000

func main() {

	internal.InitRedis()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterUserServer(grpcServer, internal.NewUserServer())
	grpcServer.Serve(lis)
}
