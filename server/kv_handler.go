package main

import (
	"context"
	"cs739-kv-store/consts"
	"cs739-kv-store/pkg"
	pb "cs739-kv-store/proto/kv739" // Import the generated package
	"cs739-kv-store/repository"
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
	//node  *server_raft.Node
	kv *pkg.KV
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

	ok, value, found := s.kv.Get(req.Key)
	if !ok {
		return &pb.GetResponse{Status: consts.InternalError}, nil
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
	oldValue, found, ok := s.kv.Put(req.Key, req.Value)
	if !ok {
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

func startKVServer(kv *pkg.KV, address string) {
	log.Printf("Starting KV server on address %s...\n", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterKVStoreServiceServer(grpcServer, &server{
		mutex: sync.Mutex{},
		kv:    kv,
	})

	log.Printf("Server is running on address %s...\n", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Close(ctx context.Context, req *pb.CloseRequest) (*pb.CloseResponse, error) {
	if req.Clean == 1 {
		// Graceful termination
		log.Println("Graceful termination initiated...")

		// Lock to prevent new requests while shutting down
		s.mutex.Lock()

		// Flush any in-memory state (this is just an example, actual logic would depend on your state handling)
		err := s.flushState()
		if err != nil {
			log.Printf("Failed to flush state: %v", err)
			return &pb.CloseResponse{Status: -1}, err
		}

		// Notify other services about the shutdown (if needed)
		s.notifyOtherServices()

		// Log shutdown and exit after a short delay to allow cleanup
		log.Println("Shutting down gracefully...")
		go func() {
			time.Sleep(1 * time.Second) // Give some time for cleanup
			os.Exit(0)                  // Exit after cleanup
		}()

		return &pb.CloseResponse{Status: 0}, nil
	} else {
		// Immediate termination
		log.Println("Immediate termination initiated...")

		// Terminate the process immediately without flushing state or notifying others
		go func() {
			os.Exit(0)
		}()
		return &pb.CloseResponse{Status: 0}, nil
	}
}

func (s *server) flushState() error {
	// Implement logic to persist in-memory state to disk or DB
	log.Println("Flushing in-memory state to persistent storage...")
	return repository.MemoryRepository.Flush()
}

// notifyOtherServices simulates notifying other nodes about the shutdown
func (s *server) notifyOtherServices() {
	// Implement logic to notify other nodes (if required)
	log.Println("Notifying other services about the shutdown...")
	// Simulate delay for notifications
	time.Sleep(500 * time.Millisecond)
	log.Println("Other services notified successfully.")
}
