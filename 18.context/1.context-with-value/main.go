package main

import (
	"context"
	"fmt"
)

// Define a context value key type

func main() {
	ctx := context.Background()
	var requestIDkey string
	ctx = context.WithValue(ctx, requestIDkey, "requestID")
	// Use the context value key to add some value to the context.

	// Afterwards retrieve your value from the context and print it.
	fmt.Println(ctx.Value(requestIDkey))
}
