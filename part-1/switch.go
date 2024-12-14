package main

import "fmt"

func main() {
	name := "Bintang"

	switch name {
	case "Bintang":
		fmt.Println("Benar")
	case "bintang":
		fmt.Println("Salah")
	default:
		fmt.Println("gada hasil")
	}
}
