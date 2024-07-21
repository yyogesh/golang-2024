package main

import "time"

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	select {
	case msg1 := <-output1:
		println("Received from server1:", msg1)
	case msg2 := <-output2:
		println("Received from server2:", msg2)
	}
}

// MUTEX
// GIN
// SQLLITE
