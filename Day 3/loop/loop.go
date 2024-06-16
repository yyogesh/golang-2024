package main

import (
	"fmt"
)

func main() {
	i := 1
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	j := 1
	for j < 10 {
		fmt.Println(j)
		j++
	}

	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
			break
		}
	}
	fmt.Println("********************************")
	for k := 1; k < 10; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Println(k)
	}

	fmt.Println("********************************")

	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("********************************")

	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
		}
	}

	fmt.Println("********************************")
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break
			}
		}
	}

	fmt.Println("********************************")
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}
	}

	fmt.Println("********************************")

	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	fmt.Println("********************************")

	for {
		fmt.Println("Hello")
	}
}
