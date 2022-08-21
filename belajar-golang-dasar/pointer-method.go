package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	ifo := Man{"Ifo"}
	ifo.Married()

	fmt.Println(ifo.Name)
}