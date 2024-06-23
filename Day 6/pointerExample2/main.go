package main

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	var x int = 5
	var y int = 10

	fmt.Println("Before swap:")
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	swap(&x, &y)

	fmt.Println("After swap:")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}
