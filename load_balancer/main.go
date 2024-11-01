package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"load_balancer/consts"
	"load_balancer/models"
	pb "load_balancer/proto/kv739"
	"load_balancer/utils"
	"log"
	"net"
	"strconv"
	"time"
)

var (
	port     int
	serverIp string

	serverPool *models.ServerPool
)

func main() {
	// Parse command-line arguments
	flag.IntVar(&port, "port", 8080, "Server port")
	flag.StringVar(&serverIp, "ip", "localhost", "Server IP")
	flag.Parse()

	var IDs []uint64
	servers := make(map[uint64]string)
	if err := utils.ReadConfigFile(consts.KVServerListFileName, &IDs, servers); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	serverPool = models.NewServerPool(IDs, servers)
	serverPool.Connect()
	defer serverPool.Close()

	go serverPool.HealthCheck(5 * time.Second)

	lis, err := net.Listen("tcp", serverIp+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterKVStoreServiceServer(grpcServer, &server{})

	fmt.Printf("Load balancer started on :%d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
