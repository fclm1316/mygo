package main

import "fmt"

//枚举 ,多一行 iota 加1
const (
	a1 = iota
	a2 = iota
	a3 = iota
	a4 = 100
	a5 = iota
)

// const(
// 	_ = iota
// 	KB = 1 << (10 * iota) // 1<<10 移位运算 2的10次方 1024
// 	MB = 1 << (10 * iota) // 1<< 20
// 	GB = 1 << (10 * iota)
// 	TB = 1 << (10 * iota)
// 	PB = 1 << (10 * iota)
// )

// const (
// 	a, b = iota + 1 ,iota + 2  //iota = 0 , a =1 b =2
// 	c, d 						// iota =1 ,c = 2 d = 3
// 	e, f 						// iota = 2 , e = 3 f = 4
// )

func main() {
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)

}
