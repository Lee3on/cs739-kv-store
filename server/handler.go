package main

import (
	"context"
	"cs739-kv-store/consts"
	pb "cs739-kv-store/proto/kv739" // Import the generated package
	"cs739-kv-store/repository"
	"cs739-kv-store/service"
	"fmt"
	"os"
	"sync"
	"time"
)

// Define a struct that implements the KeyValueServiceServer interface.
type server struct {
	pb.UnimplementedKVStoreServiceServer
	nodeID      uint64
	kvAddresses []string
	mutex       sync.Mutex
}

// Get Implement the Get method.
func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	value, found, err := service.NewGetService(repository.MemoryRepository, repository.RDSRepository).GetByKey(ctx, req.Key)
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
	oldValue, found, err := service.NewPutService(repository.MemoryRepository, repository.RDSRepository).Put(ctx, req.Key, req.Value)
	if err != nil {
		return &pb.PutResponse{Status: consts.InternalError}, err
	}

	//for i := range s.kvAddresses {
	//	if s.nodeID == uint64(i+1) {
	//		continue
	//	}
	//	err := utils.CopyFile(fmt.Sprintf("./storage/kv739_%d.db", s.nodeID), fmt.Sprintf("./storage/kv739_%d.db", i+1))
	//	if err != nil {
	//		fmt.Printf("Error copying file: %v\n", err)
	//		return &pb.PutResponse{Status: consts.InternalError}, err
	//	}
	//}

	if !found {
		return &pb.PutResponse{Status: consts.KeyNotFound}, nil
	}
	return &pb.PutResponse{Status: consts.Success, OldValue: oldValue}, nil
}

func (s *server) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Message: "pong"}, nil
}

func (s *server) Close(ctx context.Context, req *pb.CloseRequest) (*pb.CloseResponse, error) {
	if req.Clean == 0 {
		os.Exit(0)
	} else {
		s.mutex.Lock()
		tmp := (1000 * len(s.kvAddresses)) / 100
		time.Sleep(time.Millisecond * time.Duration(tmp))
		s.mutex.Unlock()
		os.Exit(0)
	}
	return &pb.CloseResponse{Status: 0}, nil
}
