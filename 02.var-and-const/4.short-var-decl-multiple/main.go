package main

import (
	"fmt"
)

func main() {
	// Declare a string and boolean variable in a single line
	line, test := "this is a string", false
	// and print both variables. Use the short variable declaration.
	fmt.Printf("%s and then it is %t\n", line, test)
	// After the first print, change the content of the variables and
	line, test = "new string", true
	fmt.Printf("%s and then it is %t\n", line, test)
	// print them again.
}
