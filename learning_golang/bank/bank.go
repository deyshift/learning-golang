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
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/rivtechprojects/learning-golang/file_handling"
)

const accountBalanceFile = "balance.txt"

var accountBalance, err = file_handling.ReadDataFromFile(accountBalanceFile)
var running = true

func main() {

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------------")
	}

	fmt.Println("Welcome to the Go bank application!")
	fmt.Println("Your account number is:", randomdata.RandStringRunes(10))

	for running {
		displayMenu()
		choice := getUserChoice()
		handleChoice(choice)
	}
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
	file_handling.WriteDataToFile(accountBalance, accountBalanceFile)
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
	file_handling.WriteDataToFile(accountBalance, accountBalanceFile)
}

func exit() {
	fmt.Println("Thank you for using the bank application!")
	running = false
}
