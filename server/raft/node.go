package raft

import (
	"context"
	"cs739-kv-store/models"
	"cs739-kv-store/repository"
	"cs739-kv-store/service"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	pb "cs739-kv-store/proto/kv739"

	raft "go.etcd.io/etcd/raft/v3"
	"go.etcd.io/etcd/raft/v3/raftpb"
)

type Node struct {
	ID        uint64
	RaftNode  raft.Node
	Transport *Transport
	//storage   *BoltStorage
	storage   *raft.MemoryStorage
	Addresses []string
}

func NewRaftNode(id uint64, peers []raft.Peer, addresses []string) *Node {
	//storage, err := NewBoltStorage(fmt.Sprintf("./storage/raft_log/log_%d.db", id))
	//if err != nil {
	//	log.Fatalf("Failed to create Bolt storage: %v", err)
	//}
	storage := raft.NewMemoryStorage()

	// Load the initial state from persistent storage
	hardState, _, err := storage.InitialState()
	if err != nil {
		log.Fatalf("Failed to load initial state: %v", err)
	}

	config := &raft.Config{
		ID:              id,      // 节点的唯一ID
		ElectionTick:    10,      // 选举超时计时器
		HeartbeatTick:   1,       // 心跳间隔
		Storage:         storage, // 存储日志的存储器
		MaxSizePerMsg:   4096,    // 每条消息的最大大小
		MaxInflightMsgs: 256,     // 最大未确认消息数
	}

	var raftNode raft.Node

	// Check if we have an existing state or not
	if hardState.Commit == 0 && hardState.Term == 0 && hardState.Vote == 0 {
		// Start a new Raft node if no state exists
		raftNode = raft.StartNode(config, peers)
	} else {
		// Restart the Raft node if a previous state exists
		raftNode = raft.RestartNode(config)
	}

	node := &Node{
		ID:        id,
		RaftNode:  raftNode,
		storage:   storage,
		Addresses: addresses,
		Transport: NewTransport(peers, addresses),
	}
	go node.monitorState()
	return node
}

// ProcessReady handles messages from the Raft Ready channel
func (n *Node) ProcessReady() {
	for {
		ready := <-n.RaftNode.Ready()

		// Log messages sent to other nodes
		for _, msg := range ready.Messages {
			if msg.Type == raftpb.MsgVote {
				log.Printf("Node %d is requesting votes from node %d in term %d", msg.From, msg.To, msg.Term)
			}
			if msg.Type == raftpb.MsgVoteResp {
				log.Printf("Node %d received a vote from node %d", msg.To, msg.From)
			}
		}

		//// Persist the entries and hard state only if there's data to save
		//if len(ready.Entries) > 0 || !raft.IsEmptyHardState(ready.HardState) {
		//	if err := n.storage.Save(ready.HardState, ready.Entries); err != nil {
		//		log.Printf("Failed to persist log entries: %v", err)
		//		return
		//	}
		//}
		// Persist the entries and hard state only if there's data to save
		if len(ready.Entries) > 0 || !raft.IsEmptyHardState(ready.HardState) {
			// Save the hard state and entries into memory storage
			if err := n.storage.Append(ready.Entries); err != nil {
				log.Printf("Failed to append log entries: %v", err)
				return
			}

			// Optionally save the hard state separately if it's not empty
			if !raft.IsEmptyHardState(ready.HardState) {
				if err := n.storage.SetHardState(ready.HardState); err != nil {
					log.Printf("Failed to save hard state: %v", err)
				}
			}
		}

		// Send messages to other nodes using the Transport
		log.Printf("Sending messages to other nodes: %v", ready.Messages)
		n.Transport.Send(ready.Messages)

		// Apply committed entries to the state machine
		for _, entry := range ready.CommittedEntries {
			n.Apply(entry)
		}

		// Notify Raft that the node has processed the Ready struct
		n.RaftNode.Advance()
	}
}

