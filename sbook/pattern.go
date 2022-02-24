package main

import (
	"fmt"
	"regexp"
)

func main() {
	searchIn := "John : 234.2 willam : 99082.3 steve : 90882.1"
	pat := "[0-9]+.[0-9]"

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("match found")
	}

	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)

}
