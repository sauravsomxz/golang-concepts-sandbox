// package main

// import "fmt"

// func main() {
// 	var revenue float64
// 	var expenses float64
// 	var taxRate float64

// 	fmt.Print("Enter You Revenue: ")
// 	fmt.Scan(&revenue)

// 	fmt.Print("Enter Your Expenses: ")
// 	fmt.Scan(&expenses)

// 	fmt.Print("Enter Your Tax Rate: ")
// 	fmt.Scan(&taxRate)

// 	fmt.Print("Profile before Tax: ")
// 	fmt.Println(revenue - expenses)

// 	var profitAfterTax = revenue * (1 - taxRate/100)
// 	fmt.Print("Profit after Tax: ")
// 	fmt.Println(profitAfterTax)

// 	fmt.Print("Ratio: ")
// 	fmt.Println(profitAfterTax / (revenue - expenses))

// 	fmt.Println("EXPENSES -> ", expenses)
// 	fmt.Println("EXPENSES -> ", expenses)
// }

package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter your Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter your Tax Rate (as a percentage): ")
	fmt.Scan(&taxRate)

	// Calculate profit before tax
	profitBeforeTax := revenue - expenses
	fmt.Printf("Profit before Tax: %.2f\n", profitBeforeTax)

	// Calculate profit after tax
	profitAfterTax := profitBeforeTax * (1 - taxRate/100)
	fmt.Printf("Profit after Tax: %.2f\n", profitAfterTax)

	// Calculate and print ratio if profit before tax is not zero
	if profitBeforeTax != 0 {
		ratio := profitAfterTax / profitBeforeTax
		fmt.Printf("Ratio: %.2f\n", ratio)
	} else {
		fmt.Println("Ratio: N/A (Profit before tax is zero)")
	}

	// Print expenses
	fmt.Printf("Expenses: %.2f\n", expenses)
}
