package main

import "fmt"

//MyInt 基于int类型的自定义类型
type MyInt int

// NewInt asdasd
type NewInt = int

func main() {
	var i MyInt
	fmt.Printf("type %T vale %v", i, i)
	var x NewInt
	fmt.Printf("type %T vale %v", x, x)
}
