package finance

import "fmt"

type Account struct {
	balance float64
}

func NewAccount(initialBalance float64) *Account {
	return &Account{balance: initialBalance}
}

func (account *Account) SetBalance(value float64) {
	account.balance = value
}

func (account *Account) CheckBalance() {
	fmt.Printf("Your current balance is: $%.2f\n", account.balance)
}

func (account *Account) Deposit(value float64) {
	account.balance += value
	fmt.Printf("You have deposited: $%.2f\n", value)
	account.CheckBalance()
}

func (account *Account) Withdraw(value float64) {
	if value > account.balance {
		fmt.Println()
		fmt.Println("---------------------------")
		fmt.Println("Error: Insufficient funds!")
		fmt.Println("---------------------------")
		fmt.Println()
	} else {
		account.balance -= value
		fmt.Printf("You have withdrawn: $%.2f\n", value)
		account.CheckBalance()
	}
}
