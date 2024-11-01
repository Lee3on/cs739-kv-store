package main

import (
	"context"
	"fmt"
	"load_balancer/consts"
	pb "load_balancer/proto/kv739" // Import the generated package
	"load_balancer/utils"
	"log"
	"os/exec"
	"syscall"
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
	client := serverPool.LoadBalance()
	resp, err := client.Put(ctx, req)
	if err != nil {
		log.Printf("Error storing key: %v", err)
		return resp, err
	}

	if resp.Status == consts.Redirect && resp.LeaderAddress != "" {
		// Redirect to the leader
		client = serverPool.GetClientByAddress(resp.LeaderAddress)
		if client == nil {
			return &pb.PutResponse{Status: consts.InternalError}, fmt.Errorf("leader address not found in the server pool. Address: %s", resp.LeaderAddress)
		}
		return client.Put(ctx, req)
	}

	return resp, err
}

func (s *server) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Message: "pong"}, nil
}

func (s *server) Close(ctx context.Context, req *pb.CloseRequest) (*pb.CloseResponse, error) {
	log.Printf("Closing connection to server: %s\n", req.ServerName)
	client := serverPool.GetClientByAddress(req.ServerName)
	if client == nil {
		return &pb.CloseResponse{Status: consts.InternalError}, fmt.Errorf("server address not found in the server pool. Address: %s", req.ServerName)
	}

	serverPool.Health[serverPool.AddressToID[req.ServerName]] = false
	return client.Close(ctx, req)
}

func (s *server) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	log.Printf("Starting server: %s\n", req.ServerName)
	if req.New == 1 {
		newId := serverPool.NextID()
		if err := utils.AddInstanceToConfigFile(newId, req.ServerName, utils.GenRaftAddr(newId)); err != nil {
			return &pb.StartResponse{Status: consts.InternalError}, err
		}
		cmd := exec.Command("./server", "--id", fmt.Sprintf("%d", newId))
		cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
		err := cmd.Start()
		if err != nil {
			return &pb.StartResponse{Status: consts.InternalError}, err
		}
		serverPool.AddServer(newId, req.ServerName)
		client := serverPool.GetClientByAddress(req.ServerName)
		if client == nil {
			return &pb.StartResponse{Status: consts.InternalError}, fmt.Errorf("server address not found in the server pool. Address: %s", req.ServerName)
		}

		serverPool.Health[serverPool.AddressToID[req.ServerName]] = true
		return client.Start(ctx, req)
	} else {
		id, ok := serverPool.AddressToID[req.ServerName]
		if !ok {
			return &pb.StartResponse{Status: consts.InternalError}, fmt.Errorf("server address not found in the server pool. Address: %s", req.ServerName)
		}

		cmd := exec.Command("./server", "--id", fmt.Sprintf("%d", id))
		cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
		err := cmd.Start()
		if err != nil {
			return &pb.StartResponse{Status: consts.InternalError}, err
		}

		serverPool.Health[id] = true
		return &pb.StartResponse{Status: consts.Success}, nil
	}
}

func (s *server) Leave(ctx context.Context, req *pb.LeaveRequest) (*pb.LeaveResponse, error) {
	log.Printf("Stopping server: %s\n", req.ServerName)
	if err := utils.RemoveInstanceFromConfigFile(req.ServerName); err != nil {
		return &pb.LeaveResponse{Status: consts.InternalError}, err
	}
	serverPool.RemoveServer(req.ServerName)
	client := serverPool.GetClientByAddress(req.ServerName)
	if client == nil {
		return &pb.LeaveResponse{Status: consts.InternalError}, fmt.Errorf("server address not found in the server pool. Address: %s", req.ServerName)
	}

	serverPool.Health[serverPool.AddressToID[req.ServerName]] = false
	return client.Leave(ctx, req)
}
