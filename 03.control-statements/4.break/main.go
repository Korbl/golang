package main

import "fmt"

func main() {
	x := 0
	// let the program print the number 0 - 10 without adjusting the next line.
	for {
		if x >= 10 {
			break
		}
		x++
		fmt.Printf("%d", x)

	}
}
