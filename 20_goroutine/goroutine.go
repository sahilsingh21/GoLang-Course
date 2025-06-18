package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task", id)
}

// main() is also go thread
// when { go task(i) } thread added in schedualer
// func main() thread end first so inside task(i) thread will not show any output
func main() {
	// concurrent(Parallel) Processing only added go
	for i := 0; i <= 10; i++ {
		// go task(i)

		//closure
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	// main() func hold for 2 second to end
	time.Sleep(time.Second * 2)
}
