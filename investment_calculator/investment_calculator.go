package main // designates this file as the main entry point of the program

import (
	"fmt"
	"math"
)

const inflationRate = 2.5 // use const keyword to declare a constant

// main entry point of the program, execution starts here
func main() {
	var investmentAmount float64 // use var keyword to declare a variable and specify the type when not initializing
	var years float64
	expectedReturnRate := 5.5 // use assignment operator to infer type and initialize the variable - kept this for example

	fmt.Print("Enter the investment amount: ") // use fmt.Println to print a message to the console
	fmt.Scan(&investmentAmount)                // use fmt.Scan to read input from the user

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Printf("The future value of the investment is: %.2f\n", futureValue)
	fmt.Printf("The future real value of the investment is: %.2f\n", futureRealValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	// return investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
