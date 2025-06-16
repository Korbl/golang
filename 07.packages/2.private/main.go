package main

import "github.com/wimspaargaren/go-training-exercises/07.packages/2.private/mypkg"

func main() {
	// Retrieve the constant from mypkg, without making the constant itself public. Hint: Use a function.
	hiddenConst := mypkg.ReturnConstant()
	println(hiddenConst)
}
