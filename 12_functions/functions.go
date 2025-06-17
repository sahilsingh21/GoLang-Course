package main

import "fmt"

// mention return type else it will give us error
// func add(a string, b int)
// if attribute is same type then we can also do
// func add(a, b int) int {
// 	return a + b
// }

// Too many values to return
// func getLanguages() (string, string, bool) {
// 	return "golang", "javascript", true
// }

// In Argu we can pass function
// func processIt(fn func(a int) int){
// 	fn(1)
// }

// return function type
func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	// result := add(3, 5)

	// if we do not need lang3 value we can spress by adding _
	// lang1, lang2, _ := getLanguages()

	// lang1, lang2, lang3 := getLanguages()
	// fmt.Println(getLanguages())
	// fmt.Println(lang1, lang2, lang3)

	// fn := func(a int)int {
	// 	return 2
	// }

	// return function type
	fn := processIt()
	fmt.Println(fn(6))

}
