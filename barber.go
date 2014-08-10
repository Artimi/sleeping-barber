package main

import (
	"flag"
	"fmt"
)

type Params struct {
	Customers      int
	Seats          int
	Customer_delay float64
	Barber_delay   float64
}

func barber() {
	fmt.Printf("I am a Barber.\n")
}

func customer() {
	fmt.Printf("I am a Customer.\n")
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
	fmt.Printf("%#v\n", params)
	fmt.Printf("Hello, Barber.\n")

	go barber()
	go customer()

	var input string
	fmt.Scanln(&input)
}
