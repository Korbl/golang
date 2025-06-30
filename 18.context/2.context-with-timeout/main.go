package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Define a context value key type

func main() {
	ctx := context.Background()
	// Use a context with timeout to wait not longer than 5 seconds until the result is returned
	// otherwise log an error message that the operation timed out.
	// Hint: Use a select statement.
	resChan := make(chan (int))
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	go flakyWork(ctx, resChan)

	for {
		x, ok := <-resChan
		if !ok {
			break
		}
		fmt.Println("X", x)
	}
}

func flakyWork(ctx context.Context, resChan chan (int)) {
	defer close(resChan)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Duration(r.Intn(10) * int(time.Second)))

	select {
	case <-ctx.Done():
		log.Println("[flakyWork] ctx Done is received inside flakyWork")
		return
	default:
		resChan <- 42
		return
	}
}
