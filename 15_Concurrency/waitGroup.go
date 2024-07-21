package main

import (
	"fmt"
	"sync"
	"time"
)

func proess(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("started Goroutine ", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", id)
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go proess(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished")
}
