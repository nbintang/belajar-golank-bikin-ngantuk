package main

import "fmt"

func main() {
	const (
		fistName   = "Bintang"
		middleName = "Nur"
	)
	getLastName := showLastName()
	fmt.Println(fistName, middleName, getLastName)
}

func showLastName() string {
	const lastName string = "Hidayat"
	return lastName
}
