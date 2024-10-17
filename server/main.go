package main

import (
	server_raft "cs739-kv-store/raft"
	"database/sql"
	"flag"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"go.etcd.io/etcd/raft/v3"
)

var (
	port          int
	serverIp      string
	nodeID        uint64
	kvAddresses   []string
	raftAddresses []string
	raftToKV      map[string]string
	db            *sql.DB
)

func main() {
	// Parse command-line arguments
	flag.IntVar(&port, "port", 50051, "Server port")
	flag.StringVar(&serverIp, "ip", "localhost", "Server IP")
	flag.Uint64Var(&nodeID, "id", 1, "Node ID")
	flag.Parse()

	initRaftConfig()
	initKVConfig()
	initDB(nodeID)
	defer db.Close()

	peers := make([]raft.Peer, 0)
	for i := range raftAddresses {
		if uint64(i+1) == nodeID {
			continue
		}
		peers = append(peers, raft.Peer{ID: uint64(i + 1)})
	}

	node := server_raft.NewRaftNode(nodeID, peers, raftAddresses)

	kvAddress := serverIp + ":" + strconv.Itoa(port)
	if len(kvAddresses) > 0 {
		kvAddress = kvAddresses[nodeID-1]
	}
	go startRaftServer(node, raftAddresses[nodeID-1])

	// Start processing Ready channel from Raft node
	go startKVServer(node, kvAddress)
	go node.ProcessReady()

	// Block and wait for exit signals or errors
	select {}
}
