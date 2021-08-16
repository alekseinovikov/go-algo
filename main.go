package main

import (
	"algo/chat"
	context "context"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	s := Server{}
	chat.RegisterChatServiceServer(grpcServer, &s)

	go clentCreation()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func clentCreation() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	client := chat.NewChatServiceClient(conn)
	defer conn.Close()

	resp, err := client.SayHello(context.Background(), &chat.Message{Body: "PIDOR!"})

	if err != nil {
		log.Fatalf("WRONG answer from server: %s", err)
	}

	log.Printf("Answer from server to client: %s", resp.Body)
}

type Server struct {
}

func (s Server) SayHello(ctx context.Context, message *chat.Message) (*chat.Message, error) {
	log.Printf("Received message %s", message.Body)

	return &chat.Message{
		Body: "Hello! Blyad!",
	}, nil
}
