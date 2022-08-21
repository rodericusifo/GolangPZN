package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuuuu from", a.Name)
}

func main() {
	var ifo Customer
	ifo.Name = "Ifo"
	ifo.Address = "Indonesia"
	ifo.Age = 23

	ifo.sayHi("Joko")
	ifo.sayHuuu()

	// fmt.Println(ifo.Name)
	// fmt.Println(ifo.Address)
	// fmt.Println(ifo.Age)

	// joko := Customer{
	// 	Name:    "Joko",
	// 	Address: "Indonesia",
	// 	Age:     35,
	// }
	// fmt.Println(joko)

	// budi := Customer{"Budi", "Indonesia", 35}
	// fmt.Println(budi)
}
