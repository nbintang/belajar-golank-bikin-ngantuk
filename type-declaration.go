package main

import "fmt"

func main() {
	type name string
	var firstName name = "Bintang"
	fmt.Println(firstName)
	fmt.Println(showMyName())
}

func showMyName() string {
	type fullName struct {
		firstName string
		lastName  string
	}
	var myFullName = fullName{"Bintang", "Hidayat"}
	return fmt.Sprintf("%s %s", myFullName.firstName, myFullName.lastName)
}
