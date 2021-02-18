package main

import "fmt"

// 继承

//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) SetDreams(dreams []string) {
	//p.dreams = dreams
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

func main() {
	p1 := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams) // ?

	//d1 := &Dog{
	//    Feet: 4,
	//    Animal: &Animal{ //注意嵌套的是结构体指针
	//        name: "乐乐",
	//    },
	//}
	//d1.wang() //乐乐会汪汪汪~
	//d1.move() //乐乐会动！
}
