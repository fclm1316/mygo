package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I love my work and I love my family too I I "
	mRet := wordCountFunc(str)

	for k, v := range mRet {
		fmt.Printf("%q:%d\n", k, v)
	}
}

func wordCountFunc(str string) map[string]int {
	s := strings.Fields(str)  // 字符串分割，默认空格
	m := make(map[string]int) // 创建切片类型
	for i := 0; i < len(s); i++ {
		//fmt.Println(s[i]) // key
		if _, ok := m[s[i]]; ok { //判断m切片的key，value 是否存在
			m[s[i]] = m[s[i]] + 1
		} else {
			m[s[i]] = 1
		}
	}
	return m
}
