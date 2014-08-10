package main

import "fmt"

func barber() {
	fmt.Printf("I am a Barber.\n")
}

func customer() {
	fmt.Printf("I am a Customer.\n")
}

func main() {
	fmt.Printf("Hello, Barber.\n")

	go barber()
	go customer()

	var input string
	fmt.Scanln(&input)
}
