package main

import (
	"time"
)

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

// construction hack
func newOrder(id string, amount float32, status string) *order {
	// initial setup goes here...
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

// To connect above struct with first letter
// receiver type || pointer to change value
func (o *order) changeStatus(status string) {
	// struct do automatically De-reference
	o.status = status
}

// We are not modified value that's why we not using pointer
func (o order) getAmount() float32 {
	return o.amount
}

func main() {

	// // inline struct
	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{"golang", true} // assign in same order

	// fmt.Println(language)

	// myOrder := newOrder("1", 30.50, "received")
	// fmt.Println(myOrder)

	// if you don't set any field, default value is zero value
	// int => 0, float => 0, string "", bool => false
	// myOrder := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "received",
	// }

	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder.getAmount())
	// myOrder.createdAt = time.Now()

	// var myOrder2 order = order{
	// 	id:        "2",
	// 	amount:    100,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }

	// fmt.Println(myOrder.status)
	// fmt.Println("Order struct", myOrder)
	// fmt.Println("2nd Order struct", myOrder2)
}
