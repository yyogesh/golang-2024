package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	const inflationRate = 5
	investmentAmount, year, expectedReturnRate := 100000, 2, 7.2

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(year))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, float64(year))

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

	var name = "RAVI"
	fmt.Println(len(name))
	fmt.Println(name[2])
	fmt.Println(name[0:2])
	fmt.Println(name[:3])
	fmt.Println(name[1:])
	fmt.Println(name[:])

	var first = "testing"
	var second = first
	second = "second"
	fmt.Println(first)
	fmt.Println(second)

	// var word = first + " " + second
	fmt.Println(strings.HasPrefix("test", "te"))

	// var email = "test123@gmail.com"

	// var myArray [3]string

	// var arr1 = [3]string{"test", "te", "asdf"}

	// arr1[0] = "asf"

	// len(arr1)
}
