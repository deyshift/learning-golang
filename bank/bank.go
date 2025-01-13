/*
The bank application provides functionality for managing bank accounts.

It provides the following features:

1. Check balance
2. Deposit money
3. Withdraw money
4. Exit
*/

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var accountBalance, err = readBalanceFromFile()
var running = true

func main() {

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------------")
	}

	fmt.Println("Welcome to the Go bank application!")

	for running {
		displayMenu()
		choice := getUserChoice()
		handleChoice(choice)
	}
}

func displayMenu() {
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func getUserChoice() int {
	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	return choice
}

func handleChoice(choice int) {
	switch choice {
	case 1: // Check balance
		checkBalance()
	case 2: // Deposit money
		depositMoney()
	case 3: // Withdraw money
		withdrawMoney()
	case 4: // Exit
		exit()
	default:
		fmt.Println("Invalid choice")
	}
}

func checkBalance() {
	fmt.Println("Your account balance is $", accountBalance)
}

func depositMoney() {
	fmt.Println("Enter the amount you want to deposit:")
	var depositAmount float64
	fmt.Scan(&depositAmount)

	if depositAmount <= 0 {
		fmt.Println("Invalid deposit amount")
		return
	}

	fmt.Printf("You have deposited $%.2f\n", depositAmount)
	accountBalance += depositAmount
	fmt.Println("Your new account balance is $", accountBalance)
	writeBalanceToFile(accountBalance)
}

func withdrawMoney() {
	fmt.Println("Enter the amount you want to withdraw:")
	var withdrawAmount float64
	fmt.Scan(&withdrawAmount)

	if withdrawAmount <= 0 {
		fmt.Println("Invalid withdraw amount")
		return
	} else if withdrawAmount > accountBalance {
		fmt.Println("Insufficient funds")
		return
	}

	fmt.Printf("You have withdrawn $%.2f\n", withdrawAmount)
	accountBalance -= withdrawAmount
	fmt.Println("Your new account balance is $", accountBalance)
	writeBalanceToFile(accountBalance)
}

func exit() {
	fmt.Println("Thank you for using the bank application!")
	running = false
}

func readBalanceFromFile() (float64, error) {
	// Read the account balance from a file
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 1000, errors.New("Failed to read balance from file")
	}

	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 1000, errors.New("Failed to read balance from file")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	// Convert the balance to a string
	balanceStr := fmt.Sprintf("%.2f", balance)

	// Write the account balance to a file
	err := os.WriteFile("balance.txt", []byte(balanceStr), 0644)
	if err != nil {
		fmt.Println("Error writing balance to file:", err)
		return
	}

	fmt.Println("Balance successfully written to file.")
}
