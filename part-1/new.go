package main

import "fmt"

type AddressProps struct {
	city, province, country string
}

func main() {
	address := new(AddressProps)
	address2 := address

	address2.city = "Jakarta"

	*address2 = AddressProps{"Depok", "Jawa Barat", "Indonesia"}
	fmt.Println(address)
	fmt.Println(address2)

}
