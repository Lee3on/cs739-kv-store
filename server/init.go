package main

import (
	"bufio"
	"cs739-kv-store/consts"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

func initDB(nodeID uint64) {
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
}

func initRaftConfig() {
	file, err := os.Open(consts.RaftServerListFileName)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer file.Close()

	raftPeers = make(map[uint64]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		if len(parts) < 2 {
			log.Println("Invalid line format:", line)
			continue
		}

		var id uint64
		if _, err := fmt.Sscanf(parts[0], "%d", &id); err != nil {
			log.Println("Invalid ID:", parts[0])
			continue
		}
		raftPeers[id] = parts[1]
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

	kvAddresses = make(map[uint64]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		if len(parts) < 2 {
			log.Println("Invalid line format:", line)
			continue
		}

		var id uint64
		if _, err := fmt.Sscanf(parts[0], "%d", &id); err != nil {
			log.Println("Invalid ID:", parts[0])
			continue
		}
		kvAddresses[id] = parts[1]
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}
}
