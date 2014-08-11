package main

import (
	"flag"
	"fmt"
	"time"
)

type Params struct {
	Customers      int
	Seats          int
	Customer_delay int
	Barber_delay   int
}

func barber(params Params, seats chan int, customer_served chan int) {
	for {
		customer_id := <-seats
		time.Sleep(time.Duration(params.Barber_delay) * time.Millisecond)
		customer_served <- customer_id
		fmt.Printf("BARBER: Customer %d served.\n", customer_id)
	}
}

func customer(id int, seats chan int, customer_served chan int) {
	select {
	case seats <- id:
		fmt.Printf("CUSTOMER %d: Entering barber shop.\n", id)
		finished_customer := -1
		for finished_customer != id {
			finished_customer = <-customer_served
		}
		fmt.Printf("CUSTOMER %d: Leaving barber shop happy.\n", id)
	default:
		fmt.Printf("CUSTOMER %d: Leaving barber shop unserviced.\n", id)
	}
}

func read_params() *Params {
	var params *Params = new(Params)
	flag.IntVar(&params.Customers, "customers", 10, "number of customers to come to barber shop")
	flag.IntVar(&params.Seats, "seats", 3, "number of seats in barber shop")
	flag.IntVar(&params.Customer_delay, "customers_delay", 5, "mean value of time of customers incoming to shop")
	flag.IntVar(&params.Barber_delay, "barber_delay", 15, "mean value of time of barbers work on customer")

	flag.Parse()
	return params
}

func main() {
	var params *Params = read_params()

	seats := make(chan int, params.Seats)
	customer_served := make(chan int)

	go barber(*params, seats, customer_served)
	for id := 0; id < params.Customers; id++ {
		go customer(id, seats, customer_served)
		time.Sleep(time.Duration(params.Customer_delay) * time.Millisecond)
	}

	time.Sleep(500 * time.Millisecond)
}
