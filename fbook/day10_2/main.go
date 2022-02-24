package main

import "fmt"

//结构体继承

//Animal 动物,结构体
type Animal struct {
	name string
}

// 结构体方法
func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗 结构体
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

// 方法
func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	// 构造 没有定义构造函数，直接赋值, &Dog 指针Dog
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}
