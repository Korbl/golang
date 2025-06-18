package main

import (
	"log"
)

type Whoopsie struct{}

func (w Whoopsie) Error() string {
	return "big bada boom"
}

func main() {
	// One of the interfaces defined in the Go builtin types is the error interface.
	// Create a struct called Whoopsie that implements the interface and returns
	// an error message.
	// Pass the error to the function below.
	woops := Whoopsie{}
	logMyError(woops)
}

func logMyError(err error) {
	log.Fatalf("the program exited with an error: %s", err)
}
