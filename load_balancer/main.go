package main

import (
	"context"
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
	"sync"
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

	forceLeaveC := make(chan string)
	go serverPool.HealthCheck(5*time.Second, forceLeaveC)

	lis, err := net.Listen("tcp", serverIp+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	newServer := &server{
		mutex: sync.Mutex{},
	}
	pb.RegisterKVStoreServiceServer(grpcServer, newServer)

	go func() {
		for serverName := range forceLeaveC {
			log.Printf("Force leaving the server pool for server: %s...\n", serverName)
			resp, err := newServer.Leave(context.TODO(), &pb.LeaveRequest{
				ServerName: serverName,
				Clean:      1,
			})

			if err != nil || resp.Status != consts.Success {
				log.Printf("Failed to leave the server pool: %v\n", err)
			} else {
				log.Printf("Force leave response: %v\n", resp)
			}
		}
	}()

	fmt.Printf("Load balancer started on :%d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
