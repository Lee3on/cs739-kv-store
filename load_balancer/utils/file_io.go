package utils

import (
	"bufio"
	"fmt"
	"load_balancer/consts"
	"log"
	"os"
	"strings"
)

func ReadConfigFile(fileName string, IDs *[]uint64, servers map[uint64]string) error {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error opening config file:", err)
		return err
	}
	defer file.Close()

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
		addr := parts[1]

		*IDs = append(*IDs, id)
		servers[id] = addr
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading config file:", err)
		return err
	}

	return nil
}

func AddInstanceToConfigFile(id uint64, kvAddr string, raftAddr string) error {
	kvConfigFile, err := os.OpenFile(consts.KVServerListFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println("Error opening config file:", err)
		return err
	}
	raftConfigFile, err := os.OpenFile(consts.RaftServerListFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println("Error opening config file:", err)
		return err
	}
	defer kvConfigFile.Close()
	defer raftConfigFile.Close()

	newLine := fmt.Sprintf("%d %s\n", id, kvAddr)
	_, err = kvConfigFile.WriteString(newLine)
	if err != nil {
		log.Println("Error writing to file:", err)
		return err
	}

	newLine = fmt.Sprintf("%d %s\n", id, raftAddr)
	_, err = raftConfigFile.WriteString(newLine)
	if err != nil {
		log.Println("Error writing to file:", err)
		return err
	}

	return nil
}

func RemoveInstanceFromConfigFile(kvAddr string) error {
	kvConfigFile, err := os.OpenFile(consts.KVServerListFileName, os.O_RDWR, 0644)
	if err != nil {
		log.Println("Error opening config file:", err)
		return err
	}
	raftConfigFile, err := os.OpenFile(consts.RaftServerListFileName, os.O_RDWR, 0644)
	if err != nil {
		log.Println("Error opening config file:", err)
		return err
	}
	defer kvConfigFile.Close()
	defer raftConfigFile.Close()

	// Remove the line from the KV config file
	scanner := bufio.NewScanner(kvConfigFile)
	var id uint64
	var lines []string
	for scanner.Scan() {
		var curId uint64
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			log.Println("Invalid line format:", line)
			continue
		}

		if _, err := fmt.Sscanf(parts[0], "%d", &curId); err != nil {
			log.Println("Invalid ID:", parts[0])
			continue
		}

		if parts[1] == kvAddr {
			id = curId
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading config file:", err)
		return err
	}

	if err := os.WriteFile(consts.KVServerListFileName, []byte(strings.Join(lines, "\n")+"\n"), 0644); err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	// Remove the line from the Raft config file
	lines = make([]string, 0)
	scanner = bufio.NewScanner(raftConfigFile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			log.Println("Invalid line format:", line)
			continue
		}

		var lineID uint64
		if _, err := fmt.Sscanf(parts[0], "%d", &lineID); err != nil {
			log.Println("Invalid ID:", parts[0])
			continue
		}

		if lineID == id {
			continue
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading config file:", err)
		return err
	}

	if err := os.WriteFile(consts.RaftServerListFileName, []byte(strings.Join(lines, "\n")+"\n"), 0644); err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}
