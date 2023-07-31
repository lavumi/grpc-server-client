package cmd

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	pb "grpc-game-server/pkg/api/proto"
	"grpc-game-server/pkg/database"
	"grpc-game-server/pkg/service/chatService"
	"log"
	"net"
)

func Start() {
	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}

	db := database.Database{}
	db.Init()

	server := &chatService.ChattingServer{}
	server.Init()

	grpcServer := grpc.NewServer()

	pb.RegisterRoomServiceServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("fail to serve: %v", err)
	}
}
