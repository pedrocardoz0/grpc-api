package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-api/api/message"
	"net"
)

type Server struct {
	message.UnimplementedMessageServiceServer
}

func (s *Server) CreateMessage(ctx context.Context, req *message.Message) (*message.Message, error) {
	return req, nil
}

func (s *Server) UpdateMessage(ctx context.Context, req *message.Message) (*message.Message, error) {
	return req, nil
}

func (s *Server) GetMessageById(ctx context.Context, req *message.GetMessageByIdRequest) (*message.Message, error) {
	return &message.Message{
		Id:      req.Id,
		Content: "Example content",
	}, nil
}

func main() {
	listenAddr := ":50051"

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	s := grpc.NewServer()
	message.RegisterMessageServiceServer(s, &Server{})
	reflection.Register(s)

	fmt.Printf("Server listening on %s\n", listenAddr)

	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
