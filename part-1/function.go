package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello World!")
}

// with parameter
func sayHelloTo(firstName string, lastName string) string {
	return firstName + " " + lastName
}

// return multiple values
func returnMultipleValue() (string, string) {
	return "Bintang", "hidayat"
}

// return named values
func namedReturnValue() (firstName, lastName string) {
	firstName = "Nur"
	lastName = "Hidayat"

	return firstName, lastName
}

// variadic function
func sumAll(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func functionAsAValue(name string) string {
	return name
}

type namingPersonProps func(string) string

func funcAsParams(name string, namingPerson namingPersonProps  ) (string, namingPersonProps) {
	return name, namingPerson
}

func main() {
	sayHello()
	f, _ := returnMultipleValue()
	fmt.Println(f)

	firstName, _ := namedReturnValue()
	fmt.Println(firstName)
	fmt.Println("Hello,", sayHelloTo("Bintang", "Hidayat"))
	fmt.Println(sumAll(10, 10, 120, 10, 10))

	namePerson := functionAsAValue
	fmt.Println(namePerson("Bintang"))

	name, namedPerson := funcAsParams("Bintang", functionAsAValue)
	fmt.Println(name, namedPerson("Hidayat"))
}
