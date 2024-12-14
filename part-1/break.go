package main

import "fmt"

func main() {
	for ranges := 1; ranges < 10; ranges++ {
		ifFive := ranges == 5
		if ifFive {
			break
		}
		fmt.Println("done", )
	}
}
