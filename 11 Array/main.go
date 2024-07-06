package main

import (
	"fmt"
	"reflect"
)

func main1() {
	// var a int = 10;
	var a [3]int
	fmt.Println(a)

	var c [3]int
	c[0] = 10
	c[1] = 20
	c[2] = 30
	fmt.Println("Value of c is: ", c)

	d := [3]int{12}
	fmt.Println("Value of d is: ", d)

	e := [...]int{12, 13, 14, 15, 16, 17}
	fmt.Println("Value of e is: ", e)

	f := [...]float64{67.7, 89.9, 21, 78}
	for i := 0; i < len(f); i++ {
		fmt.Printf("Value of h at index %d is: %.2f\n", i, f[i])
	}

	fmt.Println("--------------------------------")

	for index, value := range f {
		fmt.Printf("Value of h at index %d is: %.2f\n", index, value)
	}

	g := [5]int{1: 20, 3: 15}
	fmt.Println("Value of g is: ", g)

	h := [...]string{"USA", "China", "India", "Germany", "France"}
	i := h
	i[0] = "USA1"
	fmt.Println("Value of h is: ", h)
	fmt.Println("Value of i is: ", i)

	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)

	strArray1 := [3]string{"Japan", "Australia", "Germany"}
	strArray2 := strArray1  // data is passed by value
	strArray3 := &strArray1 // data is passed by Reference

	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)

	strArray1[0] = "Canada"
	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)
	fmt.Printf("strArray3: %v\n", *strArray3)

	fmt.Println(reflect.ValueOf(strArray1).Kind())

	// for _, value := range f {
	// }

	// z := 0
	// for range strArray1 {
	// 	fmt.Printf("strArray", strArray1[z])
	// 	z++
	// }

	strArray := [5]string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(itemExists(strArray, "Canada"))
	fmt.Println(itemExists(strArray, "Africa"))

	countries := [...]string{"India", "Canada", "Japan", "Germany", "Italy"}

	fmt.Printf("Countries: %v\n", countries)

	fmt.Printf(":2 %v\n", countries[:2])

	fmt.Printf("1:3 %v\n", countries[1:3])

	fmt.Printf("2: %v\n", countries[2:])

	fmt.Printf("2:5 %v\n", countries[2:5])

	fmt.Printf("0:3 %v\n", countries[0:3])

	fmt.Println(countries[len(countries)-1])

	fmt.Printf("All elements: %v\n", countries[0:len(countries)])

	fmt.Printf("2: %v\n", countries[:])
	fmt.Printf("2: %v\n", countries[0:])

	fmt.Printf("2: %v\n", countries[len(countries)-2:])

	a1 := [3][2]string{
		{"USA", "Washington"},
		{"India", "New Delhi"},
		{"Canada", "Ottawa"},
	}

	a1[0][0] = "Canada1"

	for _, v1 := range a1 {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}
