package main

import (
	"encoding/json"
	"fmt"
)

type class struct {
	Classname string    `json:"classname"`
	Student   []student `json:"student"`
}

type student struct {
	Id     int `json:"id"`
	Name   string
	Gender string
	age    int
}

//newStudent 构造函数
func newStudent(Id int, Name string, Gender string, Age int) *student {
	return &student{
		Id:     Id,
		Name:   Name,
		Gender: Gender,
		age:    Age, //小写，私有。无法json
	}
}

func main() {
	// 初始化class
	c := class{
		Classname: "天才",
		// 申请内存
		Student: make([]student, 0, 20),
	}

	// 构造函数
	for i := 0; i < 10; i++ {
		s := newStudent(i, fmt.Sprintf("stud%02d", i), "男", i+10)
		//添加
		c.Student = append(c.Student, *s)
		//fmt.Println(c.Student)

	}

	//fmt.Printf("%#v",c)
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json load failed")
		return
	}
	fmt.Printf("%s\n", data)

	// 序列化，反序列化
	data2 := `{"Classname":"101","Student":[{"Id":0,"Gender":"男","Name":"stu00"},{"Id":1,"Gender":"男","Name":"stu01"},{"Id":2,"Gender":"男","Name":"stu02"},{"Id":3,"Gender":"男","Name":"stu03"},{"Id":4,"Gender":"男","Name":"stu04"},{"Id":5,"Gender":"男","Name":"stu05"},{"Id":6,"Gender":"男","Name":"stu06"},{"Id":7,"Gender":"男","Name":"stu07"},{"Id":8,"Gender":"男","Name":"stu08"},{"Id":9,"Gender":"男","Name":"stu09"}]}`
	//获得指针地址
	c1 := &class{}
	str := []byte(data2)
	//Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构,第二个参数必须是指针，否则无法接收解析的数据
	err = json.Unmarshal(str, c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}

//Student 学生
//type Student struct {
//    ID     int
//    Gender string
//    Name   string
//}

//Class 班级
//type Class struct {
//    Title    string
//    Students []*Student
//}

//func main() {
//    c := &Class{
//        Title:    "101",
//        Students: make([]*Student, 0, 200),
//    }
//    for i := 0; i < 10; i++ {
//        stu := &Student{
//            Name:   fmt.Sprintf("stu%02d", i),
//            Gender: "男",
//            ID:     i,
//        }
//        c.Students = append(c.Students, stu)
//    }
//fmt.Println(c)
//JSON序列化：结构体-->JSON格式的字符串
//data, err := json.Marshal(c)
//if err != nil {
//    fmt.Println("json marshal failed")
//    return
//}
//fmt.Printf("json:%s\n", data)
//}
