// package main

// import "fmt"

// func calcSquares(number int, squareop chan int) {
// 	sum := 0
// 	for number != 0 {
// 		digit := number % 10
// 		sum += digit * digit
// 		number /= 10
// 	}
// 	squareop <- sum // set channel
// }

// func calcCubes(number int, cubeop chan int) {
// 	sum := 0
// 	for number != 0 {
// 		digit := number % 10
// 		sum += digit * digit * digit
// 		number /= 10
// 	}
// 	cubeop <- sum // set channel
// }

// func main() {
// 	number := 589
// 	sqrch := make(chan int)
// 	cubech := make(chan int)

// 	go calcSquares(number, sqrch)
// 	go calcCubes(number, cubech)

// 	squares, cubes := <-sqrch, <-cubech
// 	fmt.Printf("Square of digits: %d\n", squares+cubes)
// }

// package main

// func main() {
// 	ch := make(chan int)
// 	ch <- 5 // unbuffered channel
// }

package main

import "fmt"

// import "fmt"

// func sendData(ch chan<- int) {
// 	ch <- 10
// }

// func main() {
// 	sendCh := make(chan int)
// 	go sendData(sendCh)
// 	fmt.Println(<-sendCh)
// }
func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	ch := make(chan int)
	go producer(ch)

	// for {
	// 	v, ok := <-ch
	// 	if ok == false {
	// 		break
	// 	}
	// 	fmt.Println("Received ", v)
	// }

	for v := range ch {
		fmt.Println("Received ", v)
	}
}
