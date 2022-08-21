package main

import "fmt"

type Filter func(string)string

func spamFIlter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func sayHelloWithFilter(name string, filter Filter)  {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func main() {
	sayHelloWithFilter("Ifo", spamFIlter)

	filter := spamFIlter
	sayHelloWithFilter("Anjing", filter)
}