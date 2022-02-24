package main

import (
	"fmt"
)

func main() {
	x := min(2, 3, 4, 0)
	fmt.Printf("The minimum is : %d\n", x)
	arr := []int{7, 3, 5, 9, 10, 22}
	y := min(arr...) //传递边长参数
	fmt.Printf("the minmum arr is %d\n", y)
}

func min(a ...int) int { //接收变长参数
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v <= min {
			min = v
		}
	}
	return min
}
