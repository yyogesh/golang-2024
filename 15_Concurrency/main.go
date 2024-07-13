package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello, World!")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}

	// <- //  => // ->
}

func main() {
	go numbers()
	go alphabets()
	time.Sleep(15 * time.Second)
	fmt.Println("main function")
}
