package main

/*
This is the Profit Calculator application.
this calculator will ask for the following inputs:
1. Revenue
2. Expenses
3. Taxes

It will calculate:
1. Earnings before tax (EBT)
2. Earnings after tax (profit)
3. Ratio (EBT/profit)
4. Output EBT, profit, ratio
*/

import "fmt"

func main() {

	fmt.Println("This is the Profit Calculator application.")

	revenue := getUserInput("Enter revenue: ")
	expenses := getUserInput("Enter cost of expenses: ")
	taxRate := getUserInput("Enter tax rate: ")

	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %.2f\n", ebt)
	fmt.Printf("Earnings after tax: %.2f\n", profit)
	fmt.Printf("Profit  Ratio: %.2f\n", ratio)

}

func getUserInput(infoText string) float64 {
	// Function to get user input
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateProfit(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}
