package main

import "fmt"

//student 结构体
type student struct {
	id    int
	name  string
	class string
}

//student 构造函数
func newStudent(id int, name string, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

// 管理学生类结构体,保存所有student
type studentMgr struct {
	allStudents []*student
}

// 构造函数
func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

// 给newStudentmgr 添加方法,添加新学员
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}
func (s *studentMgr) delStudent(newStu *student) {
	for i, v := range s.allStudents { //循环id
		if newStu.id == v.id { // id一致 修改
			// 切片，把符合的排除
			s.allStudents = append(s.allStudents[:i], s.allStudents[i+1:]...) //元素被打散一个个append进
			return
		}
	}
	// 没有找到
	fmt.Printf("没有 %d 的学生 \n", newStu.id)
}

// 编辑学员
func (s *studentMgr) modfiyStudent(newStu *student) {
	for i, v := range s.allStudents { //循环id
		if newStu.id == v.id { // id一致 修改
			s.allStudents[i] = newStu
			return
		}
	}
	// 没有找到
	fmt.Printf("没有 %d 的学生 \n", newStu.id)
}
func (s *studentMgr) showStudent() {
	fmt.Println(s.allStudents)
	for _, v := range s.allStudents {
		fmt.Printf("学号：%d 姓名：%s 班级：%s \n", v.id, v.name, v.class)
	}
}
