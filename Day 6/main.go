package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func sum(a, b int) int {
	return a + b
}

func substract10(value *int) int {
	*value = *value - 5
	return *value - 10
}

func main() {
	fmt.Println("Starting", randomdata.City())

	x := 3
	y := 10

	// var num *int

	z := &x

	fmt.Println(x, y, &x, &y, *z, z)

	result := sum(x, y)
	fmt.Println(result)

	fmt.Println("********************************")
	var num int = 15
	result1 := substract10(&num)
	fmt.Println(result1)

	fmt.Println("********************************")
	var value = 25
	var num1 *int
	num1 = &value
	result2 := substract10(num1)
	fmt.Println(result2)

	fmt.Println("********************************")

	var u = 20
	result3 := substract10(&u)
	fmt.Println(result3)
	fmt.Println("value of u ", u)
}
