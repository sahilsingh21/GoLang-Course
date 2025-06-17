package main

import "fmt"

// iterating over data structure
func main() {

	// slice
	// nums := []int{6, 7, 8}

	// // Simple for loop
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// index, value
	// for i, num := range nums {
	// 	fmt.Println(num, i)
	// }

	// m := map[string]string{"fname": "john", "lname": "doe"}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// // single value always give us keys value (eg. k)
	// for k := range m {
	// 	fmt.Println(k)
	// }

	// unicode(unique code) code point rune
	// starting byte of rune
	// 255 -> 1 byte (normal ASCII have 255 characters)
	// 300 -> 2 byte (eg. ζ take 2 byte)
	for i, c := range "ζgolang" {
		fmt.Println(i, c, string(c))
	}

}
