package main // designates this file as the main entry point of the program

import (
	"fmt"
	"math"
)

// imports the fmt package for formatted I/O

// main entry point of the program, execution starts here
func main() {
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Printf("The future value of the investment is: %.2f\n", futureValue)
}
