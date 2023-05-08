package main

import (
	"time"
	"log"

	pb "github.com/KeshikaGupta20/GRPC_Stream_Demo/proto"
)

func (s *helloServer)  SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {

	log.Printf("Got request with names: ", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello" + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}

	return nil
}
