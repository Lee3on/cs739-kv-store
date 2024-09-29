package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "cs739-kv-store/proto/kv739" // Import the generated pb package

	"google.golang.org/grpc"
)

func main() {
	// Connect to the gRPC server
	// Read server address from environment variable
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		log.Fatalf("SERVER_ADDRESS environment variable is not set")
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a new client_go for the KeyValueService
	client := pb.NewKVStoreServiceClient(conn)

	// Put a key-value pair into the server
	putResponse, err := client.Put(context.Background(), &pb.PutRequest{
		Key:   "example_key",
		Value: "example_value",
	})
	if err != nil {
		log.Fatalf("Error calling Put: %v", err)
	}
	fmt.Printf("Put Response: Status = %d, OldValue = %s\n", putResponse.Status, putResponse.OldValue)

	// Get the value for the key we just put in
	getResponse, err := client.Get(context.Background(), &pb.GetRequest{
		Key: "example_key",
	})
	if err != nil {
		log.Fatalf("Error calling Get: %v", err)
	}
	if getResponse.Status == 0 {
		fmt.Printf("Get Response: Status = %d, Value = %s\n", getResponse.Status, getResponse.Value)
	} else {
		fmt.Printf("Get Response: Status = %d (key not found)\n", getResponse.Status)
	}

	// Pause for a second to ensure all logs are printed before exiting
	time.Sleep(1 * time.Second)
}
