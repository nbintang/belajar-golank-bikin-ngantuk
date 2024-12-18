package main

import "fmt"

type AddressProps struct {
	city, province, country string
}

type Man struct {
	name string
}

func ChangeAddress(address *AddressProps, name string) {
	address.city = name
}
func main() {

	//pass by value
	/*
	   address := AddressProps{"Subang", "Jawa Barat", "indonesia"}
	   	address2 := address
	   	address2.city = "Jakarta"
	   	fmt.Println(address)
	   	fmt.Println(address2)
	*/

	// pass by refrences
	/*	address := AddressProps{"Subang", "Jawa Barat", "indonesia"}
		address2 := &address

		address2.city = "Jakarta"
		fmt.Println(address)
		fmt.Println(address2)
	*/

	//asterisk
	/* 	address := AddressProps{"Subang", "Jawa Barat", "indonesia"}
	   	address2 := &address

	   	address2.city = "Jakarta"

	   	*address2 = AddressProps{"Depok", "Jawa Barat", "Indonesia"}
	   	fmt.Println(address)
	   	fmt.Println(address2)
	*/

	address := &AddressProps{
		city:     "",
		province: "Sumatera Barat",
		country:  "Indonesia",
	}

	ChangeAddress(address, "Pariaman")
	fmt.Println(address)

	isMan()
}

func (man *Man) Married() {
	man.name = "Mr. " + man.name
}

func isMan() {
	bintang := Man{name: "Bintang"}
	bintang.Married()
	fmt.Println(bintang.name)
}
