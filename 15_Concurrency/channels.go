package main

import (
	"fmt"
	"time"
)

func hello1(done chan bool) {
	fmt.Println("Hello, World!")
	time.Sleep(5 * time.Second)
	done <- false // signal to main goroutine that hello1 is done
}

func main2() {
	// ch := make(chan int)
	// ch <- 1 means set the channel
	// x := <- ch get from channel
	done := make(chan bool)
	go hello1(done)
	<-done // wait for hello1 to finish
	fmt.Println("Main function complete")
}
