package main

func addPointer(x *int, y int) {
	*x = *x + y
}

func main() {
	// Let's recreate our add function and call it addPointer, but now instead of returning the sum of two integers
	// let the function take the first argument as pointer and add the second argument as
	// value to the first one.
	// Run the test to verify your solution
}
