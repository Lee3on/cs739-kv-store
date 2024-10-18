package main

import (
	pb "cs739-kv-store/proto/kv739"
	"cs739-kv-store/repository"
	"database/sql"
	"flag"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

var (
	port        int
	serverIp    string
	nodeID      uint64
	kvAddresses []string

	db         *sql.DB
	memoryRepo *repository.MemoryRepo
	rdsRepo    *repository.RDSRepo
)

func main() {
	// Parse command-line arguments
	flag.IntVar(&port, "port", 50051, "Server port")
	flag.StringVar(&serverIp, "ip", "localhost", "Server IP")
	flag.Uint64Var(&nodeID, "id", 1, "Node ID")
	flag.Parse()

	initKVConfig()
	initDB(nodeID)
	defer db.Close()

	lis, err := net.Listen("tcp", kvAddresses[nodeID-1])
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterKVStoreServiceServer(grpcServer, &server{nodeID: nodeID, kvAddresses: kvAddresses, mutex: sync.Mutex{}})

	log.Printf("Server is running on address %s...\n", kvAddresses[nodeID-1])
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
