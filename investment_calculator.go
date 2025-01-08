package main // designates this file as the main entry point of the program

import (
	"fmt"
	"math"
)

// main entry point of the program, execution starts here
func main() {
	const inflationRate = 2.5    // use const keyword to declare a constant
	var investmentAmount float64 // use var keyword to declare a variable and specify the type
	years := 10.0                // use assignment operator to infer type and initialize the variable
	expectedReturnRate := 5.5

	fmt.Print("Enter the investment amount: ") // use fmt.Println to print a message to the console
	fmt.Scan(&investmentAmount)                // use fmt.Scan to read input from the user

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("The future value of the investment is: %.2f\n", futureValue)
	fmt.Printf("The future real value of the investment is: %.2f\n", futureRealValue)
}
