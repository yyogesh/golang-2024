package main

import (
	"fmt"
	"log"
)

func main() {
	var accountBalance = 1000.0
	// while(true){}
	for {
		fmt.Println("Welcome to ABC Bank")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit")

		var choice int
		// fmt.Println("Enter your choice")
		log.SetPrefix("Bank :: ")
		log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
		log.Println("Enter your choice")
		fmt.Scan(&choice)

		fmt.Println("You choose:", choice)

		switch choice {
		case 1:
			{
				fmt.Println("You choose to deposit")
				var depositBalance float64
				fmt.Scan(&depositBalance)

				if depositBalance <= 0 {
					fmt.Println("Invalid withdraw. Please enter greater than zero")
					return
				}

				accountBalance = accountBalance + depositBalance
				fmt.Print("Amount Deposited! New Balance:", accountBalance)
			}
		case 2:
			{
				fmt.Println("You choose to withdraw")
				var withdrawBalance float64
				fmt.Scan(&withdrawBalance)

				if withdrawBalance <= 0 {
					fmt.Println("Invalid withdraw. Please enter greater than zero")
					return
				}

				if withdrawBalance > accountBalance {
					fmt.Println("Invalid withdraw. You can not withdraw more than you have.")
					return
				}

				accountBalance -= withdrawBalance
				fmt.Print("Amount Withdrawn! New Balance: ", accountBalance)
			}
		case 3:
			{
				fmt.Println("You choose to check balance: ", accountBalance)
			}
		default:
			{
				fmt.Println("You choose to exit")
				break
			}
		}
	}
}
