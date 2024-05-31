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
	} else if choice == 2 {
		var depositAmount float64
		fmt.Print("Your deposit amount is ")
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Your updated balance is", accountBalance)
	} else if choice == 3 {
		var withdrawnAmount float64
		fmt.Print("Your amount to withdraw is ")
		fmt.Scan((&withdrawnAmount))
		accountBalance -= withdrawnAmount
		fmt.Println("Your updated balance is", accountBalance)
	} else {
		fmt.Println(("Thank you for using our app!"))
	}
}
