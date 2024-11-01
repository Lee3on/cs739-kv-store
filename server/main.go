package main

import (
	"cs739-kv-store/raft"
	"cs739-kv-store/repository"
	"flag"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var (
	port        int
	serverIp    string
	nodeID      uint64
	kvAddresses map[uint64]string
	raftPeers   map[uint64]string
	proposeCs   map[uint64]chan string
	commitCs    map[uint64]chan []string
)

func main() {
	// Parse command-line arguments
	flag.IntVar(&port, "port", 50051, "Server port")
	flag.StringVar(&serverIp, "ip", "localhost", "Server IP")
	flag.Uint64Var(&nodeID, "id", 1, "Node ID")
	flag.Parse()

	initRaftConfig()
	initKVConfig()

	kvs := make(map[uint64]*repository.KV)
	wrapper, err := raft.NewWrapper(nodeID, raftPeers, proposeCs[nodeID], commitCs[nodeID])
	if err != nil {
		log.Printf("[ERROR] %d init with error %v", nodeID, err)
		return
	}
	kvs[nodeID], err = repository.NewKV(nodeID, proposeCs[nodeID], commitCs[nodeID])
	if err != nil {
		log.Printf("[ERROR] %d init with error %v", nodeID, err)
		return
	}
	startKVServer(kvs[nodeID], kvAddresses[nodeID], wrapper)

	// Block and wait for exit signals or errors
	select {}
}
