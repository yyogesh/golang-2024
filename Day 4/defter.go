package main

import "fmt"

func displayValue(a int) {
	fmt.Println("value of a in deferred function", a)
}

func calc(a, b int) int {
	res := a * b
	fmt.Println("Product: ", res)
	return 0
}

func show() {
	fmt.Println("The product of two numbers.")
}

func main1() {
	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")
	defer fmt.Println("Statement 3")
	// LIFO

	fmt.Println("Statement 4")

	a := 5
	defer displayValue(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)

	defer calc(23, 56)
	show()
}
