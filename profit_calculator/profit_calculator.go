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

/*
Second iteration:
1. Adds the ability to store calculated results to file.
2. Validates user input:
	- Show error messages if invalid input is provided
	- No negative numbers
	- No zero values
*/

import (
	"errors"
	"fmt"
	"os"
)

var running = true

func main() {

	fmt.Println("This is the Profit Calculator application.")

	for running {
		choiceMenu()
		choice := getUserChoice()
		handleChoice(choice)
	}
}

func choiceMenu() {
	fmt.Println("What would you like to do?")
	fmt.Println("1. Calculate profit")
	fmt.Println("2. View stored calculations")
	fmt.Println("3. Exit")
}

func getUserChoice() int {
	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	return choice
}

func handleChoice(choice int) {
	switch choice {
	case 1: // Calculate profit
		calculateAndStoreProfit()
	case 2: // View stored calculations
		viewStoredCalculations()
	case 3: // Exit
		running = false
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}

func getUserInput(infoText string) (float64, error) {
	// Function to get user input
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("you must enter a positive number")
	}

	return userInput, nil
}

func calculateAndStoreProfit() {
	revenue, err := getUserInput("Enter revenue: ")
	if err != nil {
		printError(err)
		return
	}
	expenses, err := getUserInput("Enter cost of expenses: ")
	if err != nil {
		printError(err)
		return
	}
	taxRate, err := getUserInput("Enter tax rate: ")
	if err != nil {
		printError(err)
		return
	}

	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	fmt.Printf("\nEarnings before tax: %.2f\n", ebt)
	fmt.Printf("Earnings after tax: %.2f\n", profit)
	fmt.Printf("Profit Ratio: %.2f\n", ratio)

	message, err := writeResultsToFile(ebt, profit, ratio)
	if err != nil {
		printError(err)
	} else {
		fmt.Println(message)
	}
}

func calculateProfit(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func writeResultsToFile(ebt, profit, ratio float64) (string, error) {
	// convert data to string
	ebtStr := fmt.Sprintf("%.2f", ebt)
	profitStr := fmt.Sprintf("%.2f", profit)
	ratioStr := fmt.Sprintf("%.2f", ratio)

	data := fmt.Sprintf("EBT: %s\nProfit: %s\nRatio: %s\n", ebtStr, profitStr, ratioStr)

	// write data to file
	err := os.WriteFile("profit.txt", []byte(data), 0644)
	if err != nil {
		return "", errors.New("error writing balance to file")
	}

	return "Results stored to file", nil
}

func viewStoredCalculations() {
	data, err := os.ReadFile("profit.txt")
	if err != nil {
		fmt.Println()
		fmt.Println("Error reading file:", err)
		fmt.Println()
		return
	}

	fmt.Printf("\nStored Calculations:\n")
	fmt.Println(string(data))
}

func printError(err error) {
	fmt.Printf("\n%s\n\n", err)
}
