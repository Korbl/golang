// more package comments
package main

import (
	"errors"
	"log"
)

var errHappened = errors.New("what could go wrong?")

// Run make lint and see if you can fix the lint errors.
func main() {
	err := whatCouldGoWrong()
	if err != nil {
		log.Default()
	}
}

func whatCouldGoWrong() error {
	return errHappened
}
