package main

import "fmt"

func main() {
	var name = "Ifo"

	if name == "Ifo" {
		fmt.Println("Hello Ifo")
	}
	if name == "Joko" {
		fmt.Println("Hello Joko")
	}
	if name == "Budi" {
		fmt.Println("Hello Budi")
	}else {
		fmt.Println("Hello kenalan donk")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}