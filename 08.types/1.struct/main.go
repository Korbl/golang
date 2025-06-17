package main

import "fmt"

type train struct {
	numberPassengers int
	velocity         int
	name             string
}

func main() {
	// Create a struct representing a train which holds 3 fields: The amount of passengers, the velocity and a name.
	// The main function should initialise the struct with some values and print the result of the struct.
	// Hint: try using the `%+v` verb to print the struct value.
	//
	newTrain := train{
		numberPassengers: 50,
		velocity:         100,
		name:             "intercity",
	}
	fmt.Printf("train: %+v", newTrain)
}
