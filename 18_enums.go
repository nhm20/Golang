package main

//enumerated types

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Dispatched
)

func changeOrderStatus(status OrderStatus) {
	switch status {
	case Received:
		println("Order has been received")
	case Confirmed:
		println("Order has been confirmed")
	case Prepared:
		println("Order has been prepared")
	case Dispatched:
		println("Order has been dispatched")
	default:
		println("Unknown order status")
	}
}

type Gender string

const (
	Male   Gender = "male"
	Female        = "female"
)

func changeGender(gender Gender) {
	switch gender {
	case Male:
		println("Male is the Gender")
	case Female:
		println("Female is the Gender")
	}
}

func main() {
	changeOrderStatus(Received)
	// changeOrderStatus("Confirmed") // This will cause a compile-time error

	changeGender(Male)

}
