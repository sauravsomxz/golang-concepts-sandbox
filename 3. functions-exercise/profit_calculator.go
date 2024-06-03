package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserInput("Enter Your Revenue: ")
	expenses = getUserInput("Enter Your Expenses: ")
	taxRate = getUserInput("Enter Your TaxRate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func getUserInput(infoText string) (revenue float64) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profile, ratio float64) {

	ebt = revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}
