package main

import "fmt"

type dog struct{}

func (d dog) say() {
	fmt.Println("aaa")
}

type cat struct{}

func (c cat) say() {
	fmt.Println("bbb")
}

type person struct {
	name string
	age  int
}

func (p person) say() {
	fmt.Printf("ccc %s", p.name)
}

func (p person) move() {
	fmt.Printf("ddd %d", p.age)
}

type sayer interface {
	say()
}

// 值接收者
func da(arg sayer) {
	arg.say()
}

func main() {
	//实例化类
	c1 := cat{}
	da(c1)
	c2 := dog{}
	da(c2)
	c3 := person{
		name: "水水水水",
		age:  99,
	}
	da(c3)
	fmt.Println(c3)

	var s sayer
	//实例化类
	c4 := cat{}
	//实例赋值
	s = c4
	fmt.Println(s)
	c3.move()

	//空接口  函数的参数
	var x interface{}
	x = "hello"
	fmt.Println(x)
	//    空接口 map的value
	var m = make(map[string]interface{})
	m["name"] = "水水水水"
	m["age"] = 19
	m["bobby"] = []string{"a", "b", "c"}
	fmt.Println(m)

}
