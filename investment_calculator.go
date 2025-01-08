package main // designates this file as the main entry point of the program

import (
	"fmt"
	"math"
)

// imports the fmt package for formatted I/O

// main entry point of the program, execution starts here
func main() {
	var investmentAmount float64 = 1000 // use var keyword to declare a variable and specify the type
	years := 10.0                       // use assignment operator to infer type and initialize the variable
	expectedReturnRate := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Printf("The future value of the investment is: %.2f\n", futureValue)
}
