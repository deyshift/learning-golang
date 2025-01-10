package main // designates this file as the main entry point of the program

import (
	"fmt"
	"math"
)

const inflationRate = 2.5 // use const keyword to declare a constant

// main entry point of the program, execution starts here
func main() {
	investmentAmount := getUserInput("Enter the investment amount: ")
	years := getUserInput("Enter the number of years: ")
	expectedReturnRate := getUserInput("Enter the expected return rate: ")

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Printf("The future value of the investment is: %.2f\n", futureValue)
	fmt.Printf("The future real value of the investment is: %.2f\n", futureRealValue)
}

// calculateFutureValues calculates the future value and real future value of an investment
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
