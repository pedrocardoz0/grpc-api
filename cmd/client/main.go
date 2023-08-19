package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"grpc-api/api/message"
)

func main() {
	serverAddress := "localhost:50051"
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := message.NewMessageServiceClient(conn)

	createResponse, err := client.CreateMessage(context.Background(), &message.Message{
		Id:      "1",
		Content: "Hello, World!",
	})
	if err != nil {
		log.Fatalf("CreateMessage failed: %v", err)
	}
	fmt.Printf("Created message: %v\n", createResponse)

	updateResponse, err := client.UpdateMessage(context.Background(), &message.Message{
		Id:      "1",
		Content: "Updated content",
	})
	if err != nil {
		log.Fatalf("UpdateMessage failed: %v", err)
	}
	fmt.Printf("Updated message: %v\n", updateResponse)

	getResponse, err := client.GetMessageById(context.Background(), &message.GetMessageByIdRequest{
		Id: "1",
	})
	if err != nil {
		log.Fatalf("GetMessageById failed: %v", err)
	}
	fmt.Printf("Message by ID: %v\n", getResponse)
}
