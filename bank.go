package main

import "fmt"
import "os"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	var accountBalance float64 = 1000
	fmt.Println("welcome to the bank")
	for {
		fmt.Println("How can I help you?")
		fmt.Println("1. check balance")
		fmt.Println("2. deposit")
		fmt.Println("3. withdraw")
		fmt.Println("4. exit")

		var choice int
		fmt.Scan(&choice)

		//wantsCheckBalance := choice == 1
		switch choice {

		case 1:
			fmt.Println("Your balance is: ", accountBalance)
			continue

		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your new balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)

		case 3:
			fmt.Print("Enter the amount you want to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Your new balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)

		default:
			fmt.Println("Goodbye")
			fmt.Println("Thank you for visiting the bank")
			return

		}
	}
}
