package main

import (
	pb "cs739-kv-store/proto/kv739"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

var (
	port     = 6666
	serverIp = "localhost"
)

func main() {
	// Create a new gRPC server.
	lis, err := net.Listen("tcp", serverIp+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterKeyValueServiceServer(grpcServer, &server{})

	log.Printf("Server is running on port %d...\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
