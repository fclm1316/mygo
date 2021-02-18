package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "6666"
	var an int
	var newS string

	fmt.Printf("OS is %d \n", strconv.IntSize)

	// 字符串转数字 , 字符串转其他时不一定成功
	an, _ = strconv.Atoi(orig)
	// 其他转字符串必定成功
	newS = strconv.Itoa(an)
	fmt.Printf("%s", newS)
}
