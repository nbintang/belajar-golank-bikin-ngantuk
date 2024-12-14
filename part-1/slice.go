package main

import "fmt"

func main() {
	var arrayExample = [...]string{
		"tes",
		"tes2",
		"tes3",
	}

	var tes = arrayExample[1:]
	fmt.Println(tes)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(days)
	daySlice := days[5:]
	fmt.Println(daySlice)

	daySlice[1] = "Jumat"
	fmt.Println(daySlice)
	fmt.Println(days)

	liburBaru := append(daySlice, "Libur Baru")

	fmt.Println(liburBaru)
	fmt.Println(days)


	newSlice := make([]string, 2,3)
	newSlice[0] = "Nur"
	newSlice[1] = "Bintang"

	fmt.Println(newSlice)
	isiLen := len(newSlice)
	isiCap := cap(newSlice)
	fmt.Println(isiLen)
fmt.Println(isiCap)

	fromSLice := days[:]
	fmt.Println(fromSLice)
	toSlice := make([]string, len(fromSLice), cap(fromSLice))
	fmt.Println(toSlice)

	fmt.Println(copy(fromSLice, toSlice))
	
}
