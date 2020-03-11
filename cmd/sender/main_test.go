package main

import (
	"context"
	"testing"
)

func TestStream(t *testing.T) {
	ctx := context.Background()
	client := GetClient("127.0.0.1:8080")
	StreamData(ctx, client, 1000, 1, false)
}
