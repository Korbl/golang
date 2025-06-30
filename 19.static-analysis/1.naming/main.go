// Linting requirement
package main

import (
	"fmt"
	"log"
)

func main() {
	helloFunc()
}

func helloFunc() {
	_, err := fmt.Println("hi")
	if err != nil {
		log.Default()
	}
}
