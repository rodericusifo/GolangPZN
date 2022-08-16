package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Ifo",
		"address": "Blitar",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = map[string]string{}
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Ifo"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
