package main

import "fmt"

// 函数的返回值是函数
func main() {
	var f = Adder()
	fmt.Print(f(1), "-")
	fmt.Print(f(10), "-")
	fmt.Print(f(100), "-")
}

func Adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