func (n *Node) HandlePutRequest(req *pb.PutRequest) (string, error) {
	if n.RaftNode.Status().RaftState != raft.StateLeader {
		// Redirect client to the leader
		leader := n.RaftNode.Status().Lead
		// Ensure the leader is valid
		if leader == raft.None {
			return "", fmt.Errorf("no leader elected yet")
		}
		// Return the leader address to the client
		return n.Addresses[leader-1], nil
	} else {
		// Process the put request
		command := &models.Command{
			Key:   req.Key,
			Value: req.Value,
		}
		data, err := json.Marshal(command)
		if err != nil {
			log.Fatalf("Failed to marshal command: %v", err)
			return "", err
		}
		if err = n.RaftNode.Propose(context.Background(), data); err != nil {
			log.Fatalf("Failed to propose data: %v", err)
			return "", err
		}
	}

	return "", nil
}

func (n *Node) Apply(entry raftpb.Entry) {
	// Only process normal entries (application data)
	if entry.Type == raftpb.EntryNormal && len(entry.Data) > 0 {
		log.Printf("Applying entry: %s\n", string(entry.Data))

		var command models.Command
		if err := json.Unmarshal(entry.Data, &command); err != nil {
			log.Fatalf("Failed to unmarshal command: %v", err)
			return
		}

		// Apply the command to the key-value store
		_, _, err := service.NewPutService(repository.MemoryRepository, repository.RDSRepository).Put(context.Background(), command.Key, command.Value)
		if err != nil {
			log.Printf("Failed to put key: %s with value: %s: %v", command.Key, command.Value, err)
			return
		}
	} else if entry.Type == raftpb.EntryConfChange {
		// Handle configuration change entries
		var cc raftpb.ConfChange
		if err := cc.Unmarshal(entry.Data); err != nil {
			log.Fatalf("Failed to unmarshal ConfChange: %v", err)
			return
		}

		// Apply the configuration change to the Raft node
		//n.applyConfChange(cc)
	}
}

func (n *Node) Close() {
	log.Println("Node is shutting down, notifying other Raft nodes...")

	// Notify other Raft nodes about the shutdown
	n.Transport.NotifyShutdown(n.ID)

	// Perform any other shutdown logic, such as state flushing
	n.flushState()

	// Exit the process after some delay
	time.Sleep(1 * time.Second)
	os.Exit(0)
}

func (n *Node) flushState() {
	// Flush any in-memory state to disk
	log.Printf("Flushing state to disk...")
}

func (n *Node) applyConfChange(cc raftpb.ConfChange) {
	log.Printf("Applying configuration change: %+v", cc)

	// Apply the configuration change to the Raft node
	n.RaftNode.ApplyConfChange(cc)

	// Update the transport layer to reflect the new cluster configuration
	switch cc.Type {
	case raftpb.ConfChangeAddNode:
		// Add the new node to the transport layer
		n.Transport.AddPeer(cc.NodeID, n.Addresses[cc.NodeID-1])
		log.Printf("Node %d added to the cluster", cc.NodeID)

	case raftpb.ConfChangeRemoveNode:
		// Remove the node from the transport layer
		if cc.NodeID == n.ID {
			log.Println("This node is being removed from the cluster. Shutting down.")
			n.Close()
		} else {
			n.Transport.RemovePeer(cc.NodeID)
			log.Printf("Node %d removed from the cluster", cc.NodeID)
		}

	default:
		log.Printf("Unknown configuration change type: %v", cc.Type)
	}
}

func (n *Node) monitorState() {
	ticker := time.NewTicker(1 * time.Second) // Check state every second
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Check and log the current state of the node
			status := n.RaftNode.Status()
			switch status.RaftState {
			case raft.StateFollower:
				log.Printf("Node %d is a Follower in term %d", n.ID, status.Term)
			case raft.StateCandidate:
				log.Printf("Node %d is a Candidate in term %d", n.ID, status.Term)
			case raft.StateLeader:
				log.Printf("Node %d is the Leader in term %d", n.ID, status.Term)
			}
		}
	}
}
