package main

import (
	"fmt"

	"github.com/wimspaargaren/go-training-exercises/07.packages/1.addition/calc"
)

func Add(x, y int) int {
	return x + y
}

func main() {
	// Re-create the Add function, but now move it to a separate package as defined by the import statement above.
	fmt.Println(calc.Add(21, 21))
}
