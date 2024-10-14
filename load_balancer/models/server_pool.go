package models

import (
	"context"
	"google.golang.org/grpc"
	pb "load_balancer/proto/kv739"
	"log"
	"sync/atomic"
	"time"
)

type ServerPool struct {
	Servers     []string
	connections []*grpc.ClientConn
	clients     []pb.KVStoreServiceClient
	current     uint32
	health      []bool
}

func NewServerPool(servers []string) *ServerPool {
	return &ServerPool{
		Servers:     servers,
		connections: make([]*grpc.ClientConn, len(servers)),
		clients:     make([]pb.KVStoreServiceClient, len(servers)),
		health:      make([]bool, len(servers)),
	}
}

func (p *ServerPool) Connect() {
	// Connect to all servers in the server list
	for i, server := range p.Servers {
		conn, err := grpc.Dial(server, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}

		p.connections[i] = conn
		p.clients[i] = pb.NewKVStoreServiceClient(conn)
	}
}

// Close closes all connections to servers in the pool
func (p *ServerPool) Close() {
	for _, conn := range p.connections {
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close connection: %v", err)
		}
	}
}

func (p *ServerPool) getNextServer() pb.KVStoreServiceClient {
	count := 0
	for {
		if count == len(p.Servers) {
			log.Fatalf("All servers are unhealthy")
		}
		// Atomically increment and get the next server index
		next := atomic.AddUint32(&p.current, 1)
		index := next % uint32(len(p.clients))

		// Check if the server is healthy
		if p.health[index] {
			return p.clients[index]
		}
		// If not healthy, continue to the next server
		log.Printf("Skipping unhealthy server: %s\n", p.Servers[index])
		count += 1
	}
}

func (p *ServerPool) LoadBalance() pb.KVStoreServiceClient {
	return p.getNextServer()
}

// HealthCheck runs a periodic health check on all servers
func (p *ServerPool) HealthCheck(interval time.Duration) {
	for {
		time.Sleep(interval)
		for i, client := range p.clients {
			// Perform a health check on the server (using a Ping method or similar)
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			resp, err := client.Ping(ctx, &pb.PingRequest{})
			cancel()

			// Update health status based on response
			if err == nil && resp.GetMessage() == "pong" {
				log.Printf("Server %s is healthy", p.Servers[i])
				p.health[i] = true
			} else {
				log.Printf("Server %s is unhealthy: err=%v, message=%s", p.Servers[i], err, resp.GetMessage())
				p.health[i] = false
			}
		}
	}
}
