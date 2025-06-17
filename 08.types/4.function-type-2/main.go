package main

type CalcType int

const (
	// Add a type for Subtract and Multiplication
	CalcTypeSum CalcType = iota // iota: https://go.dev/wiki/Iota
	CalcTypeSubtract
	CalcTypeMultiplication
)

type Calc func(x, y int) int

func sum(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiplication(x, y int) int {
	return x * y
}

func main() {
	// Call the GetCalculationFunction with the different CalcTypes and execute them with the following variables
	_, _ = 42, 42
	// Execute the test to verify  your solution
}

func GetCalculationFunction(calcType CalcType) Calc {
	// Implement the GetCalculationFunction
	// Hint: Use a switch statement to return the correct function
	switch calcType {
	case CalcTypeSum:
		return sum
	case CalcTypeSubtract:
		return subtract
	case CalcTypeMultiplication:
		return multiplication
	default:
		return nil
	}
}
