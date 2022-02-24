package main

import (
	"fmt"
)

func modify1(x int) {
	//拷贝了有一个x ，他不是原来的a的地址
	x = 100

	fmt.Println(&x)

}

func modify2(x *int) {
	//对指针变量取值，再赋值
	//传过来一个内存地址，把这个内存地址中的值改了
	*x = 100
}

func newMake() {
	//声明
	var a *int
	//分配内存,初始化
	a = new(int)
	*a = 100
	fmt.Println(*a)

	//声明
	var b map[string]int
	//分配内存,初始化
	b = make(map[string]int, 10)
	b["baba"] = 8
	fmt.Println(b)
	/*
	   1 两者都是分配内存
	   2 make 只用于 slice map 以及channel 的初始化 返回的还是这三个引用类型
	   3 new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针
	*/
}

func main() {
	a := 10
	fmt.Println(&a)
	a = 200
	fmt.Println(&a)
	modify1(a)
	fmt.Println(a)
	//对变量取内存地址,获得变量的指针变量
	modify2(&a)
	fmt.Println(a)
	//newMake()

}
