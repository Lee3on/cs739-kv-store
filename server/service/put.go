package service

import (
	"context"
	"log"
)

func Put(ctx context.Context, key string, value string) (string, error) {
	// Placeholder logic for putting a key-value pair.
	oldValue := "old_value"
	log.Printf("Storing key: %s with value: %s\n", key, value)
	return oldValue, nil
}
