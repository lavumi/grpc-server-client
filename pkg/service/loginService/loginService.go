package loginService

import (
	"golang.org/x/net/context"
	pb "grpc-game-server/pkg/api/proto"
	"grpc-game-server/pkg/database"
)

type LoginServer struct {
	pb.UnimplementedLoginServiceServer
	db *database.MongoDb
}

func (s *LoginServer) Init(db *database.MongoDb) {
	s.db = db
}

func (s *LoginServer) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {

	//platformId := req.PlatformId
	//platformToken := req.PlatformToken

	return &pb.LoginResponse{
		Status: pb.LoginResponse_SUCCESS,
		User:   nil,
	}, nil
}
