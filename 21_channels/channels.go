package main

import (
	"fmt"
)

// // send
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("processing number", num)
// 		time.Sleep(time.Second)
// 	}
// }

// recevied
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

func main() {
	// messageChan := make(chan string)

	// //send data  to channel
	// messageChan <- "ping"

	// // received data from channel
	// msg := <-messageChan

	// // both are waiting to each to finsh that's why it's showing Deadlock
	// fmt.Println(msg)

	//STEP 2

	// numChan := make(chan int)

	// // Below both statment run parallel in different thread to prevent Deadlock
	// go processNum(numChan) // Different thread or chan
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// time.Sleep(time.Second * 2)

	// STEP 3
	result := make(chan int)

	go sum(result, 4, 5)

	res := <-result // blocking (don't need time sleep)

	fmt.Println(res)
}
