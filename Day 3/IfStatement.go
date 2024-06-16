package main

import (
	"fmt"
)

func main1() {
	num := 10
	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	var mark = 100

	if mark >= 90 {
		fmt.Println("A")
	} else if mark < 90 && mark > 70 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	if num%2 == 0 {
		fmt.Println("even")
		if mark == 0 {
			fmt.Println("zero")
		}
	} else {
		fmt.Println("odd")
	}

}
