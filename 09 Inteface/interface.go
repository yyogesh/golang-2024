package main

import "fmt"

// Define the MathOperation interface
type MathOperation interface {
	Square() float64
}

// Define a custom type MyFloat
type MyFloat float64

// Implement the Square method for MyFloat
func (m MyFloat) Square() float64 {
	return float64(m * m)
}

func CalculateSquare(m MathOperation) {
	fmt.Printf("The square is: %.2f\n", m.Square())
}

func main2() {
	var number MyFloat = 3.14
	// Calculate the square of the number
	CalculateSquare(number)
}
