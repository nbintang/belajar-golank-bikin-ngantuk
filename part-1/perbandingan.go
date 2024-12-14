package main

import "fmt"

func main() {
	var j = 10
	j++
	j++
	fmt.Println(j)
	var (
		tes  = "tes"
		tes2 = "tes"
	)
	if(tes != tes2) {
		fmt.Println("tidak sama")
	} else {
		fmt.Println("sama")
	}
}
