package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}

	doubled := Mapping(&numbers, double)
	tripled := Mapping(&numbers, triple)
	tripledx := Mapping(&numbers, func(n int) int {
		return n * 3 * 3
	})
	fmt.Println(tripledx)
	fmt.Println(doubled)
	fmt.Println(tripled)

	fn1 := getMapFn(&numbers)
	fn2 := getMapFn(&moreNumbers)

	doubled = Mapping(&numbers, fn1)
	tripled = Mapping(&numbers, fn2)

	fmt.Println(doubled)
	fmt.Println(tripled)

	result1 := sumOfArrays(1, numbers...)

	fmt.Println("result", result1)

	find(10, 30, 20, 10, 24)
	find(10, 30, 20, 10, 24, 20, 90, 40)

	// go find() // Goroutines
}

// variadic  function
func find(num int, nums ...int) {
	found := false
	for i, n := range nums {
		if n == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func sumOfArrays(b int, a ...int) int {
	sum := 0
	for _, num := range a {
		sum += num
	}
	return sum
}

type mapFn func(int) int

func Mapping(number *[]int, fn mapFn) []int {
	numbers := []int{}
	for _, num := range *number {
		numbers = append(numbers, fn(num))
	}

	return numbers
}

func getMapFn(number *[]int) mapFn {
	if (*number)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
