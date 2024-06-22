package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const balanceFile = "balance.txt"

const (
	// YYYY-MM-DD: 2022-03-23
	YYYYMMDD = "2006-01-02"
	// 24h hh:mm:ss: 14:23:20
	HHMMSS24h = "15:04:05"
	// 12h hh:mm:ss: 2:23:20 PM
	HHMMSS12h = "3:04:05 PM"
	// text date: March 23, 2022
	TextDate = "January 2, 2006"
	// text date with weekday: Wednesday, March 23, 2022
	TextDateWithWeekday = "Monday, January 2, 2006"
	// abbreviated text date: Mar 23 Wed
	AbbrTextDate = "Jan 2 Mon"
)

func writeValueInFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func readDataFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return defaultValue, errors.New("Failed to read data from file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		fmt.Println(err)
		return defaultValue, errors.New("Failed to parse float value: " + valueText)
	}
	return value, nil
}

func main() {
	// Defines flags
	flags := log.Lshortfile
	// Define date and time format
	datetime := time.Now().UTC().Format(YYYYMMDD+" "+HHMMSS12h) + ": "

	// Open a file if it exist or create one if file does not exist
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// Set file as the destination to log messages with log level "FATAL"
		logger := log.New(file, "", flags)
		logger.SetPrefix("FATAL: " + datetime)
		logger.Println(err)
	}
	defer file.Close()
	logger := log.New(file, "", flags)
	logger.SetPrefix("Bank : " + datetime)

	mw := io.MultiWriter(os.Stdout, file)
	logger.SetOutput(mw)

	var accountBalance, fileError = readDataFromFile(balanceFile, 1000)
	if fileError != nil {
		logger.Println("Error reading")
		logger.Println(fileError)
		logger.Println("--------------------------------")
		return
	}
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
		// log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
		logger.Println("Enter your choice")
		fmt.Scan(&choice)

		logger.Println("You choose:", choice)

		switch choice {
		case 1:
			{
				logger.Println("You choose to deposit")
				var depositBalance float64
				fmt.Scan(&depositBalance)

				if depositBalance <= 0 {
					logger.Println("Invalid withdraw. Please enter greater than zero")
					return
				}

				accountBalance = accountBalance + depositBalance
				writeValueInFile(balanceFile, accountBalance)
				logger.Print("Amount Deposited! New Balance:", accountBalance)
			}
		case 2:
			{
				logger.Println("You choose to withdraw")
				var withdrawBalance float64
				fmt.Scan(&withdrawBalance)

				if withdrawBalance <= 0 {
					logger.Println("Invalid withdraw. Please enter greater than zero")
					return
				}

				if withdrawBalance > accountBalance {
					logger.Println("Invalid withdraw. You can not withdraw more than you have.")
					return
				}

				accountBalance -= withdrawBalance
				writeValueInFile(balanceFile, accountBalance)
				logger.Print("Amount Withdrawn! New Balance: ", accountBalance)
			}
		case 3:
			{
				logger.Println("You choose to check balance: ", accountBalance)
			}
		default:
			{
				logger.Println("You choose to exit")
				break
			}
		}
	}
}
