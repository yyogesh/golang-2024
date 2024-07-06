package main

import (
	"fmt"
	"reflect"
)

func main() {
	var intSlice []int
	var intArray [3]int
	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Println(reflect.ValueOf(intArray).Kind())
	a := [5]int{1, 2, 3, 4, 5}
	b := a[1:3]
	fmt.Println("********************************")
	fmt.Println(b)
	fmt.Println(reflect.ValueOf(a).Kind())
	fmt.Println(reflect.ValueOf(b).Kind())

	var intSlice1 = make([]int, 5)
	var strSlice = make([]string, 10, 20)

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice1), cap(intSlice1))
	fmt.Println(reflect.ValueOf(intSlice1).Kind())

	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
	fmt.Println(reflect.ValueOf(strSlice).Kind())

	// [0, 1, 2]

	b1 := []int{10, 20, 30, 40, 50}
	fmt.Printf("b1 \tLen: %v \tCap: %v\n", len(b1), cap(b1))

	var c1 = new([50]int)[0:10]
	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Printf("c1 \tLen: %v \tCap: %v\n", len(c1), cap(c1))

	d1 := make([]int, 2, 5)
	d1[0] = 10
	d1[1] = 20
	// d1[2] = 30
	fmt.Println("Slice A:", d1)
	fmt.Printf("Length is %d Capacity is %d\n", len(d1), cap(d1))

	p1 := append(d1, 30, 40, 50, 60, 70, 80)
	fmt.Println("Slice A after appending data:", p1)
	fmt.Printf("Length is %d Capacity is %d\n", len(d1), cap(p1))

	fmt.Println("Slice A:", d1)
	fmt.Printf("Length is %d Capacity is %d\n", len(d1), cap(d1))

	// fmt.Println()
	strSlice = RemoveIndex(strSlice, 3)
	fmt.Println(strSlice)

	a3 := []int{5, 6, 7} // Create a smaller slice
	fmt.Printf("[Slice:A] Length is %d Capacity is %d\n", len(a3), cap(a3))

	b3 := make([]int, 5, 10) // Create a bigger slice
	copy(b3, a3)             // Copy function
	fmt.Printf("[Slice:B] Length is %d Capacity is %d\n", len(b3), cap(b3))

	fmt.Println("Slice B after copying:", b3)

	var slice1 = []string{"india", "japan", "canada"}
	var slice2 = []string{"australia", "russia"}
	slice3 := append(slice2, slice1...)
	fmt.Println(slice3)

	var slice []int
	fmt.Println(slice == nil)

	currencyMode := make(map[string]string)
	currencyMode["USD"] = "United States Dollar"
	currencyMode["EUR"] = "Euro"
	currencyMode["GBP"] = "British Pound"
	currencyMode["JPY"] = "Japanese Yen"
	currencyMode["INR"] = "Indian Rupee"
	fmt.Println(currencyMode)

	a5 := map[string]string{
		"USA":    "United States",
		"UK":     "United Kingdom",
		"Canada": "Canada",
	}

	fmt.Println(a5["USA1"])
	name, ok := a5["USA"]
	if ok {
		fmt.Println("Currency name for currency code", name, " is", ok)
	}

	// if name, ok := a5["USD"]; ok {
	// }

	//b5 := map[string]string

	for code, name := range a5 {
		fmt.Println("Currency name for currency code", code, " is", name)
	}

	delete(a5, "EUR")
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1])
}
