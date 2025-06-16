package main

import (
	"fmt"
	"strings"
)

func printTriangle(a int) {
	for x := 1; x <= a; x++ {
		fmt.Println(strings.Repeat("*", x))
	}
}

func main() {
	/**
	Create a function which takes an integer as argument and prints a triangle with
	a height and with of the given integer.

	Output for x = 10:
	```
	*
	* *
	* * *
	* * * *
	* * * * *
	* * * * * *
	* * * * * * *
	* * * * * * * *
	* * * * * * * * *
	```
	*/
	printTriangle(10)
}
