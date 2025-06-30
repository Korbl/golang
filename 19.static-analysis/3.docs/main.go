// Linting requirement.
package main

import (
	"fmt"
	"log"

	"github.com/wimspaargaren/go-training-exercises/19.static-analysis/3.docs/mypkg"
)

// Run make lint and see if you can fix the lint errors.
func main() {
	_, err := fmt.Println(mypkg.HelloWorld())
	if err != nil {
		log.Default()
	}
}
