package main

import (
	"flag"
	"fmt"
	"time"
)

type Params struct {
	Customers             int
	Seats                 int
	Customer_delay        int
	Customer_return_delay int
	Barber_delay          int
}

func serve_customer(customer_id int, delay int, customer_served chan int) {
	time.Sleep(time.Duration(delay) * time.Millisecond)
	customer_served <- customer_id
	fmt.Printf("BARBER: Customer %d served.\n", customer_id)
}

func barber(delay int, seats chan int, customer_served chan int) {
	for {
		select {
		case customer_id := <-seats:
			serve_customer(customer_id, delay, customer_served)
		default:
			fmt.Printf("BARBER: Sleeping.\n")
			customer_id := <-seats
			fmt.Printf("BARBER: Woken up by customer %d.\n", customer_id)
			serve_customer(customer_id, delay, customer_served)
		}
	}
}

func customer(id int, return_delay int, seats chan int, customer_served chan int) {
	var serviced bool = false
	for serviced != true {
		select {
		case seats <- id:
			fmt.Printf("CUSTOMER %d: Entering barber shop.\n", id)
			finished_customer := -1
			for finished_customer != id {
				finished_customer = <-customer_served
			}
			fmt.Printf("CUSTOMER %d: Leaving barber shop happy.\n", id)
			serviced = true
		default:
			fmt.Printf("CUSTOMER %d: Leaving barber shop unserviced.\n", id)
			time.Sleep(time.Duration(return_delay) * time.Millisecond)
		}
	}
}

func read_params() *Params {
	var params *Params = new(Params)
	flag.IntVar(&params.Customers, "customers", 10, "number of customers to come to barber shop")
	flag.IntVar(&params.Seats, "seats", 3, "number of seats in barber shop")
	flag.IntVar(&params.Customer_delay, "customers_delay", 5, "time between customers incoming to shop")
	flag.IntVar(&params.Customer_return_delay, "customers_return_delay", 20, "time between customers' retry if the barber shop is full")
	flag.IntVar(&params.Barber_delay, "barber_delay", 15, "time of barber working on customer")

	flag.Parse()
	return params
}

func main() {
	var params *Params = read_params()

	seats := make(chan int, params.Seats)
	customer_served := make(chan int)

	go barber(params.Barber_delay, seats, customer_served)
	for id := 0; id < params.Customers; id++ {
		go customer(id, params.Customer_return_delay, seats, customer_served)
		time.Sleep(time.Duration(params.Customer_delay) * time.Millisecond)
	}

	time.Sleep(500 * time.Millisecond)
}
