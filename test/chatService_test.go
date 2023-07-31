package test

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	pb "grpc-game-server/pkg/api/proto"
	chat_service "grpc-game-server/pkg/service/chat-service"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener
var testUser *pb.User
var testRoom *pb.Room

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	server := &chat_service.ChattingServer{}
	server.Init()
	pb.RegisterRoomServiceServer(s, server)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	name := "test_user"
	id := sha256.Sum256([]byte("time_time" + name))
	testUser = &pb.User{
		UserId:      hex.EncodeToString(id[:]),
		DisplayName: name,
	}

	testRoom = &pb.Room{
		RoomId:   "",
		Name:     "test_room_1",
		Members:  nil,
		Messages: nil,
	}
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestJoinRoom(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "test_buf_conn", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufcon: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	client := pb.NewRoomServiceClient(conn)

	resp, err := client.CreateRoom(ctx, &pb.RoomRequest{
		User:   testUser,
		Action: 0,
		Room:   testRoom,
	})
	if err != nil {
		t.Fatalf("Create Room failed: %v", err)
	}
	log.Printf("CreateRoomResponse: %+v", resp)
	testRoom.RoomId = resp.Room.RoomId
	resp, err = client.JoinRoom(ctx, &pb.RoomRequest{
		User:   testUser,
		Action: 0,
		Room:   testRoom,
	})
	if err != nil {
		t.Fatalf("Join failed: %v", err)
	}
	log.Printf("JoinRoomResponse: %+v", resp)

	stream, err := client.CreateStream(ctx, &pb.RoomRequest{
		User:   testUser,
		Action: 0,
		Room:   testRoom,
	})
	go func() {
		for {
			msg, err := stream.Recv()
			log.Printf("Receive Message")
			if err != nil {
				break
			}
			//fmt.Printf("%v : %s\n", msg.User.DisplayName, msg.Message)
			log.Printf("%v : %s\n", msg.User.DisplayName, msg.Message)
		}
	}()

	testMsg := &pb.Message{
		RoomId:    testRoom.RoomId,
		User:      testUser,
		Message:   "this is test message",
		Timestamp: "",
	}
	_, err = client.BroadcastMessage(ctx, testMsg)
	if err != nil {
		return
	}

}
