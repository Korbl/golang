package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Repeat the previous assignment using a switch statement.
	x := s1.Intn(100)
	y := s1.Intn(100)
	switch {
	case x < y:
		fmt.Printf("x %d is smaller than y  %d", x, y)
	case x > y:
		fmt.Printf("y %d is smaller than x %d", y, x)

	}
}
