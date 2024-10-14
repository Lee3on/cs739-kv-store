package main

import (
	"bufio"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"load_balancer/consts"
	"load_balancer/models"
	pb "load_balancer/proto/kv739"
	"log"
	"net"
	"os"
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

	file, err := os.Open(consts.ServerListFileName)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer file.Close()

	var servers []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		servers = append(servers, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	serverPool = models.NewServerPool(servers)
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
