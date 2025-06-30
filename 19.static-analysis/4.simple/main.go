// package main docs.
package main

import (
	"fmt"
	"log"
)

// Run make lint and see if you can fix the lint errors.
// Hint try running gofumpt.
func main() {
	_, err := fmt.Println(isTrue(true))
	if err != nil {
		log.Default()
	}
	_, er := fmt.Println(isTrue(false))
	if er != nil {
		log.Default()
	}
}

func isTrue(x bool) bool {
	return x
}
