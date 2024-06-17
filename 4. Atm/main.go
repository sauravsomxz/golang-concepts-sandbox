package main

import "fmt"

func main() {

	var accountBalance = 1000

	fmt.Println("Welcome to Bank")
	fmt.Println("1. Check Account Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4 Exit")

	var choice int64
	fmt.Print("Choose Option: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Account Balance: ", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your Deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount > 0 {
			accountBalance = accountBalance + int(depositAmount)

			fmt.Println("Updated Account Balance: ", accountBalance)
		} else {
			fmt.Println("Deposit Amount cannot be negative.")
		}

	} else if choice == 3 {
		fmt.Print("Your Withdrawl: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		if withdrawAmount > 0 && withdrawAmount < float64(accountBalance) {
			accountBalance = accountBalance - int(withdrawAmount)

			fmt.Println("Updated Account Balance: ", accountBalance)
		} else {
			if withdrawAmount <= 0 {
				fmt.Println("Invalid Amount. Withdraw Amount is invalid")
			} else {
				fmt.Println("Withdrawl Amount is greater than available amount")
			}
		}

	} else {
		fmt.Println("Thank you for using our bank.")
	}
}
