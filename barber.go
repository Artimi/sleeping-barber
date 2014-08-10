package main

import (
	"flag"
	"fmt"
	"time"
)

type Params struct {
	Customers      int
	Seats          int
	Customer_delay float64
	Barber_delay   float64
}

func barber(customer_ready chan int) {
	fmt.Printf("I am a Barber.\n")
	for {
		customer_id := <-customer_ready
		fmt.Printf("BARBER: Customer %d served.\n", customer_id)
	}
}

func customer(id int, customer_ready chan int) {
	fmt.Printf("I am a Customer #%d.\n", id)
	customer_ready <- id

}

func read_params() *Params {
	var params *Params = new(Params)
	flag.IntVar(&params.Customers, "customers", 10, "number of customers to come to barber shop")
	flag.IntVar(&params.Seats, "seats", 3, "number of seats in barber shop")
	flag.Float64Var(&params.Customer_delay, "customers_delay", 100, "mean value of time of customers incoming to shop")
	flag.Float64Var(&params.Barber_delay, "barber_delay", 150, "mean value of time of barbers work on customer")

	flag.Parse()
	return params
}

func main() {
	params := read_params()

	customer_ready := make(chan int)

	go barber(customer_ready)
	for id := 0; id < params.Customers; id++ {
		go customer(id, customer_ready)
	}

	time.Sleep(250 * time.Millisecond)
}
