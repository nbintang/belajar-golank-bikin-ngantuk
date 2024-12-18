package main

import "fmt"

type HasNameProps interface {
	GetName() string
}
type PersonProps struct{
	name string
}

func sayHelloName(hasName HasNameProps) {
	fmt.Println("Hello", hasName.GetName())
}

func (person PersonProps) GetName() string{
	return person.name
}



func returnAnyResult(anything interface{}) interface{}{
	return anything
}

func main() {
	person := PersonProps{
		name: "Bintang",
	}
	sayHelloName(person)
	fmt.Println(returnAnyResult(21))
}
