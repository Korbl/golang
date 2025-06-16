package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	// write a program that sums all number [0,n] and prints
	// the result.
	n := s1.Intn(100)
	sum := 0
	for x := 0; x <= n; x++ {
		sum += x
		fmt.Printf("%d plus %d\n", x, sum)

	}
	fmt.Println("n", n)
}
