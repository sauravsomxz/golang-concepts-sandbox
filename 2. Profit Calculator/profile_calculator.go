package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter You Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Your Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter Your Tax Rate: ")
	fmt.Scan(&taxRate)

	fmt.Print("Profile before Tax: ")
	fmt.Println(revenue - expenses)

	var profitAfterTax = revenue * (1 - taxRate/100)
	fmt.Print("Profit after Tax: ")
	fmt.Println(profitAfterTax)

	fmt.Print("Ratio: ")
	fmt.Println(profitAfterTax / (revenue - expenses))
}
