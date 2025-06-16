package main

import "fmt"

const age = 30

//not allowed outside
// name := "golang"

func main() {
	const name string = "golang"

	// not able to assign it again
	// name = "javascript"

	// fmt.Println(age)

	// Multiple constants
	const (
		port = 5000
		host = "localhost"
	)

	//not allowed
	// host := "golang"

	fmt.Println(port, host)
}
