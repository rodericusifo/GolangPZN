package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 75

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus bool = lulusUjian && lulusAbsensi
	fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absensi >= 80)
}
