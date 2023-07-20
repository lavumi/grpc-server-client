package chat_service

import (
	"crypto/sha256"
	"encoding/hex"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	pb "grpc-game-server/pkg/api/proto"
	"time"
)

type RoomManager struct {
	RoomStreamMap map[string][]*Connection
	Rooms         map[string]*pb.Room
}

func (m *RoomManager) Init() {
	m.RoomStreamMap = make(map[string][]*Connection)
	m.Rooms = make(map[string]*pb.Room)
}

func (m *RoomManager) CreateNewRoom(roomName string) (string, error) {
	if m.Rooms[roomName] != nil {
		return "", status.Errorf(codes.AlreadyExists, "room AlreadyExist")
	}

	id := sha256.Sum256([]byte(roomName + time.Now().String()))
	stringId := hex.EncodeToString(id[:])
	m.Rooms[stringId] = &pb.Room{
		RoomId:   stringId,
		Name:     roomName,
		Members:  nil,
		Messages: nil,
	}

	m.RoomStreamMap[stringId] = nil

	grpclog.Infof("%s room created", stringId)
	return stringId, nil
}

func (m *RoomManager) JoinRoom(roomId string, user *pb.User, conn *Connection) error {
	if m.Rooms[roomId] == nil {
		return status.Errorf(codes.AlreadyExists, "roomId %s No tExist", roomId)
	}
	//if m.RoomStreamMap[roomId] == nil {
	//	return status.Errorf(codes.AlreadyExists, "stream %s  Not Exist", roomId)
	//}

	m.Rooms[roomId].Members = append(m.Rooms[roomId].Members, user)
	m.RoomStreamMap[roomId] = append(m.RoomStreamMap[roomId], conn)

	grpclog.Infof("Join Room Success")
	grpclog.Infof("UserId %s", user.UserId)
	grpclog.Infof("RoomId %s", roomId)
	grpclog.Infof("Rooms Member Len %d", len(m.RoomStreamMap[roomId]))
	//grpclog.Infof("RoomId %s", roomId)
	grpclog.Infof("---------------------------------------------")
	return nil
}

func (m *RoomManager) GetRoomInfo(roomId string) (*pb.Room, error) {
	if m.Rooms[roomId] == nil {
		return nil, status.Errorf(codes.NotFound, "room No tExist")
	}
	//if m.RoomStreamMap[roomId] == nil {
	//	return nil, status.Errorf(codes.NotFound, "stream Not Exist")
	//}

	return m.Rooms[roomId], nil
}

func (m *RoomManager) GetStream(roomId string) ([]*Connection, error) {
	if m.Rooms[roomId] == nil {
		return nil, status.Errorf(codes.AlreadyExists, "room No tExist")
	}
	//if m.RoomStreamMap[roomId] == nil {
	//	return nil, status.Errorf(codes.AlreadyExists, "stream Not Exist")
	//}

	return m.RoomStreamMap[roomId], nil

}
