package main

import (
	"context"
	"log"
	"net"

	pb "cs739-kv-store/proto/raft" // Import the generated proto package
	server_raft "cs739-kv-store/raft"

	"go.etcd.io/etcd/raft/v3/raftpb"
	"google.golang.org/grpc"
)

type raftServer struct {
	pb.UnimplementedRaftServiceServer
	node *server_raft.Node // Reference to the local Raft node
}

// SendRaftMessage is the unary gRPC call to receive Raft messages
func (s *raftServer) SendRaftMessage(ctx context.Context, msg *pb.RaftMessage) (*pb.RaftResponse, error) {
	// Deserialize the message and send it to the local Raft node
	var raftMsg raftpb.Message
	err := raftMsg.Unmarshal(msg.Data)
	if err != nil {
		return &pb.RaftResponse{Success: false}, err
	}

	// Pass the Raft message to the local Raft node for processing
	err = s.node.RaftNode.Step(ctx, raftMsg)
	if err != nil {
		return &pb.RaftResponse{Success: false}, err
	}

	return &pb.RaftResponse{Success: true}, nil
}

func startRaftServer(node *server_raft.Node, address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRaftServiceServer(grpcServer, &raftServer{node: node})

	log.Printf("Raft server listening on address %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
