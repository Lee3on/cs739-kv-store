package models

import (
	"context"
	pb "load_balancer/proto/kv739"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
)

type ServerPool struct {
	IDs             []uint64
	IdToServers     map[uint64]string
	IdToConnections map[uint64]*grpc.ClientConn
	IdToClient      map[uint64]pb.KVStoreServiceClient
	current         uint32
	Health          map[uint64]bool
	AddressToID     map[string]uint64
	globalNextID    uint64

	mu sync.Mutex
}

func NewServerPool(IDs []uint64, servers map[uint64]string) *ServerPool {
	return &ServerPool{
		IDs:             IDs,
		IdToServers:     servers,
		IdToConnections: make(map[uint64]*grpc.ClientConn, len(servers)),
		IdToClient:      make(map[uint64]pb.KVStoreServiceClient, len(servers)),
		Health:          make(map[uint64]bool, len(servers)),
		AddressToID:     make(map[string]uint64),
		globalNextID:    IDs[len(IDs)-1] + 1,
		mu:              sync.Mutex{},
	}
}

func (p *ServerPool) Connect() {
	// Connect to all servers in the server list
	for id, server := range p.IdToServers {
		conn, err := grpc.Dial(server, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}

		p.IdToConnections[id] = conn
		p.IdToClient[id] = pb.NewKVStoreServiceClient(conn)
		p.AddressToID[server] = id
	}
}

// Close closes all IdToConnections to servers in the pool
func (p *ServerPool) Close() {
	for _, conn := range p.IdToConnections {
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close connection: %v", err)
		}
	}
}

func (p *ServerPool) getNextServer() pb.KVStoreServiceClient {
	count := 0
	for {
		if count == len(p.IdToServers) {
			log.Fatalf("All servers are unhealthy")
		}
		// Atomically increment and get the next server index
		next := atomic.AddUint32(&p.current, 1)
		index := next % uint32(len(p.IDs))
		id := p.IDs[index]

		// Check if the server is healthy
		if p.Health[id] {
			return p.IdToClient[id]
		}
		// If not healthy, continue to the next server
		log.Printf("Skipping unhealthy server: %s\n", p.IdToServers[id])
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
		for id, client := range p.IdToClient {
			// Perform a Health check on the server (using a Ping method or similar)
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			resp, err := client.Ping(ctx, &pb.PingRequest{})
			cancel()

			// Update Health status based on response
			if err == nil && resp.GetMessage() == "pong" {
				//log.Printf("Server %s is healthy", p.IdToServers[id])
				p.Health[id] = true
			} else {
				log.Printf("Server %s is unhealthy: err=%v, message=%s", p.IdToServers[id], err, resp.GetMessage())
				p.Health[id] = false
			}
		}
	}
}

func (p *ServerPool) NextID() uint64 {
	p.mu.Lock()
	defer p.mu.Unlock()

	id := p.globalNextID
	p.globalNextID += 1
	return id
}

func (p *ServerPool) GetClientByAddress(address string) pb.KVStoreServiceClient {
	return p.IdToClient[p.AddressToID[address]]
}

func (p *ServerPool) AddServer(id uint64, address string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.IDs = append(p.IDs, id)
	p.IdToServers[id] = address
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	p.IdToConnections[id] = conn
	p.IdToClient[id] = pb.NewKVStoreServiceClient(conn)
	p.AddressToID[address] = id
}

func (p *ServerPool) RemoveServer(address string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	id := p.AddressToID[address]
	for i, v := range p.IDs {
		if v == id {
			p.IDs = append(p.IDs[:i], p.IDs[i+1:]...)
			break
		}
	}
	delete(p.AddressToID, address)
	delete(p.IdToServers, id)
	delete(p.IdToConnections, id)
	delete(p.IdToClient, id)
	delete(p.Health, id)
}
