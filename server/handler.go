package main

import (
	"context"
	pb "cs739-kv-store/proto/kv739" // Import the generated package
	"cs739-kv-store/service"
	"fmt"
)

// Define a struct that implements the KeyValueServiceServer interface.
type server struct {
	pb.UnimplementedKeyValueServiceServer
}

// Init Implement the Init method.
func (s *server) Init(ctx context.Context, req *pb.InitRequest) (*pb.InitResponse, error) {
	// Example logic for Init method
	fmt.Printf("Server initialized with: %s\n", req.ServerName)
	err := service.Init(ctx)
	if err != nil {
		return &pb.InitResponse{Status: 1}, err
	}
	return &pb.InitResponse{Status: 0}, nil
}

// Shutdown Implement the Shutdown method.
func (s *server) Shutdown(ctx context.Context, req *pb.ShutdownRequest) (*pb.ShutdownResponse, error) {
	fmt.Println("Server shutting down.")
	err := service.Shutdown(ctx)
	if err != nil {
		return &pb.ShutdownResponse{Status: 1}, err
	}
	return &pb.ShutdownResponse{Status: 0}, nil
}

// Get Implement the Get method.
func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	// Placeholder logic for retrieving a value.
	if req.Key == "example" {
		return &pb.GetResponse{Status: 0, Value: "example_value"}, nil
	}
	value, err := service.GetByKey(ctx, req.Key)
	if err != nil {
		return &pb.GetResponse{Status: 1}, err
	}
	return &pb.GetResponse{Status: 1, Value: value}, nil // Key not found
}

// Put Implement the Put method.
func (s *server) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	fmt.Printf("Storing key: %s with value: %s\n", req.Key, req.Value)
	oldValue, err := service.Put(ctx, req.Key, req.Value)
	if err != nil {
		return &pb.PutResponse{Status: 1}, err
	}
	return &pb.PutResponse{Status: 0, OldValue: oldValue}, nil
}
