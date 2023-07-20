package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	proto "grpc-game-server/pkg/api/proto"
	"os"

	"encoding/hex"
	"log"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var client proto.ChatServiceClient
var wait *sync.WaitGroup

func init() {
	wait = &sync.WaitGroup{}
}

func connect(user *proto.User) error {
	var streamError error

	stream, err := client.CreateStream(context.Background(), &proto.ChatRoomRequest{
		User:     user,
		Action:   0,
		ChatRoom: nil,
	})

	if err != nil {
		return fmt.Errorf("connection failed: %v", err)
	}

	wait.Add(1)
	go func(str proto.ChatService_CreateStreamClient) {
		defer wait.Done()

		for {
			msg, err := str.Recv()
			if err != nil {
				streamError = fmt.Errorf("error reading message: %v", err)
				break
			}

			fmt.Printf("%v : %s\n", msg.User.DisplayName, msg.Message)

		}
	}(stream)

	return streamError
}

func main() {
	timestamp := time.Now()
	done := make(chan int)

	//name := flag.String("N", "Anon", "The name of the user")
	//flag.Parse()

	fmt.Printf("input your name\n")
	nameScanner := bufio.NewScanner(os.Stdin)
	nameScanner.Scan()
	name := nameScanner.Text()

	id := sha256.Sum256([]byte(timestamp.String() + name))

	conn, err := grpc.Dial("localhost:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//conn, err := grpc.Dial("112.170.43.143:15829", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Couldnt connect to service: %v", err)
	}
	fmt.Printf("Chatting Server Connected\n")

	client = proto.NewChatServiceClient(conn)
	user := &proto.User{
		Id:          hex.EncodeToString(id[:]),
		DisplayName: name,
	}

	err = connect(user)

	if err != nil {
		log.Fatalf("Couldnt connect to service: %v", err)
		return
	}

	md := metadata.Pairs("authorization", "my_token") // 예시로 'authorization' 헤더에 Bearer 토큰을 설정
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	wait.Add(1)
	go func() {
		defer wait.Done()

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			msg := &proto.Message{
				Id:        user.Id,
				User:      user,
				Message:   scanner.Text(),
				Timestamp: timestamp.String(),
			}

			_, err := client.BroadcastMessage(ctx, msg)
			if err != nil {
				fmt.Printf("Error Sending Message: %v", err)
				break
			}
		}

	}()

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
}
