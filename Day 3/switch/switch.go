package main

import (
	"fmt"
	"time"
)

func main() {
	finger := 4
	fmt.Printf("Finger %d is ", finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	// case 6:
	case 5:
		fmt.Println("Pinky")
	default: //default case
		fmt.Println("incorrect finger number")
	}

	fmt.Println("****************************************************************")

	letter := "i"
	fmt.Printf("Letter %s is a ", letter)
	switch letter {
	case "a", "e", "i", "o", "u": //multiple expressions in case
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

	fmt.Println("****************************************************************")

	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Printf("%d is greater than 0 and less than 50", num)
	case num >= 51 && num <= 100:
		fmt.Printf("%d is greater than 51 and less than 100", num)
	case num >= 101:
		fmt.Printf("%d is greater than 100", num)
	}

	fmt.Println("****************************************************************")

	switch num1 := 75; {
	case num1 < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num1 < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num1 < 200:
		fmt.Printf("%d is lesser than 200", num)
	}

	fmt.Println("****************************************************************")

	switch num1 := 75; {
	case num1 < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num1 < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num1 == 200:
		fmt.Printf("%d is lesser than 200", num)
	}

	fmt.Println("****************************************************************")

	switch num1 := 75; {
	case num1 < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num1 < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num1 == 200:
		if num1 < 100 {
			break
		}
		fmt.Printf("%d is lesser than 200", num)
	}

	fmt.Println("****************************************************************")

	today := time.Now()
	fmt.Println(today)
	fmt.Println(today.Day())

}
