package main

import (
	"example.com/bank/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------")
		panic("could not read balance")
	}

	fmt.Println("welcome to the bank")
	fmt.Println("Reach Us 24/7 at:", randomdata.PhoneNumber())

	for {
		presentOptins()
		var choice int
		fmt.Scan(&choice)

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
			fileops.WriteBFloatToFile(accountBalance, accountBalanceFile)

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
			fileops.WriteBFloatToFile(accountBalance, accountBalanceFile)

		default:
			fmt.Println("Goodbye")
			fmt.Println("Thank you for visiting the bank")
			return

		}
	}

}
