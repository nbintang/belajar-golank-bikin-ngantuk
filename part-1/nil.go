package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {
	result := newMap("Bintang")
	if result == nil {
		fmt.Println("Empty")
	} else {
		fmt.Println(result["name"])
	}
}
