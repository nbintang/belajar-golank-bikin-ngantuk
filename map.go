package main

import "fmt"

func main() {
	maps := map[int]string{
		1: "bintang",
		2: "hidayat",
	}

	fmt.Println(maps[2])

	books := make(map[string]string)

	books["title"] = "hidayat"
	books["description"] = "lorem ipsum dolot amet"
	books["author"] = "Bintang"

	delete(books, "author")
	fmt.Println(books)
}
