package main

import "fmt"

func main() {
	var arrayExample [5]string

	arrayExample[0] = "Bintang"
	arrayExample[1] = "Nur"
	arrayExample[2] = "Hidayat"
	arrayExample[3] = "Bintang"
	arrayExample[4] = "Hidayat"
	fmt.Println(arrayExample)

	var arrayExample2 = [...]int{
		1,
		2,
		3,
		4,
	}

	fmt.Println(arrayExample2[3])
	fmt.Println(arrayExample2)
	arrayExample2[3] = 0
	var tesString = "Nur Bintang Hidayat"
	fmt.Println(len(tesString))
	fmt.Println(arrayExample2)

}
