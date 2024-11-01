package main

import (
	"bufio"
	"cs739-kv-store/consts"
	"fmt"
	"log"
	"os"
	"strings"
)

func initRaftConfig() {
	file, err := os.Open(consts.RaftServerListFileName)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer file.Close()

	raftPeers = make(map[uint64]string)
	proposeCs = make(map[uint64]chan string)
	commitCs = make(map[uint64]chan []string)
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
		proposeCs[id] = make(chan string)
		commitCs[id] = make(chan []string)
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
