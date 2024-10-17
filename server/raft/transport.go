package raft

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"go.etcd.io/etcd/raft/v3"

	pb "cs739-kv-store/proto/raft"

	"go.etcd.io/etcd/raft/v3/raftpb"
	"google.golang.org/grpc"
)

type Transport struct {
	connections map[uint64]pb.RaftServiceClient
}

// NewTransport creates a new Transport instance
func NewTransport(peers []raft.Peer, addresses []string) *Transport {
	transport := &Transport{
		connections: make(map[uint64]pb.RaftServiceClient),
	}

	// Establish gRPC connections with each peer
	for _, peer := range peers {
		conn, err := grpc.Dial(addresses[peer.ID-1], grpc.WithInsecure())
		if err != nil {
			log.Printf("Failed to connect to peer %d: %v, address: %s", peer.ID, err, addresses[peer.ID-1])
		}
		client := pb.NewRaftServiceClient(conn)
		transport.connections[peer.ID] = client
	}

	return transport
}

// Send sends a Raft message to the appropriate node
func (t *Transport) Send(msgs []raftpb.Message) {
	for _, msg := range msgs {
		client := t.connections[msg.To]

		// Serialize the Raft message
		data, err := msg.Marshal()
		if err != nil {
			log.Printf("Failed to marshal raft message: %v", err)
			continue
		}

		// Send the message via gRPC
		raftMsg := &pb.RaftMessage{
			From: msg.From,
			To:   msg.To,
			Data: data,
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err = client.SendRaftMessage(ctx, raftMsg)
		if err != nil {
			log.Printf("Failed to send raft message to %d: %v", msg.To, err)
		}
	}
}

// NotifyShutdown sends a shutdown notification to all peers
func (t *Transport) NotifyShutdown(nodeID uint64) {
	// Create a shutdown notification message
	notification := &pb.ShutdownNotification{
		NodeId:  nodeID,
		Message: "Node is shutting down.",
	}

	// Serialize the shutdown notification
	data, err := json.Marshal(notification)
	if err != nil {
		log.Printf("Failed to marshal shutdown notification: %v", err)
		return
	}

	// Send the notification to all peers
	for peerID, client := range t.connections {
		raftMsg := &pb.RaftMessage{
			From: nodeID,
			To:   peerID,
			Data: data,
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err = client.SendRaftMessage(ctx, raftMsg)
		if err != nil {
			log.Printf("Failed to send shutdown notification to %d: %v", peerID, err)
		}
	}
}

// AddPeer adds a new peer to the transport layer
func (t *Transport) AddPeer(nodeID uint64, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		log.Fatalf("Failed to connect to new peer %d: %v", nodeID, err)
	}
	client := pb.NewRaftServiceClient(conn)
	t.connections[nodeID] = client
	log.Printf("Added peer %d at %s", nodeID, address)
}

// RemovePeer removes a peer from the transport layer
func (t *Transport) RemovePeer(nodeID uint64) {
	if _, ok := t.connections[nodeID]; ok {
		delete(t.connections, nodeID)
		// Here, you can also close the gRPC connection if necessary
		log.Printf("Removed peer %d", nodeID)
	}
}
