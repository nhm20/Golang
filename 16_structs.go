package main

import (
	"fmt"
	"time"
)

// Structs are used to create complex data types in Go.
type order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time // nanosecond precision
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}
func (o *order) getAmount() float64 {
	return o.amount
}

// constructor function
func newOrder(id string, amount float64, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// struct embedding
type chip struct {
	name  string
	price float64
}
type iPhone struct {
	id     string
	amount float64
	chip   // embedding chip struct
}

func main() {
	// var order order=

	//if variable is not initialized, it will be set to the zero value of its type.
	myOrder := order{
		id:     "1",
		amount: 100.0,
		status: "received",
	}

	fmt.Println(myOrder)
	myOrder.createdAt = time.Now()
	fmt.Println(myOrder)
	fmt.Println(myOrder.status)

	myOrder.changeStatus("shipped")
	fmt.Println(myOrder)
	fmt.Println(myOrder.getAmount())

	// constructor function
	myOrder2 := newOrder("2", 200.0, "received")
	fmt.Println(myOrder2)
	fmt.Println(myOrder2.amount)

	language := struct {
		name   string
		isGood bool
	}{"golang", true}
	fmt.Println(language)

	// struct embedding
	myIphone := iPhone{
		id:     "1",
		amount: 1000.0,
	}
	fmt.Println(myIphone)

	newChip := chip{
		name:  "A15",
		price: 100.0,
	}
	myIphone2 := iPhone{
		id:     "2",
		amount: 2000.0,
		chip:   newChip,
	}
	fmt.Println(myIphone2)

	myIphone3 := iPhone{
		id:     "3",
		amount: 3000.0,
		chip:   chip{"A16", 200.0},
	}
	fmt.Println(myIphone3)
	myIphone3.chip.name = "A17"
	fmt.Println(myIphone3)


}
