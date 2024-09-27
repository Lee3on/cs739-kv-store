package service

import "context"

func GetByKey(ctx context.Context, key string) (string, error) {
	return "example_value", nil
}
