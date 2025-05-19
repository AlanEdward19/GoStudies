package main

import (
	"bank/finance"
	"bank/utils"
	"fmt"
)

func main() {
	accountBalance := 1000.0
	finance.SetBalance(accountBalance)

	fmt.Println("Welcome to the bank!")
	for {
		fmt.Println()
		choice := utils.BuildMenu()

		switch choice {
		case 1:
			finance.CheckBalance()
			break

		case 2:
			var depositAmount float64
			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println()
				fmt.Println("---------------------------")
				fmt.Println("Error: Deposit amount must be greater than zero.")
				fmt.Println("---------------------------")
				fmt.Println()
				break
			}

			finance.Deposit(depositAmount)
			break

		case 3:
			var withdrawAmount float64
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println()
				fmt.Println("---------------------------")
				fmt.Println("Error: Withdraw amount must be greater than zero.")
				fmt.Println("---------------------------")
				fmt.Println()
				break
			}

			finance.Withdraw(withdrawAmount)
			break

		case 4:
			fmt.Println("Thank you for using the bank. Goodbye!")
			return
		}
	}
}
