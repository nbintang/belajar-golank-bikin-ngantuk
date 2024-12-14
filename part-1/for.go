package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("tessss", counter)
		counter++
	}

	for counters := 1; counters <= 10; counters++ {
		fmt.Println("tes", counters)
	}


	names := []string{"bintang", "Hidayat", "Nur"}


	for index, name := range names {
		fmt.Println(index, name)
	}
}
