package main

type X struct {
	V int
}

func itsAlways42(x *X) {
	x.V = 42
}

func main() {
	// Given the struct X, create a function called itsAlways42 taking a pointer to X
	// and setting it's value V to 42
	// run the test to verify your result
}
