package main

import (
	"fmt"
	"strings"
)

func kebian(x ...int) int {
	//    可变参数x
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

func kebian2(x int, y ...int) int {
	for _, v := range y {
		x = x + v
	}
	return x
}

// 多返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 定义返回值
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func add(x, y int) int {
	return x + y
}

// 函数传入函数
func calc1(x, y int, op func(int, int) int) int {
	return op(x, y)

}

//函数的返回值是函数
func abc() func() {
	name := "dddd"
	return func() {
		// 查找外层变量
		fmt.Println("aaaaaaaa", name)
	}
}

func abcd(name string) func() {
	//定义，接收一个参数
	return func() {
		// 查找外层变量
		fmt.Println("aaaaaaaa", name)
	}
}

// 闭包=函数+外层变量的应用
func checkSuffix(suffix string) func(string) string {
	//定义函数，返回函数,函数也返回值
	return func(name string) string {
		//定义匿名函数，接收参数 name
		if !strings.HasSuffix(name, suffix) {
			//检测匿名函数的参数是否是 外层函数结尾
			return name + suffix
		}
		// esle true 直接返回 name
		return name
	}
}

func calcs(base int) (func(int) int, func(int) int) {
	addd := func(a int) int {
		base = base + a
		return base
	}

	subb := func(b int) int {
		base = base - b
		return base
	}
	return addd, subb
}

func main() {
	k, l := calcs(1)
	k1 := k(2)
	k2 := k(3)
	k3 := k(4)
	fmt.Println(k1) //base = 1 , a = 2   return 3
	fmt.Println(k2) //base = 3 , a = 3   return 6
	fmt.Println(k3) // base = 6 ,a = 4   return 10
	l1 := l(5)      // base = 10  ,b = 5 return 5
	fmt.Println(l1)
	//aa := kebian(1, 2, 3, 4, 5)
	//bb := kebian2(1, 2, 3, 4, 5)
	//fmt.Println(aa)
	//fmt.Println(bb)
	// 将add函数传入calc1中 op = add
	//ret := calc1(2, 3, add)
	//fmt.Println(ret)
	////匿名函数
	//add := func(x, y int) {
	//    fmt.Println(x + y)
	//}
	//add(10,20)
	////自执行函数()
	//func(x,y int) {
	//    fmt.Println(x+y)
	//}(10,20)
	//
	//f :=abcd("aaaaas")
	//f()
	f := checkSuffix(".txt")
	ret := f("阿萨大大")
	ret2 := f("欧康纳")
	fmt.Println(ret, ret2)
}
