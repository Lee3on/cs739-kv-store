package main

import (
	"cs739-kv-store/raft"
	"cs739-kv-store/service"
	"database/sql"
	"flag"
	_ "github.com/mattn/go-sqlite3"
	"go.etcd.io/etcd/raft/v3/raftpb"
)

var (
	port        int
	serverIp    string
	nodeID      uint64
	kvAddresses map[uint64]string
	raftPeers   map[uint64]string
	join        bool
	db          *sql.DB
)

func main() {
	// Parse command-line arguments
	flag.IntVar(&port, "port", 50051, "Server port")
	flag.StringVar(&serverIp, "ip", "localhost", "Server IP")
	flag.Uint64Var(&nodeID, "id", 1, "Node ID")
	flag.BoolVar(&join, "join", false, "Whether to join a new node")
	flag.Parse()

	initDB(nodeID)
	initRaftConfig()
	initKVConfig()

	proposeC := make(chan string)
	confChangeC := make(chan raftpb.ConfChange)
	defer close(confChangeC)
	defer close(proposeC)

	var kvs *service.Kvstore
	getSnapshot := func() ([]byte, error) { return kvs.GetSnapshot() }
	raftNode, commitC, errorC := raft.NewRaftNode(nodeID, raftPeers, join, getSnapshot, proposeC, confChangeC)
	kvs = service.NewKVStore(<-raftNode.SnapshotterReady, proposeC, commitC, errorC, db)
	startKVServer(kvs, kvAddresses[nodeID], raftNode, confChangeC, errorC)

	// Block and wait for exit signals or errors
	select {}
}
