package main

import (
	pb "github.com/KeshikaGupta20/GRPC_Stream_Demo/proto"

	"context"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello there !!!",
	}, nil
}
