package main

import (
	"bufio"
	"cs739-kv-store/repository"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"cs739-kv-store/consts"
)

func initDB(nodeID uint64) {
	repository.MemoryRepository = repository.NewMemoryRepo(consts.KVStoreCapacity, 3*time.Second)

	var err error
	db, err = sql.Open("sqlite3", fmt.Sprintf("./storage/kv739_%d.db", nodeID))
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Create table if it doesn't exist
	createTableSQL := `CREATE TABLE IF NOT EXISTS kv (
        key TEXT PRIMARY KEY,
        value TEXT
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	repository.RDSRepository = repository.NewRDSRepo(db)
}

func initRaftConfig() {
	file, err := os.Open(consts.RaftServerListFileName)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer file.Close()

	id := uint64(1)
	raftPeers = make(map[uint64]string)
	proposeCs = make(map[uint64]chan string)
	commitCs = make(map[uint64]chan []string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//raftAddresses = append(raftAddresses, scanner.Text())
		raftPeers[id] = scanner.Text()
		proposeCs[id] = make(chan string)
		commitCs[id] = make(chan []string)
		id++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}
}

func initKVConfig() {
	file, err := os.Open(consts.KVServerListFileName)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//raftToKV = make(map[string]string)
	//index := 0
	for scanner.Scan() {
		kvAddresses = append(kvAddresses, scanner.Text())
		//raftToKV[raftAddresses[index]] = kvAddresses[index]
		//index++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}
}
