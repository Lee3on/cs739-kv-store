package main

import (
	"cs739-kv-store/pkg"
	"database/sql"
	"flag"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var (
	port          int
	serverIp      string
	nodeID        uint64
	kvAddresses   []string
	raftAddresses []string
	raftPeers     map[uint64]string
	proposeCs     map[uint64]chan string
	commitCs      map[uint64]chan []string
	//raftToKV      map[string]string
	db *sql.DB
)

func main() {
	// Parse command-line arguments
	flag.IntVar(&port, "port", 50051, "Server port")
	flag.StringVar(&serverIp, "ip", "localhost", "Server IP")
	//flag.Uint64Var(&nodeID, "id", 1, "Node ID")
	flag.Parse()

	initRaftConfig()
	initKVConfig()
	//initDB(nodeID)
	//defer db.Close()

	//peers := make([]raft.Peer, 0)
	//for i := range raftAddresses {
	//	if uint64(i+1) == nodeID {
	//		continue
	//	}
	//	peers = append(peers, raft.Peer{ID: uint64(i + 1)})
	//}
	//
	//node := server_raft.NewRaftNode(nodeID, peers, raftAddresses)
	//
	//kvAddress := serverIp + ":" + strconv.Itoa(port)
	//if len(kvAddresses) > 0 {
	//	kvAddress = kvAddresses[nodeID-1]
	//}
	//go startRaftServer(node, raftAddresses[nodeID-1])
	//
	//// Start processing Ready channel from Raft node
	//go startKVServer(node, kvAddress)
	//go node.ProcessReady()

	kvs := make([]*pkg.KV, len(kvAddresses))
	for i := range kvAddresses {
		id := uint64(i + 1)
		_, err := pkg.NewWrapper(id, raftPeers, proposeCs[id], commitCs[id])
		if err != nil {
			log.Printf("[ERROR] %d init with error %v", id, err)
			return
		}
		kvs[i], err = pkg.NewKV(id, proposeCs[id], commitCs[id])
		if err != nil {
			log.Printf("[ERROR] %d init with error %v", id, err)
			return
		}
		go startKVServer(kvs[i], kvAddresses[i])
	}

	// Block and wait for exit signals or errors
	select {}
}
