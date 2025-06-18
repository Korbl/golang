package main

import "fmt"

type emptyStruct struct{}

func checkType(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("It's an integer with value %d\n", v)
	case string:
		fmt.Printf("It's a string with value '%s'\n", v)
	case bool:
		fmt.Printf("It's a boolean with value %t\n", v)
	default:
		fmt.Printf("It's some other type: %T\n", v)
	}
}

func emptyFunc() {
}

func main() {
	// Create a function which takes an empty interface. The function should print the type of the variable passed and the value.
	// Pass the following types:
	// - An empty struct
	// - A pointer to an empty struct
	// - An integer
	// - A function
	// - The function itself
	// - A string
	newStruct := emptyStruct{}

	checkType(newStruct)
	checkType(&newStruct)
	checkType(54)
	checkType(emptyFunc)
	checkType("foo")
}
