package main

import "fmt"

func array1() {
	// 初始化
	var testArray1 [3]int
	// 赋值
	var testArray2 = [3]int{1, 2}
	testArray3 := [3]string{"杭州", "北京"}
	// 根据值来确定长度
	testArray4 := [...]string{"杭州", "北京"}
	// 切片赋值
	testArray5 := [5]int{1: 1, 3: 5}
	fmt.Println(testArray1)
	fmt.Println(testArray2)
	fmt.Println(testArray3)
	fmt.Println(testArray4)
	fmt.Println(testArray5)
}

func array2() {
	testArray6 := [...]string{"双儿", "沐剑萍", "方怡", "苏荃", "曾柔"}
	for i := 0; i < len(testArray6); i++ {
		fmt.Println(testArray6[i])
	}

	for index, value := range testArray6 {
		fmt.Printf("%d --> %s \n", index, value)
	}
}

func arrary3() {
	//第一层可用推导式
	testArray7 := [...][3]string{
		//	testArray7 := [3][...]string {
		{"北京1", "上海1", "杭州1"},
		{"北京2", "上海2", "杭州2"},
		{"北京3", "上海3", "杭州3"},
	}
	//fmt.Println(testArray7)
	//fmt.Println(testArray7[0])
	//fmt.Println(testArray7[1][0])
	for _, value1 := range testArray7 {
		//fmt.Println(value1)
		for _, value2 := range value1 {
			fmt.Printf("---> %s \n", value2)
		}
	}
}

func modifyArray(x [3]int) {
	//函数接收数组
	x[0] = 100
	fmt.Println(x)
}

func modifArray2(x [3][2]int) [3][2]int {
	x[2][1] = 100
	return x
}

func main() {
	//array1()
	//array2()
	//arrary3()
	a := [3]int{1, 2, 3}
	modifyArray(a)
	b := [3][2]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	c := modifArray2(b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
