package loginService

import (
	"golang.org/x/net/context"
	pb "grpc-game-server/pkg/api/proto"
)

type LoginServer struct {
	pb.UnimplementedLoginServiceServer
}

func (s *LoginServer) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {

	//platformId := req.PlatformId
	//platformToken := req.PlatformToken

	return &pb.LoginResponse{
		Status: pb.LoginResponse_SUCCESS,
		User:   nil,
	}, nil
}
