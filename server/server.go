package main

import (
	"cs739-kv-store/consts"
	pb "cs739-kv-store/proto/kv739"
	"cs739-kv-store/repository"
	"database/sql"
	"flag"
	"log"
	"net"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

var (
	port     int
	serverIp string

	db         *sql.DB
	memoryRepo *repository.MemoryRepo
	rdsRepo    *repository.RDSRepo
)

func initDB() {
	memoryRepo = repository.NewMemoryRepo(consts.KVStoreCapacity, 3*time.Second)

	var err error
	db, err = sql.Open("sqlite3", "kv739.db")
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
	rdsRepo = repository.NewRDSRepo(db)
}

func main() {
	// Parse command-line arguments
	flag.IntVar(&port, "port", 50051, "Server port")
	flag.StringVar(&serverIp, "ip", "localhost", "Server IP")
	flag.Parse()

	initDB()
	defer db.Close()

	lis, err := net.Listen("tcp", serverIp+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterKVStoreServiceServer(grpcServer, &server{
		memoryRepo: memoryRepo,
		rdsRepo:    rdsRepo,
	})

	log.Printf("Server is running on port %d...\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
