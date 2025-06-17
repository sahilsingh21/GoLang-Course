package main

import "fmt"

// // variadic function
// // nums (Slice) integer type
// func sum(nums ...int) int {
// 	total := 0

// 	for _, num := range nums {
// 		total = total + num
// 	}
// 	return total
// }

// nums (Slice) any type we can use interface{}
func sum2(nums ...interface{}) string {

	for _, num := range nums {
		fmt.Println(num)
	}

	return "Noting"
}

func main() {

	// this called variadic function
	// we can pass n numbers of parameters (any typed)
	// fmt.Println(1, 2, 3, 4, "hello")
	nums := []interface{}{3, 4, 5, "golng"}
	result := sum2(nums)
	fmt.Println(result)
}
