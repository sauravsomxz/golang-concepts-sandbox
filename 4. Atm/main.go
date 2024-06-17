package main

import (
	"fmt"
	"os"
	"strconv"
)

var balanceText = "balanceOfUser.txt"

func main() {
	var accountBalance = readBalanceFromFile()

	fmt.Println("Welcome to Bank")
	fmt.Println("1. Check Account Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4 Exit")

	var choice int64
	fmt.Print("Choose Option: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Account Balance: ", accountBalance)
	case 2:
		fmt.Print("Your Deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount > 0 {
			accountBalance = accountBalance + depositAmount

			fmt.Println("Updated Account Balance: ", accountBalance)
			writeBalanceToFile(float64(accountBalance))
		} else {
			fmt.Println("Deposit Amount cannot be negative.")
		}
	case 3:
		fmt.Print("Your Withdrawl: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		if withdrawAmount > 0 && withdrawAmount < float64(accountBalance) {
			accountBalance = accountBalance - withdrawAmount

			fmt.Println("Updated Account Balance: ", accountBalance)
			writeBalanceToFile(float64(accountBalance))
		} else {
			if withdrawAmount <= 0 {
				fmt.Println("Invalid Amount. Withdraw Amount is invalid")
			} else {
				fmt.Println("Withdrawl Amount is greater than available amount")
			}
		}
	default:
		fmt.Println("Thank you for using our bank.")
	}

	// if choice == 1 {
	// 	fmt.Println("Account Balance: ", accountBalance)
	// } else if choice == 2 {
	// 	fmt.Print("Your Deposit: ")
	// 	var depositAmount float64
	// 	fmt.Scan(&depositAmount)

	// 	if depositAmount > 0 {
	// 		accountBalance = accountBalance + int(depositAmount)

	// 		fmt.Println("Updated Account Balance: ", accountBalance)
	// 	} else {
	// 		fmt.Println("Deposit Amount cannot be negative.")
	// 	}

	// } else if choice == 3 {
	// 	fmt.Print("Your Withdrawl: ")
	// 	var withdrawAmount float64
	// 	fmt.Scan(&withdrawAmount)

	// 	if withdrawAmount > 0 && withdrawAmount < float64(accountBalance) {
	// 		accountBalance = accountBalance - int(withdrawAmount)

	// 		fmt.Println("Updated Account Balance: ", accountBalance)
	// 	} else {
	// 		if withdrawAmount <= 0 {
	// 			fmt.Println("Invalid Amount. Withdraw Amount is invalid")
	// 		} else {
	// 			fmt.Println("Withdrawl Amount is greater than available amount")
	// 		}
	// 	}

	// } else {
	// 	fmt.Println("Thank you for using our bank.")
	// }

	//// Looping in GO
	for i := 0; i < 2; i++ {
		fmt.Println("GO LOOP PRINTING")
	}

	//// Infinte Loop
	// for {
	// 	fmt.Println("HI!")
	// }
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balanceOfUser.txt", []byte(balanceText), 0644)
}

func readBalanceFromFile() float64 {
	data, _ := os.ReadFile(balanceText)

	balanceText := string(data)

	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance

}
