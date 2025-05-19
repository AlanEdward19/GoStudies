package finance

import "fmt"

var balance float64

func SetBalance(value float64) {
	balance = value
}

func CheckBalance() {
	fmt.Printf("Your current balance is: $%.2f\n", balance)
}

func Deposit(value float64) {
	balance += value
	fmt.Printf("You have deposited: $%.2f\n", value)
	fmt.Printf("Your new balance is: $%.2f\n", balance)
}

func Withdraw(value float64) {
	if value > balance {
		fmt.Println()
		fmt.Println("---------------------------")
		fmt.Println("Error: Insufficient funds!")
		fmt.Println("---------------------------")
		fmt.Println()
	} else {
		balance -= value
		fmt.Printf("You have withdrawn: $%.2f\n", value)
		fmt.Printf("Your new balance is: $%.2f\n", balance)
	}
}
