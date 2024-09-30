package main

import (
	"context"
	"cs739-kv-store/consts"
	pb "cs739-kv-store/proto/kv739" // Import the generated package
	"cs739-kv-store/repository"
	"cs739-kv-store/service"
	"fmt"
)

// Define a struct that implements the KeyValueServiceServer interface.
type server struct {
	pb.UnimplementedKVStoreServiceServer

	memoryRepo *repository.MemoryRepo
	rdsRepo    *repository.RDSRepo
}

// Get Implement the Get method.
func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	value, found, err := service.NewGetService(s.memoryRepo, s.rdsRepo).GetByKey(ctx, req.Key)
	if err != nil {
		return &pb.GetResponse{Status: consts.InternalError}, err
	}
	if !found {
		return &pb.GetResponse{Status: consts.KeyNotFound}, nil
	}

	return &pb.GetResponse{Status: consts.Success, Value: value}, nil // Key found
}

// Put Implement the Put method.
func (s *server) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	fmt.Printf("Storing key: %s with value: %s\n", req.Key, req.Value)
	oldValue, found, err := service.NewPutService(s.memoryRepo, s.rdsRepo).Put(ctx, req.Key, req.Value)
	if err != nil {
		return &pb.PutResponse{Status: consts.InternalError}, err
	}
	if !found {
		return &pb.PutResponse{Status: consts.KeyNotFound}, nil
	}
	return &pb.PutResponse{Status: consts.Success, OldValue: oldValue}, nil
}
