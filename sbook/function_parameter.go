package main

import "fmt"

//函数作为参数
func main() {
	callback(1, Add)
	var g int
	func(i int) {
		s := 0
		for j := 0; j <= i; j++ {
			s += j
		}
		g = s // 闭包，函数内部没有g 变量，使得去找函数外部
	}(100)
	fmt.Println(g)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is : %d\n", a, b, a+b)
}

//func callback(x int,f func(a,b int))  {
func callback(x int, f func(int, int)) {
	f(x, 2)
}
