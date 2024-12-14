package main

import "fmt"

func main() {
	var a = 10
	var b = 3
	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	fmt.Println(c, d, e, f)

	var i = 10 - 1
	i += 1

	fmt.Println(i)
}