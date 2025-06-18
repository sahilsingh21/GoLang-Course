package main

import "fmt"

// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// to prevent Duplicates
// printSlice[T any]  == any type
// printSlice[T interface{}]  == any type
// printSlice[T int | string]
// comparable most type is there
func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

//LIFO
type stack struct {
	elements []int
}

// Generics
type stack2[T any] struct {
	elements []T
}

func main() {
	nums := []int{1, 2, 3}
	// names := []string{"golang", "typeScript"}
	// printSlice(nums)
	// // printStringSlice(names)

	printSlice(nums, "Alice")
	// printSlice(names)

	myStack := stack{
		elements: []int{1, 2, 3},
	}

	fmt.Println(myStack)

	myStack2 := stack2[string]{
		elements: []string{"golang", "typeScript"},
	}

	fmt.Println(myStack2)

}
