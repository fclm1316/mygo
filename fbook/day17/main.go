package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("欢迎来到信息管理系统")
	fmt.Println("1 添加学员")
	fmt.Println("2 编辑学员")
	fmt.Println("3 删除学员")
	fmt.Println("4 查询学员")
	fmt.Println("5 退出系统")
}

func getInput() *student {
	var (
		id    int
		name  string
		class string
	)

	fmt.Println("请输入学员信息")
	fmt.Println("请输入学员学号 : ")
	fmt.Scanf("%d\n", &id)
	fmt.Println("请输入学员名字 : ")
	fmt.Scanf("%s\n", &name)
	fmt.Println("请输入学员班级 : ")
	fmt.Scanf("%s\n", &class)
	stu := newStudent(id, name, class) //调用构造函数，返回的是指针
	//fmt.Println(*stu)
	return stu
}

func main() {
	sm := newStudentMgr()
	for {
		showMenu()
		var inputNum int
		fmt.Print("请输入需要的序列号: ")
		fmt.Scanf("%d\n", &inputNum)
		fmt.Printf("用户输入的是： %d \n", inputNum)

		switch inputNum {
		case 1:
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			stu := getInput()
			sm.modfiyStudent(stu)
		case 3:
			stu := getInput()
			sm.delStudent(stu)
		case 4:
			sm.showStudent()
		case 5:
			os.Exit(0)
		}

	}

}
