package main

import "fmt"

type Currency string

const (
	EUR Currency = "€"
	USD Currency = "$"
	GBP Currency = "£"
)

func currency(input Currency, amount int) {
	switch input {
	case "EUR":
		fmt.Printf("€%d", amount)

	case "USD":
		fmt.Printf("$%d", amount)
	case "GBP":
		fmt.Printf("£%d", amount)

	default:
		fmt.Println(input)
	}
}

func main() {
	// Create a custom "Currency" type which has a string as underlying type. Define three constant currency values: EUR, USD, GBP.
	// Write a function that takes the Currency type as argument and tries to detect which currency we're dealing with(hint: use a switch case).
	// If we find a known currency we will format it with the corresponding currency sign (€, $, £).
	currency("EUR", 500)
}
