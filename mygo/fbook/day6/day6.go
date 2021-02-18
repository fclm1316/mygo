package main

import "fmt"

func slice1() {
	a := [5]int{1, 2, 3, 4, 5}
	// s := a[2:4:5]  a[low:high:max_low]
	//切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。
	// s:= a[2:4]
	s := a[2:4:5]
	fmt.Println(s)
	//cap() 容量
	fmt.Println(len(s), cap(s))
	//构造b数组，使用2容量，默认0，最大10
	b := make([]int, 2, 10)
	fmt.Println(b, len(b), cap(b))
	//    是否为空 只能使用 len(s) == 0 ,不能 s == nil
	for index, value := range a {
		fmt.Printf("%d ---> %d \n", index, value)
	}
}

func slice2() {
	//两个变量使用同一个底层数组，改变一个对另一个影响
	s0 := []int{1, 2, 3, 4} //slice  切片类型 , 上下不是同一个类型
	//s0 := [4]int{1, 2, 3, 4}  //定长数组类型,可切片
	//s0 := [...]int{1, 2, 3, 4}  //定长数组类型,通过长度倒推,可切片
	s1 := make([]int, 3)
	s2 := s1
	s2[0] = 100
	//s4 必须用make 定义切片,才能copy 不能 var s4 []int
	//有值无法拷贝
	//var s4 []int
	s4 := make([]int, 4)
	copy(s4, s0)
	fmt.Println(s1)
	fmt.Println(s4)
}

func appendSlice() {
	//通过 var 定义，可以直接append
	var numSlice []int
	for i := 0; i < 20; i++ {
		numSlice = append(numSlice, i)
		// 占位符 %p 指针  %v 相应的默认格式
		fmt.Printf("%p , %d,%d,%v\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
	a := []string{"智利", "冰岛"}
	var stringSlice []string
	stringSlice = append(stringSlice, "中国", "西班牙")
	// a...
	stringSlice = append(stringSlice, a...)
	fmt.Println(stringSlice)

}

func findAdd(x []int) {
	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			//fmt.Println(i, j)
			k := x[i] + x[j]
			for _, value := range x {
				if k == value {
					fmt.Println(i, j)
				}
			}
		}
	}

}

func main() {
	a := []int{1, 3, 5, 7, 8}
	findAdd(a)
	//slice1()
	//slice2()
	//appendSlice()
}
