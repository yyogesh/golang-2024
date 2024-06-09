package main

import "fmt"

// int string float32, int8  int16  int32 int65
// byte
// string
// bool
// complex64

func main1() {
	var n int
	fmt.Print("Enter the temperature in Celsius: ")
	fmt.Scanf("%d", &n)
	var f = (float32(n) * 1.8) + 32
	fmt.Println("Temperature in Fahrenheit: ", f)
}

// func main() {
// 	var ageInYear int
// 	ageInYear = 20
// 	fmt.Println(`Hello World "asdf" 'hi' %d`, ageInYear)
// }

// GOOS=linux GOARCH=amd64 go build demo.go
