package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("welcome to the bank")
	fmt.Println("How can I help you?")
	fmt.Println("1. check balance")
	fmt.Println("2. deposit")
	fmt.Println("3. withdraw")
	fmt.Println("4. exit")

	var choice int
	fmt.Scan(&choice)

	//wantsCheckBalance := choice == 1

	if choice == 1 {
		fmt.Println("Your balance is: ", accountBalance)
	} else if choice == 2 {
		fmt.Print("Enter the amount you want to deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Your updated balance is: ", accountBalance)
	} else if choice == 3 {
		fmt.Print("Enter the amount you want to withdraw: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		accountBalance -= withdrawAmount
		fmt.Println("Your updated balance is: ", accountBalance)
	} else {
		fmt.Println("Thank you for visiting the bank")
	}
}
