package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var KTPIfo NoKTP = "1234567890"
	var marriedStatus Married = false
	fmt.Println(KTPIfo)
	fmt.Println(marriedStatus)
}