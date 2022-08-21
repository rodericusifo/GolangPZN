package main

import "fmt"

func sayHelloTo(fristname string, lastname string) {
	fmt.Println("Hello", fristname, lastname)
}

func main() {
	firstName := "Rodericus"
	sayHelloTo(firstName, "Ifo")
	sayHelloTo("Budi", "Nugraha")
}
