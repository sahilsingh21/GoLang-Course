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

// // recevied
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }

// func task(done chan bool){
// 	defer func () {
// 		done <- true
// 	}()
// 	fmt.Println("Processing...")

// 	// done <- true  // if write here working fine but if some case process is failed then this will never execute
// }

// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() { done <- true }()

// 	for email := range emailChan {
// 		fmt.Println("sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }

func main() {

	// STEP 6

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("reveived data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("reveived data from chan1", chan2Val)
		}
	}

	// // STEP 5

	// emailChan := make(chan string, 5)

	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("done sending")
	// //this is important
	// close(emailChan)
	// <-done

	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// // STEP 4

	// done := make(chan bool)

	// go task(done)

	// // (wait until the func task execute) (we are not storing)
	// <-done // block

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
	// result := make(chan int)

	// go sum(result, 4, 5)

	// res := <-result // blocking (don't need time sleep)

	// fmt.Println(res)

}
