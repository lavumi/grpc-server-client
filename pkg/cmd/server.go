package cmd

import (
	"google.golang.org/grpc"
	pb "grpc-game-server/pkg/api/proto"
	"grpc-game-server/pkg/service"
	"log"
	"net"
)

func Start() {
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	server := &service.ChattingServer{}
	server.Init()

	server.Rooms["test"] = &service.ChatRoom{
		Connection: nil,
		Id:         "test",
		Name:       "test_room",
		Members:    nil,
		Messages:   nil,
	}

	pb.RegisterChatServiceServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("fail to serve: %v", err)
	}
}
