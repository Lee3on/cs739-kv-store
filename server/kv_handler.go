package main

import (
	"context"
	"cs739-kv-store/consts"
	pb "cs739-kv-store/proto/kv739" // Import the generated package
	"cs739-kv-store/raft"
	"cs739-kv-store/service"
	"cs739-kv-store/utils"
	"go.etcd.io/etcd/raft/v3/raftpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

// Define a struct that implements the KeyValueServiceServer interface.
type server struct {
	pb.UnimplementedKVStoreServiceServer
	mutex sync.Mutex

	kv          *service.Kvstore
	confChangeC chan<- raftpb.ConfChange
	raftNode    *raft.RaftNode
	//raftWrapper *raft.Wrapper
}

func startKVServer(kv *service.Kvstore, address string, raftNode *raft.RaftNode, confChangeC chan<- raftpb.ConfChange, errorC <-chan error) {
	log.Printf("Starting KV server on address %s...\n", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterKVStoreServiceServer(grpcServer, &server{
		mutex:       sync.Mutex{},
		kv:          kv,
		confChangeC: confChangeC,
		raftNode:    raftNode,
	})

	go func() {
		log.Printf("Server is running on address %s...\n", address)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// exit when raft goes down
	if err, ok := <-errorC; ok {
		log.Fatal(err)
	}
}

// Get Implement the Get method.
func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	//value, found, err := service.NewGetService(repository.MemoryRepository, repository.RDSRepository).GetByKey(ctx, req.Key)
	//if err != nil {
	//	return &pb.GetResponse{Status: consts.InternalError}, err
	//}
	//if !found {
	//	return &pb.GetResponse{Status: consts.KeyNotFound}, nil
	//}

	log.Printf("Processing get request for key: %s, id: %d\n", req.Key, nodeID)
	value, found, err := s.kv.Get(req.Key)
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
	//log.Printf("Storing key: %s with value: %s\n", req.Key, req.Value)
	//leaderAddress, err := s.node.HandlePutRequest(req)
	//if err != nil {
	//	return &pb.PutResponse{Status: consts.InternalError}, err
	//}
	//if leaderAddress != "" {
	//	log.Printf("Redirecting to leader: %s\n", raftToKV[leaderAddress])
	//	return &pb.PutResponse{Status: consts.Redirect, LeaderAddress: raftToKV[leaderAddress]}, nil
	//}

	log.Printf("Processing put request for key: %s, value: %s\n", req.Key, req.Value)
	//oldValue, found, err := service.NewPutService(repository.MemoryRepository, repository.RDSRepository).Put(ctx, req.Key, req.Value)
	//if err != nil {
	//	return &pb.PutResponse{Status: consts.InternalError}, err
	//}
	//if !found {
	//	return &pb.PutResponse{Status: consts.KeyNotFound}, nil
	//}
	if !s.raftNode.IsLeader() {
		// Redirect client to the leader
		leader := s.raftNode.GetLeader()
		return &pb.PutResponse{Status: consts.Redirect, LeaderAddress: kvAddresses[leader]}, nil
	}
	oldValue, found, err := s.kv.Put(req.Key, req.Value)
	if err != nil {
		return &pb.PutResponse{Status: consts.InternalError}, nil
	}
	if !found {
		return &pb.PutResponse{Status: consts.KeyNotFound}, nil
	}
	return &pb.PutResponse{Status: consts.Success, OldValue: oldValue}, nil
}

func (s *server) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Message: "pong"}, nil
}

func (s *server) Close(ctx context.Context, req *pb.CloseRequest) (*pb.CloseResponse, error) {
	if req.Clean == 1 {
		// Graceful termination
		log.Println("Graceful termination initiated...")

		// Lock to prevent new requests while shutting down
		s.mutex.Lock()

		// Flush any in-memory state (this is just an example, actual logic would depend on your state handling)
		if err := s.kv.Flush(); err != nil {
			return &pb.CloseResponse{Status: consts.InternalError}, err
		}
		s.raftNode.Stop()

		// Log shutdown and exit after a short delay to allow cleanup
		log.Println("Shutting down gracefully...")
		go func() {
			time.Sleep(1 * time.Second) // Give some time for cleanup
			os.Exit(0)                  // Exit after cleanup
		}()

		return &pb.CloseResponse{Status: consts.Success}, nil
	} else {
		// Immediate termination
		log.Println("Immediate termination initiated...")

		// Terminate the process immediately without flushing state or notifying others
		go func() {
			os.Exit(0)
		}()
		return &pb.CloseResponse{Status: consts.Success}, nil
	}
}

func (s *server) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	// Start the cluster
	if !s.raftNode.IsLeader() {
		// Redirect client to the leader
		leader := s.raftNode.GetLeader()
		return &pb.StartResponse{Status: consts.Redirect, LeaderAddress: kvAddresses[leader]}, nil
	}
	if req.New != 1 {
		return &pb.StartResponse{Status: consts.InternalError}, nil
	}

	cc := raftpb.ConfChange{
		Type:    raftpb.ConfChangeAddNode,
		NodeID:  req.Id,
		Context: []byte(utils.GenRaftAddr(req.Id)),
	}
	s.confChangeC <- cc

	return &pb.StartResponse{Status: consts.Success}, nil
}

func (s *server) Leave(ctx context.Context, req *pb.LeaveRequest) (*pb.LeaveResponse, error) {
	// Leave the cluster
	if !s.raftNode.IsLeader() {

		// Redirect client to the leader
		leader := s.raftNode.GetLeader()
		return &pb.LeaveResponse{Status: consts.Redirect, LeaderAddress: kvAddresses[leader]}, nil
	}

	cc := raftpb.ConfChange{
		Type:   raftpb.ConfChangeRemoveNode,
		NodeID: req.Id,
	}
	s.confChangeC <- cc

	if req.Clean == 1 {
		// Lock to prevent new requests while shutting down
		s.mutex.Lock()
		defer s.mutex.Unlock()

		// Flush any in-memory state (this is just an example, actual logic would depend on your state handling)
		if err := s.kv.Flush(); err != nil {
			return &pb.LeaveResponse{Status: consts.InternalError}, err
		}

		return &pb.LeaveResponse{Status: consts.Success}, nil
	}
	return &pb.LeaveResponse{Status: consts.Success}, nil
}
