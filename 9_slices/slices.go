package main

import (
	"fmt"
)

// slice -> dynamic array
// most used  construct in go
// + useful methods
func main() {
	// uninitialized slice is nil
	// var nums []int

	// fmt.Println(nums == nil) // true
	// fmt.Println(len(nums))   // 0

	// zeroed slice // nums := []int{} // empty slice
	// var nums = make([]int, 0, 5)
	// capacity is the maximum number of elements that can be stored in the slice
	// nums := []int{}

	// fmt.Println(cap(nums))

	// fmt.Println(nums) // [0 0]
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// fmt.Println(nums)

	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// copying slices
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// // copy function
	// copy(nums2, nums)
	// // nums2 = nums // this will not work, it will not copy the elements, it will just point to the same slice

	// fmt.Println(nums)  // [2]
	// fmt.Println(nums2) // [0 0 0 0 0]

	// slice operator
	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[0:2]) // [1 2]
	// fmt.Println(nums[1:])  // [2 3]
	// fmt.Println(nums[:2])  // [1 2]

	// slice package
	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 4}

	// fmt.Println(slices.Equal(nums1, nums2)) // false
	// fmt.Println(slices.Contains(nums1, 2))  // true
	// fmt.Println(slices.Index(nums1, 2))     // 1

	// 2d slices
	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums) // [[1 2 3] [4 5 6]]
}
