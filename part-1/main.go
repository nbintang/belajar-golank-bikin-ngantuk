package main

import (
	"fmt"
	"golang-dasar/database"
	"golang-dasar/example"
	_"golang-dasar/internal"
)

func main() {
	result := example.SayHello("Bintang")
	fmt.Println(result)

	fmt.Println(database.GetDatabase())

}
