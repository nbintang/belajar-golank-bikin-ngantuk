package main

import "fmt"

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value - 1)
	}
}

func main() {
	result := factorialRecursive(10)
	fmt.Println(result)

}