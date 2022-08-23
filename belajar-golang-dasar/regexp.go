package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("i([a-z])o")

	fmt.Println(regex.MatchString("ifo"))
	fmt.Println(regex.MatchString("ito"))
	fmt.Println(regex.MatchString("iDo"))

	search := regex.FindAllString("iko ika ido ito iyo iki", -1)
	fmt.Println(search)
}