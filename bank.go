package main

import "fmt"

func main() {
	accountBalance := 1000.0

	fmt.Println("Welcome to Go Bank!!")
	fmt.Println("What transaction would you like to make?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is", accountBalance)
	}
}