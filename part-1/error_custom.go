package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func checkData(id string) error {
	if id == "" {
		return &validationError{Message: "No ID Provided..."}
	}

	return nil
}

func main() {
	err := checkData("2")
	if err != nil {
		if validationError, ok := err.(*validationError); ok {
			fmt.Println("err", validationError.Error())
		} else {
			fmt.Println("unknown", err.Error())
		}
	} else {
		fmt.Println("Succeed")
	}
}

