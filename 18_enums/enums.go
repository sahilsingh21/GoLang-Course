package main

import "fmt"

// enumerated types

type OrderStatus string

//iota (only for integer) Increase value auto
const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepaid"
	Delivered             = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	changeOrderStatus(Delivered)
}
