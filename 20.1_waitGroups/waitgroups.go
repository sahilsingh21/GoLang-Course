package main

import (
	"fmt"
	"sync"
)

// Wait group should pass as a pointer
func task(id int, w *sync.WaitGroup) {
	// defer is similer to react.js useState cleaner func, it run when func is completed
	defer w.Done() //go routine subtract 1
	fmt.Println("doing task", id)
}

func main() {
	var wg sync.WaitGroup // defult 0

	for i := 0; i <= 10; i++ {
		wg.Add(1) //go routine added 1
		go task(i, &wg)
	}

	wg.Wait() // if 0 then end else wait to finish task
}
