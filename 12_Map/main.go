package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20, "Rocky": 30, "Rajiv": 40, "Kate": 50}

	for key, _ := range employee {
		delete(employee, key)
	}

	employee = make(map[string]int)

	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}
	fmt.Println(unSortedMap)

	keys := make([]string, 0, len(unSortedMap))

	for key := range unSortedMap {
		keys = append(keys, key)
	}

	fmt.Println(keys)

	sort.Strings(keys)

	fmt.Println(keys)

	for _, key := range keys {
		fmt.Printf("Country: %s, Population: %d\n", key, unSortedMap[key])
	}

	values := make([]int, 0, len(unSortedMap))

	for _, value := range unSortedMap {
		values = append(values, value)
	}

	sort.Ints(values)

	for _, v := range values {
		fmt.Printf("Population: %d\n", v)
	}

	first := map[string]int{"a": 1, "b": 2, "c": 3}
	second := map[string]int{"a": 30, "e": 5, "c": 3, "d": 4}

	for k, v := range second {
		first[k] = v
	}

	fmt.Println(first, second)

	curUSD := currency{
		name:   "US Dollar",
		symbol: "$",
	}
	curGBP := currency{
		name:   "Pound Sterling",
		symbol: "£",
	}
	curEUR := currency{
		name:   "Euro",
		symbol: "€",
	}

	currencyCode := map[string]currency{
		"USD": curUSD,
		"GBP": curGBP,
		"EUR": curEUR,
	}

	fmt.Println(currencyCode)

	for code, info := range currencyCode {
		fmt.Printf("Currency Code: %s, Name: %s, Symbol: %s\n", code, info.name, info.symbol)
	}

	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}

	fmt.Println("Original employee salary", employeeSalary)
	modified := employeeSalary
	modified["mike"] = 18000
	fmt.Println("Employee salary changed", employeeSalary)

	arr := []int{15, 42, 3, 7, 21, 8, 9, 2, 23, 44, 1}
	max, min := findMaxAndMin(arr)
	fmt.Printf("Max: %d, Min: %d\n", max, min)

	words := []string{"apple", "banana", "apple", "orange", "banana", "apple", "grape", "orange"}
	fmt.Println("Original array:", words)

	result := wordFrequency(words)
	fmt.Println("Word frequency:", result)

	arrInt := []int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", arrInt)

	mutipleByTwo := func(n int) int {
		return n * 2
	}

	mapResult := Map(arrInt, mutipleByTwo)
	fmt.Println("Modified array:", mapResult)

	userNames := make([]string, 2, 5)

	userNames[0] = "abc"
	userNames = append(userNames, "xyz")
	userNames = append(userNames, "pqr")

	fmt.Println("Modified array:", userNames)

	rating := make(floatMap, 3)
	rating["A"] = 4.5
	rating["B"] = 3.2
	rating["C"] = 5.0

	rating.output()

}

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m["A"], m["B"], m["C"])
}

func Map(arr []int, fn func(int) int) []int {
	result := make([]int, len(arr))

	for i, num := range arr {
		result[i] = fn(num)
	}

	return result
}

func wordFrequency(words []string) map[string]int {
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func findMaxAndMin(arr []int) (int, int) {
	if len(arr) == 0 {
		panic("Array is empty")
	}

	max := math.MinInt64 // 1
	min := math.MaxInt64

	for _, num := range arr {
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}
	return max, min
}

type currency struct {
	name   string
	symbol string
}
