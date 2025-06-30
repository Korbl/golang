package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	// Speed up the following code using goroutines and a waitgroup.
	for i := 0; i < 1000; i++ {
		go longRunningTask(&wg)
	}

	fmt.Println("Pfff... finally done! That took: ", time.Since(start))
}

func longRunningTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Executing long running task")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Duration(r.Intn(1000) * int(time.Millisecond)))
}
