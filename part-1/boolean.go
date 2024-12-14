package main

import "fmt"

func main() {
	var (
		a = true
		b = false
		c = a
	)

	if a && c == b {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	var (
		nilaiAkhir = 90
		absensi    = 70
	)

	var nilaiKKM = 80
	var lulusNilai bool = nilaiAkhir >= nilaiKKM
	var lulusAbsensi bool = absensi >= nilaiKKM
	var lulus bool = lulusNilai && lulusAbsensi
	fmt.Println(lulus)
}
