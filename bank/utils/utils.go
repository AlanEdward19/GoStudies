package utils

import (
	"fmt"
)

func BuildMenu() int {
	fmt.Println(
		`What would you like to do?
1. Check balance
2. Deposit
3. Withdraw
4. Exit`)

	var choice int

	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	return choice
}
