package main

import (
	"context"
	"fmt"
	pb "load_balancer/proto/kv739" // Import the generated package
)

// Define a struct that implements the KeyValueServiceServer interface.
type server struct {
	pb.UnimplementedKVStoreServiceServer
}

// Get Implement the Get method.
func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	fmt.Printf("Getting key: %s\n", req.Key)
	client := serverPool.LoadBalance()
	return client.Get(ctx, req)
}

// Put Implement the Put method.
func (s *server) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	fmt.Printf("Storing key: %s with value: %s\n", req.Key, req.Value)
	client := serverPool.LoadBalance()
	return client.Put(ctx, req)
}

func (s *server) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Message: "pong"}, nil
}
