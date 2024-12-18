package main

import (
	"errors"
	"fmt"
)

func pembagian(value int, nums int) (int, error) {
	if nums == 0 {
		return 0, errors.New("Tidak bisa di bagi")
	} else {
		return value / nums, nil
	}
}

func main() {
	result, err := pembagian(12, 0)
	if( err ==nil){
	fmt.Println("hasil", result)
	} else{
		fmt.Println("Error", err.Error())
	}
}
