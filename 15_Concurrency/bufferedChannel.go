package main

import "fmt"

func main() {
	// ch := make(chan string, 2)
	// ch <- "Hello"
	// ch <- "World"
	// ch <- "World123"
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	ch := make(chan int, 5)
	ch <- 5
	ch <- 6
	close(ch)
	n, open := <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
}
