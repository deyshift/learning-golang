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
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Println("This is the Profit Calculator application.")
	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter cost of expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Printf("Earnings before tax: %.2f\n", ebt)
	fmt.Printf("Earnings after tax: %.2f\n", profit)
	fmt.Printf("Profit  Ratio: %.2f\n", ratio)

}
