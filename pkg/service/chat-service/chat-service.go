package chat_service

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	pb "grpc-game-server/pkg/api/proto"
	"log"
	"sync"
)

type Connection struct {
	stream pb.RoomService_CreateStreamServer
	id     string
	active bool
	error  chan error
}

var roomManager RoomManager

type ChattingServer struct {
	//pb.mustEmbedUnimplementedRoomServiceServer
	pb.UnimplementedRoomServiceServer
}

func (s *ChattingServer) Init() {
	roomManager.Init()
}

func (s *ChattingServer) CreateStream(request *pb.RoomRequest, stream pb.RoomService_CreateStreamServer) error {

	if request.Room == nil {
		return status.Errorf(codes.InvalidArgument, "room id needed")
	}
	roomId := request.Room.RoomId

	conn := &Connection{
		stream: stream,
		id:     request.User.UserId,
		active: true,
		error:  make(chan error),
	}

	err := roomManager.JoinRoom(roomId, request.User, conn)
	if err != nil {
		return err
	}

	//streamList, err2 := roomManager.GetStream(roomId)
	//if err2 != nil {
	//	return err2
	//}

	//room := s.WaitingRooms[roomName]
	//if room == nil {
	//	return status.Errorf(codes.NotFound, "room Not Found")
	//}
	//
	//conn := &Connection{
	//	stream: stream,
	//	id:     request.User.UserId,
	//	active: true,
	//	error:  make(chan error),
	//}
	//room.Connection = append(room.Connection, conn)
	//
	//s.UserRoomHash[request.User.Id] = room.Name
	//
	//log.Printf("Connection Length : %d", len(room.Connection))

	//systemUser := pb.User{
	//	Id:          "1",
	//	DisplayName: "System",
	//}
	//
	//timestamp := time.Now()
	//joinMessage := pb.Message{
	//	Id:        systemUser.Id,
	//	User:      &systemUser,
	//	Message:   fmt.Sprintf("%v has entered the room", request.User.DisplayName),
	//	Timestamp: timestamp.String(),
	//}
	//
	//_, err := s.BroadcastMessage(context.Background(), &joinMessage)
	//if err != nil {
	//	return err
	//}

	return <-conn.error
}

func (s *ChattingServer) BroadcastMessage(ctx context.Context, msg *pb.Message) (*pb.Close, error) {
	wait := sync.WaitGroup{}
	done := make(chan int)
	disconnected := false

	//md, ok := metadata.FromIncomingContext(ctx)
	//if ok {
	//	authorizationHeader := md.Get("authorization")
	//	log.Print(authorizationHeader)
	//}

	//log.Printf(msg.User.UserId)

	connections, err := roomManager.GetStream(msg.RoomId)
	if err != nil {
		return &pb.Close{}, status.Errorf(codes.FailedPrecondition, "room not exist")
	}

	for _, conn := range connections {
		wait.Add(1)
		grpclog.Infof("connection check %s", conn.id)
		go func(msg *pb.Message, conn *Connection) {
			if conn.active {
				err = conn.stream.SendMsg(msg)
				if err != nil {
					conn.active = false
					conn.error <- err
					disconnected = true
				}
			}
			defer wait.Done()
		}(msg, conn)
	}

	go func() {
		wait.Wait()
		if true == disconnected {
			//이게 문제 없나...
			//Connection 에 대한 가바지 컬렉션도 잘 작동 할까?
			var newConn []*Connection
			for _, conn := range connections {
				if conn.active == true {
					newConn = append(newConn, conn)
				} else {
					log.Printf("found disconnected user %s", conn.id)
				}
			}
			connections = newConn
		}
		close(done)
	}()
	<-done
	return &pb.Close{}, nil
}

func (s *ChattingServer) GetRoomList(_ context.Context, _ *pb.RoomRequest) (*pb.RoomList, error) {
	//md, ok := metadata.FromIncomingContext(ctx)
	//if ok {
	//	authorizationHeader := md.Get("authorization")
	//	log.Print(authorizationHeader)
	//}

	//var chatRoomList []*pb.ChatRoom
	//
	//for _, v := range s.WaitingRooms {
	//
	//	chatRoomList = append(chatRoomList, &pb.ChatRoom{
	//		Id:   v.Id,
	//		Name: v.Name,
	//	})
	//}

	return &pb.RoomList{
		RoomList: nil,
	}, nil

	//return nil, status.Errorf(codes.Unimplemented, "method GetChatRoomList not implemented")
}

func (s *ChattingServer) CreateRoom(_ context.Context, req *pb.RoomRequest) (*pb.RoomResponse, error) {

	roomId, err := roomManager.CreateNewRoom(req.Room.GetName())
	if err != nil {
		return nil, err
	}

	return &pb.RoomResponse{
		Status:  0,
		Message: "create success",
		Room:    &pb.Room{RoomId: roomId},
	}, nil
}

func (s *ChattingServer) JoinRoom(_ context.Context, req *pb.RoomRequest) (*pb.RoomResponse, error) {
	roomId := req.Room.GetRoomId()
	//if s.WaitingRooms[roomName] == nil {
	//	return nil, status.Errorf(codes.NotFound, "room NotFound")
	//}

	//roomManager.JoinRoom(roomId, req.User, conn)

	room, err := roomManager.GetRoomInfo(roomId)
	if err != nil {
		return nil, err
	}
	return &pb.RoomResponse{
		Status:  0,
		Message: "create success",
		Room:    room,
	}, nil
}

func (s *ChattingServer) LeaveRoom(context.Context, *pb.RoomRequest) (*pb.RoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveChatRoom not implemented")
}
