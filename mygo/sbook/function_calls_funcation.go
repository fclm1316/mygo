package main

import "fmt"

var a string

func main() {
	a = "G"
	fmt.Println(a)
	f1()
	f2()
}

func f1() {
	a := "O"
	print(a)
}

func f2() {
	fmt.Println(a)
}
