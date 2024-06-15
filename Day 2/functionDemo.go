package main

import "fmt"

// func functionName(param) returnType {
//})

func calculateBill(price int, no int) int {
	return price * no
}

// func rectProps(length, width float64) (float64, float64) {
// 	var area = length * width
// 	var parameter = (length + width) * 2
// 	return area, parameter
// }

// Named return values
func rectProps(length, width float64) (area, parameter float64) {
	area = length * width
	parameter = (length + width) * 2
	return // area, parameter
}

// var rect = function() {}

var (
	area1 = func(l, b int) int {
		return l * b
	}
)

func sum(x, y int) int {
	return x + y
}

func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}

func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

// redux connect()()

func main() {
	// partialFun := partialSum(3)
	// fmt.Println(partialFun(10))
	fmt.Println(partialSum(3)(10))

	// squareSum(5)(6)(7)
	fmt.Println("Hello World")
	totalPrice := calculateBill(10, 2)
	fmt.Println(totalPrice)
	// area, parameter := rectProps(10.8, 5.6)
	// fmt.Printf("Area: %f Parameter: %f\n", area, parameter)

	// Blank Identifer
	area, _ := rectProps(10.8, 5.6)
	fmt.Printf("Area: %f", area)
	fmt.Println(area1(20, 30))

	// iife
	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 2)

	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)

	// var (
	// 	name = "sad"
	// )

	// Closures

	l := 20
	b := 30

	func() {
		fmt.Println(l * b)
	}()
}
