package main

import "fmt"

func main() {
	for ranges := 1; ranges < 10; ranges++ {
		ifFive := ranges%2 == 2
		if ifFive {
			continue
		}
		fmt.Println("done", ranges)
	}
}
