package main

import "fmt"

type Calc func(x, y int) int

func sum(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func executeCalc(c Calc, x int, y int) int {
	return c(x, y)
}

func main() {
	// given the custom type definition Calc
	// declare two variables with the Calc type
	// one function adds two integers and a second function subtracts two integers
	// declare a third function executeCalc which takes the Calc type as first argument
	// and the two integers as second an third argument.
	fmt.Println(executeCalc(sum, 21, 21))
	fmt.Println(executeCalc(subtract, 21, 21))
}
