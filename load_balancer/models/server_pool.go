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
	Servers         []string
	connections     []*grpc.ClientConn
	Clients         []pb.KVStoreServiceClient
	current         uint32
	Health          []bool
	addressToClient map[string]pb.KVStoreServiceClient
	AddressToIndex  map[string]int
}

func NewServerPool(servers []string) *ServerPool {
	return &ServerPool{
		Servers:         servers,
		connections:     make([]*grpc.ClientConn, len(servers)),
		Clients:         make([]pb.KVStoreServiceClient, len(servers)),
		Health:          make([]bool, len(servers)),
		addressToClient: make(map[string]pb.KVStoreServiceClient),
		AddressToIndex:  make(map[string]int),
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
		p.Clients[i] = pb.NewKVStoreServiceClient(conn)
		p.addressToClient[server] = p.Clients[i]
		p.AddressToIndex[server] = i
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
		index := next % uint32(len(p.Clients))

		// Check if the server is healthy
		if p.Health[index] {
			return p.Clients[index]
		}
		// If not healthy, continue to the next server
		log.Printf("Skipping unhealthy server: %s\n", p.Servers[index])
		count += 1
	}
}

func (p *ServerPool) LoadBalance() pb.KVStoreServiceClient {
	return p.getNextServer()
}

// HealthCheck runs a periodic Health check on all servers
func (p *ServerPool) HealthCheck(interval time.Duration) {
	for {
		time.Sleep(interval)
		for i, client := range p.Clients {
			// Perform a Health check on the server (using a Ping method or similar)
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			resp, err := client.Ping(ctx, &pb.PingRequest{})
			cancel()

			// Update Health status based on response
			if err == nil && resp.GetMessage() == "pong" {
				log.Printf("Server %s is healthy", p.Servers[i])
				p.Health[i] = true
			} else {
				log.Printf("Server %s is unhealthy: err=%v, message=%s", p.Servers[i], err, resp.GetMessage())
				p.Health[i] = false
			}
		}
	}
}

func (p *ServerPool) GetClientByAddress(address string) pb.KVStoreServiceClient {
	return p.addressToClient[address]
}

func (p *ServerPool) GetOtherClients() []pb.KVStoreServiceClient {
	clients := make([]pb.KVStoreServiceClient, 0)
	for i, client := range p.Clients {
		if i == int(p.current) {
			continue
		}
		if p.Health[i] {
			clients = append(clients, client)
		}
	}
	return clients
}
