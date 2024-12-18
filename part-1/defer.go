package main

import "fmt"

func fooBar() (isError bool) {
	defer fmt.Println("Tes")
	isError = true
	return isError
}

func isFooBar() {
	isError :=  fooBar()
	if isError {
		panic("Something was Wrong")
	}
}

func main() {
	isFooBar()
}
