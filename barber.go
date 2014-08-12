package main

import (
	"flag"
	"fmt"
	"time"
)

func serve_customer(customer_id int, delay int, customer_served chan int) {
	time.Sleep(time.Duration(delay) * time.Millisecond)
	customer_served <- customer_id
	fmt.Printf("BARBER: Customer %d served.\n", customer_id)
}

func barber(delay int, seats chan int, customer_served chan int, going_home chan int, customer_to_serve int) {
	var customer_id int = -1
	var number_of_customers_served int = 0
	for {
		select {
		case customer_id = <-seats:
		default:
			fmt.Printf("BARBER: Sleeping.\n")
			customer_id = <-seats
			fmt.Printf("BARBER: Woken up by customer %d.\n", customer_id)
		}
		serve_customer(customer_id, delay, customer_served)
		number_of_customers_served++
		if number_of_customers_served == customer_to_serve {
			break
		}
	}
	fmt.Printf("BARBER: Closing shop, going home.\n")
	going_home <- 1
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

type Params struct {
	Customers             int
	Seats                 int
	Customer_delay        int
	Customer_return_delay int
	Barber_delay          int
}

func read_params() *Params {
	var params *Params = new(Params)
	flag.IntVar(&params.Customers, "customers", 10, "Number of customers to come to barber shop")
	flag.IntVar(&params.Seats, "seats", 3, "Number of seats in barber shop")
	flag.IntVar(&params.Customer_delay, "customers_delay", 5, "Time between customers incoming to shop [ms]")
	flag.IntVar(&params.Customer_return_delay, "customers_return_delay", 20, "Time between customers' retry if the barber shop is full [ms]")
	flag.IntVar(&params.Barber_delay, "barber_delay", 15, "Time of barber working on customer [ms]")

	flag.Parse()
	return params
}

func main() {
	var params *Params = read_params()

	seats := make(chan int, params.Seats)
	customer_served := make(chan int)
	barber_going_home := make(chan int)

	go barber(params.Barber_delay, seats, customer_served, barber_going_home, params.Customers)
	for id := 0; id < params.Customers; id++ {
		go customer(id, params.Customer_return_delay, seats, customer_served)
		time.Sleep(time.Duration(params.Customer_delay) * time.Millisecond)
	}
	<-barber_going_home
}
