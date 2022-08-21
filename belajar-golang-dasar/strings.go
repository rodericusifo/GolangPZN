package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Rodericus Ifo", "Ifo"))
	fmt.Println(strings.Contains("Rodericus Ifo", "Budi"))

	fmt.Println(strings.Split("Rodericus Ifo Krista", " "))

	fmt.Println(strings.ToLower("Rodericus Ifo Krista"))
	fmt.Println(strings.ToUpper("Rodericus Ifo Krista"))
	fmt.Println(strings.Title("rodericus ifo krista"))

	fmt.Println(strings.Trim("      Rodericus Ifo     ", " "))
	fmt.Println(strings.ReplaceAll("Ifo Joko Ifo", "Ifo", "Budi"))
}
