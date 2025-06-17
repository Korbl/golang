package main

type counter struct {
	val int
}

func (c *counter) Add(input int) {
	c.val = c.val + input
}

func main() {
	// Define a method on counter called Add taking an integer value
	// and adding this value to the val field in the counter struct.
	// Afterwards define a loop which adds all values from 1 to 10 to the
	// counter using the Add method.
	newCounter := counter{}
	for x := 1; x <= 10; x++ {
		newCounter.Add(x)
	}
}
