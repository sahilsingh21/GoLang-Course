package main

import "fmt"

func counter() func() int {
	var count int = 0

	// count is always available after execution also (eg. closures)
	return func() int {
		count += 1
		return count
	}
}

func main() {
	increament := counter()

	fmt.Println(increament())
	fmt.Println(increament())
}
