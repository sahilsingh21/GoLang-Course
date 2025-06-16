package main

import "fmt"

// numberd sequence of specific length
func main() {
	// zeroed value
	// int -> 0, string -> "", bool -> false
	// var nums [4]int

	// nums[0] = 1

	// fmt.Println(nums[0]) // access first element
	// fmt.Println(nums[1]) // access second element
	// fmt.Println(nums)    // print entire array
	// fmt.Println(len(nums)) // length of array

	// var vals [4]bool
	// vals[2] = true
	// fmt.Println(vals) // [false false false false]

	// var names [3]string
	// names[0] = "golang"
	// fmt.Println(names)

	// to declare it in one line
	// nums := [3]int{1, 2, 3} // array literal
	// fmt.Println(nums) // [1 2 3]

	// 2d arrays
	nums := [2][2]int{{3, 4}, {5, 6}} // 2d array literal
	fmt.Println(nums)                 // [[3 4] [5 6]]

	// - fixed size, that is predictable
	// - Memory optimization
	// - Contant time access
}
