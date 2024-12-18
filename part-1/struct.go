package main

import "fmt"

type Customer struct {
	name, address string
	age           int
}

func showCustomer() {
	adit := Customer{
		name:    "Adit",
		address: "Indonesia",
		age:     21,
	}
	fmt.Println(adit)
	adit.sayHello("adit")
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.name)
}

func main() {
	showCustomer()

}
