package main

import "fmt"

func a() {
	fmt.Println("func a")
}

func b() {
	//注意搭配
	//z在出错前 defer 一个匿名函数
	defer func() {
		// 回魂
		err := recover()
		// 判断时候有错误值
		if err != nil {
			fmt.Println("func b error")
			fmt.Println(err)
		}
	}()
	// 异常退出
	panic("painc func b")
}

func c() {
	fmt.Println("func c")
}

func main() {
	a()
	b()
	c()
}
