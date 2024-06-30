package main

import "fmt"

// a: int | string

// Employee<T, U>

// Employee<int> employee = new Employee()

func main() {
	fmt.Println("Hello World")
	response := sum(4, 5)
	fmt.Println("Sum:", response)

	response1 := sum("asdf", "asdf")
	fmt.Println("Sum:", response1)
}

func sum[T int | float64 | string](a, b T) T {
	return a + b
}

// Collection // Array slices and Maps

// Project

// Concurrency

// Project

// Rest API

// Project
