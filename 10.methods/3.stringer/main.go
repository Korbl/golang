package main

import "fmt"

type train struct {
	name       string
	velocity   float64
	passengers int
}

func (t *train) String() string {
	return fmt.Sprintf("Train %s is going %v kph with %d passengers ", t.name, t.velocity, t.passengers)
}

func (t *train) GoString() string {
	return fmt.Sprintf("Train %s is going %v kph with %d passengers ", t.name, t.velocity, t.passengers)
}

func main() {
	// Recall the train struct from the types exercises. Define a
	// method on the train with the signature: String() string
	// Add some nice formatting to represent the name, speed and amount of passengers
	// and see what happens when we pass the struct to fmt.Println.
	// Bonus: Define a method GoString() string and pass it to fmt.Printf with the %#v verb.
	newTrain := train{
		name:       "choochoo",
		velocity:   80.05,
		passengers: 300,
	}
	fmt.Println(newTrain.GoString())
}
