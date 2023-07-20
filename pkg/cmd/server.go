package cmd

import (
	"google.golang.org/grpc"
	pb "grpc-game-server/pkg/api/proto"
	"grpc-game-server/pkg/service/chat-service"
	"log"
	"net"
)

func Start() {
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	server := &chat_service.ChattingServer{}
	server.Init()

	//server.WaitingRooms["test_room_1"] = &chat_service.Room{
	//	Connection: nil,
	//	Id:         "test1",
	//	Name:       "test_room_1",
	//	Members:    nil,
	//	Messages:   nil,
	//}
	//
	//server.WaitingRooms["test_room_2"] = &chat_service.Room{
	//	Connection: nil,
	//	Id:         "test2",
	//	Name:       "test_room_2",
	//	Members:    nil,
	//	Messages:   nil,
	//}

	pb.RegisterRoomServiceServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("fail to serve: %v", err)
	}
}
