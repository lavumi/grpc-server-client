package service

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	pb "grpc-game-server/pkg/api/proto"
	"log"
	"sync"
	"time"
)

type Connection struct {
	stream pb.ChatService_CreateStreamServer
	id     string
	active bool
	error  chan error
}

type ChatRoom struct {
	Connection []*Connection
	Id         string
	Name       string
	Members    []*pb.User
	Messages   []*pb.Message
}

type ChattingServer struct {
	pb.UnimplementedChatServiceServer
	Rooms map[string]*ChatRoom
	//Connection []*Connection
}

func (s *ChattingServer) Init() {
	s.Rooms = make(map[string]*ChatRoom)
}

func (s *ChattingServer) CreateStream(request *pb.ChatRoomRequest, stream pb.ChatService_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id:     request.User.Id,
		active: true,
		error:  make(chan error),
	}
	room := s.Rooms["test"]
	room.Connection = append(room.Connection, conn)

	log.Printf("Connection Length : %d", len(room.Connection))

	systemUser := pb.User{
		Id:          "1",
		DisplayName: "System",
	}

	timestamp := time.Now()
	joinMessage := pb.Message{
		Id:        systemUser.Id,
		User:      &systemUser,
		Message:   fmt.Sprintf("%v has entered the room", request.User.DisplayName),
		Timestamp: timestamp.String(),
	}

	_, err := s.BroadcastMessage(context.Background(), &joinMessage)
	if err != nil {
		return err
	}

	return <-conn.error
}

func (s *ChattingServer) BroadcastMessage(ctx context.Context, msg *pb.Message) (*pb.Close, error) {
	wait := sync.WaitGroup{}
	done := make(chan int)
	disconnected := false

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		authorizationHeader := md.Get("authorization")
		log.Print(authorizationHeader)
	}

	room := s.Rooms["test"]

	for _, conn := range room.Connection {
		wait.Add(1)
		go func(msg *pb.Message, conn *Connection) {
			if conn.active {
				err := conn.stream.Send(msg)
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
			for _, conn := range room.Connection {
				if conn.active == true {
					newConn = append(newConn, conn)
				} else {
					log.Printf("found disconnected user %s", conn.id)
				}
			}
			room.Connection = newConn
		}
		close(done)
	}()
	<-done
	return &pb.Close{}, nil
}

func (s *ChattingServer) GetChatRoomList(context.Context, *pb.ChatRoomRequest) (*pb.ChatRoomList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatRoomList not implemented")
}

func (s *ChattingServer) CreateChatRoom(context.Context, *pb.ChatRoomRequest) (*pb.ChatRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChatRoom not implemented")
}

func (s *ChattingServer) JoinChatRoom(context.Context, *pb.ChatRoomRequest) (*pb.ChatRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinChatRoom not implemented")
}

func (s *ChattingServer) LeaveChatRoom(context.Context, *pb.ChatRoomRequest) (*pb.ChatRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveChatRoom not implemented")
}
