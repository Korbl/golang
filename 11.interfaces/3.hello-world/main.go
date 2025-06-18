package main

import "fmt"

type Stringer interface {
	String() string
}

type stringer struct{}

func (s stringer) String() string {
	return fmt.Sprintln("this is a string")
}

func main() {
	s := stringer{}
	// Implement the stringer interface and pass it to the log function.
	log(s)
}

func log(s Stringer) {
	println(s.String())
}
