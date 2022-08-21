package main

import (
	"fmt"
	"GolangPZN/belajar-golang-dasar/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
