package main

import "fmt"

type BlackListedProps func(string) bool

func register(name string, isBlackListed BlackListedProps) {
	if isBlackListed(name) {
		fmt.Println("You were blocked, sorry")
	} else {
		fmt.Println("Welcome to dashboard", name)
	}
}

func isBlackListed(name string) bool {
	if name == "Bintang" {
		return true
	}
	return false
}

func main() {
	register("Bintang", isBlackListed)
	register("Hidayat", isBlackListed)
}
