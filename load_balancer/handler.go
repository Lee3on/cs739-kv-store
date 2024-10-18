package main

import (
	"context"
	"load_balancer/consts"
	pb "load_balancer/proto/kv739" // Import the generated package
	"log"
	"time"
)

// Define a struct that implements the KeyValueServiceServer interface.
type server struct {
	pb.UnimplementedKVStoreServiceServer
}

// Get Implement the Get method.
func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("Getting key: %s\n", req.Key)
	client := serverPool.LoadBalance()
	return client.Get(ctx, req)
}

// Put Implement the Put method.
func (s *server) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	log.Printf("Storing key: %s with value: %s\n", req.Key, req.Value)
	//client := serverPool.LoadBalance()
	//go func() {
	//	clients := serverPool.GetOtherClients()
	//	log.Printf("Storing key: %s with value: %s in other servers: %v\n", req.Key, req.Value, clients)
	//	for _, c := range clients {
	//		_, err := c.Put(ctx, req)
	//		if err != nil {
	//			log.Printf("Error storing key: %s with value: %s\n", req.Key, req.Value)
	//		}
	//	}
	//}()
	//return client.Put(ctx, req)
	var finalErr error
	found := false
	oldValue := ""
	for i, client := range serverPool.Clients {
		if !serverPool.Health[i] {
			continue
		}
		resp, err := client.Put(ctx, req)
		if err != nil {
			log.Printf("Error storing key: %s with value: %s\n", req.Key, req.Value)
			finalErr = err
		}
		if resp != nil && resp.Status == int32(consts.Success) {
			found = true
			oldValue = resp.OldValue
		}
		time.Sleep(time.Microsecond * 10)
	}

	if finalErr != nil {
		return &pb.PutResponse{Status: int32(consts.InternalError)}, finalErr
	}
	if !found {
		return &pb.PutResponse{Status: int32(consts.KeyNotFound)}, nil
	}
	return &pb.PutResponse{Status: int32(consts.Success), OldValue: oldValue}, nil
}

func (s *server) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Message: "pong"}, nil
}

func (s *server) Close(ctx context.Context, req *pb.CloseRequest) (*pb.CloseResponse, error) {
	log.Printf("Closing connection to server: %s\n", req.ServerName)
	client := serverPool.GetClientByAddress(req.ServerName)
	_, err := client.Close(ctx, req)
	if err != nil {
		log.Printf("Error closing connection to server: %s\n", req.ServerName)
		return &pb.CloseResponse{Status: int32(consts.InternalError)}, err
	}
	serverPool.Health[serverPool.AddressToIndex[req.ServerName]] = false
	return &pb.CloseResponse{Status: int32(consts.Success)}, nil
}
