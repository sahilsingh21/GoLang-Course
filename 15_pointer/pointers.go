package main

import "fmt"

// // By value
// func changeeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num) // 5
// }

// By reference
func changeeNum(num *int) {
	// De-reference
	*num = 5
	fmt.Println("In changeNum", *num) // 5
}

func main() {
	num := 1
	changeeNum(&num)
	// fmt.Println("Memory address", &num)
	fmt.Println("After change in main", num)
}
