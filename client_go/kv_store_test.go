package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	pb "cs739-kv-store/proto/kv739" // Import the generated pb package
	"google.golang.org/grpc"
)

// Helper function to create a gRPC client connection
func createClient() (pb.KVStoreServiceClient, *grpc.ClientConn) {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:6666", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	client := pb.NewKVStoreServiceClient(conn)
	return client, conn
}

// TestPut tests the Put operation of the KeyValue store
func TestPut(t *testing.T) {
	client, conn := createClient()
	defer conn.Close()

	// Test: Put a key-value pair
	putResponse, err := client.Put(context.Background(), &pb.PutRequest{
		Key:   "test_key",
		Value: "test_value",
	})
	if err != nil {
		t.Fatalf("Error calling Put: %v", err)
	}
	if putResponse.Status != 1 {
		t.Errorf("Expected Status 1, got %d", putResponse.Status)
	}
}

// TestGet tests the Get operation of the KeyValue store
func TestGet(t *testing.T) {
	client, conn := createClient()
	defer conn.Close()

	// First, Put a key-value pair to make sure we have something to get
	client.Put(context.Background(), &pb.PutRequest{
		Key:   "test_key",
		Value: "test_value",
	})

	// Test: Get the value we just put
	getResponse, err := client.Get(context.Background(), &pb.GetRequest{
		Key: "test_key",
	})
	if err != nil {
		t.Fatalf("Error calling Get: %v", err)
	}
	if getResponse.Status != 0 {
		t.Errorf("Expected Status 0, got %d", getResponse.Status)
	}
	if getResponse.Value != "test_value" {
		t.Errorf("Expected Value 'test_value', got '%s'", getResponse.Value)
	}
}

// TestPutOverwrite tests overwriting a value for the same key
func TestPutOverwrite(t *testing.T) {
	client, conn := createClient()
	defer conn.Close()

	// Put an initial key-value pair
	client.Put(context.Background(), &pb.PutRequest{
		Key:   "overwrite_key",
		Value: "initial_value",
	})

	// Overwrite the value for the same key
	putResponse, err := client.Put(context.Background(), &pb.PutRequest{
		Key:   "overwrite_key",
		Value: "new_value",
	})
	if err != nil {
		t.Fatalf("Error calling Put: %v", err)
	}
	// Verify the new value was stored
	getResponse, err := client.Get(context.Background(), &pb.GetRequest{
		Key: "overwrite_key",
	})
	if err != nil {
		t.Fatalf("Error calling Get: %v", err)
	}
	if getResponse.Value != "new_value" {
		t.Errorf("Expected Value 'new_value', got '%s'", getResponse.Value)
	}

	if putResponse.OldValue != "initial_value" {
		t.Errorf("Expected OldValue 'initial_value', got '%s'", putResponse.OldValue)
	}
	if putResponse.Status != 0 {
		t.Errorf("Expected Status 0, got %d", putResponse.Status)
	}

}

// TestGetNonExistentKey tests attempting to Get a non-existent key
func TestGetNonExistentKey(t *testing.T) {
	client, conn := createClient()
	defer conn.Close()

	// Test: Try to get a key that doesn't exist
	getResponse, err := client.Get(context.Background(), &pb.GetRequest{
		Key: "non_existent_key",
	})
	if err != nil {
		t.Fatalf("Error calling Get: %v", err)
	}
	if getResponse.Status != 1 {
		t.Errorf("Expected Status 1 (key not found), got %d", getResponse.Status)
	}
	if getResponse.Value != "" {
		t.Errorf("Expected empty Value, got '%s'", getResponse.Value)
	}
}

// TestConcurrentPuts tests concurrent Put operations
func TestConcurrentPuts(t *testing.T) {
	client, conn := createClient()
	defer conn.Close()

	// Simulate concurrent puts for the same key
	for i := 0; i < 10; i++ {
		go func(i int) {
			client.Put(context.Background(), &pb.PutRequest{
				Key:   "concurrent_key",
				Value: fmt.Sprintf("value_%d", i),
			})
		}(i)
	}

	// Pause to ensure goroutines finish
	time.Sleep(2 * time.Second)

	// Get the value after the concurrent puts
	getResponse, err := client.Get(context.Background(), &pb.GetRequest{
		Key: "concurrent_key",
	})
	if err != nil {
		t.Fatalf("Error calling Get: %v", err)
	}

	if getResponse.Status != 0 {
		t.Errorf("Expected Status 0, got %d", getResponse.Status)
	}
	// The final value can be any of the concurrent values, so we just check that it's not empty
	if getResponse.Value == "" {
		t.Errorf("Expected a non-empty Value, got empty string")
	}
}

func TestRecovery(t *testing.T) {
	client, conn := createClient()
	defer conn.Close()

	// Test: Get the value we just put
	getResponse, err := client.Get(context.Background(), &pb.GetRequest{
		Key: "test_key",
	})
	if err != nil {
		t.Fatalf("Error calling Get: %v", err)
	}
	if getResponse.Status != 0 {
		t.Errorf("Expected Status 0, got %d", getResponse.Status)
	}
	if getResponse.Value != "test_value" {
		t.Errorf("Expected Value 'test_value', got '%s'", getResponse.Value)
	}
}
