package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  byte
}

func main() {

	p := new(Person)
	p.name = "aaa"
	p.age = 2
	p.sex = 1
	fmt.Println(p.name)
}
