package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {
	// creating map
	// m := make(map[string]string)

	// // setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"

	// get an element
	// fmt.Println(m["name"], m["area"])

	//IMP: if key does not exists in the map than it return zero value
	// fmt.Println(m["phone"])

	// m := make(map[string]int)

	// m["age"] = 25
	// m["price"] = 100
	// fmt.Println(m["phone"])

	// fmt.Println(len(m))

	// delete(m, "price")
	// clear(m)

	// fmt.Println(m)

	// m := map[string]int{"price": 14, "phone": 3}
	// // fmt.Println(m)

	// v, ok := m["price"]

	// fmt.Println(v)

	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	// checking Equal or not
	m1 := map[string]int{"price": 14, "phone": 3}
	m2 := map[string]int{"price": 14, "phone": 3}

	fmt.Println(maps.Equal(m1, m2))
}
