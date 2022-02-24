package main

import "fmt"

type person1 struct {
	name, city string
	age        int8
}

func person2() {
	//键值对初始化
	p5 := person1{
		name: "小",
		city: "啊啊",
		age:  9,
	}

	fmt.Println(p5)
}

type test struct {
	//恰到好处的内存对齐
	//https://segmentfault.com/a/1190000017527311?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com
	a int8
	b int8
	c int8
	d int8
}

//Person 结构体
type Person struct {
	name string
	age  int8
}

//NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	//
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream Person做梦的方法. 给Person添加了Dream方法
func (p Person) Dream() {
	//func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	//    函数体
	//}
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

//指针类型的接收者，直接修改源值
func (p *Person) SetAge(newAge int8) {
	//需要修改接收者中的值
	//接收者是拷贝代价比较大的大对象
	//保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
	p.age = newAge
}

//值类型接收者，复制源数据修改
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("小王子", 25)
	p1.Dream()
	p1.SetAge(30)
	fmt.Println(p1.age)
	p1.SetAge2(50)
	fmt.Println(p1.age)
	//var p1 person1
	//p1.name = "合理"
	//p1.city = "转换"
	//p1.age = 2
	//fmt.Println(p1)
	//
	//var p2 struct {
	//    Name string
	//    Age  int
	//}
	//p2.Name="升级"
	//p2.Age=2
	//fmt.Println(p2)
	//new 关键字实体化
	//var p3=new(person1)
	//fmt.Println(p3)
	//取结构体的地址实例化
	//p4 := &person1{}
	//fmt.Println(p4)
	//p4.name="啊啊"
	//p4.age=9
	//fmt.Println(p4)
	//person2()
	//n:=test{1,2,3,4}
	//fmt.Printf("n.a %p \n",&n.a)
	//fmt.Printf("n.b %p \n",&n.b)
	//fmt.Printf("n.c %p \n",&n.c)
	//fmt.Printf("n.d %p \n",&n.d)

}
