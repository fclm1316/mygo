package main

import (
	"fmt"
	"strings"
)

func main() {
	//字符串开头结尾
	var str string = " ThiS IS An, ExaMple Of, A STring "
	// 头
	fmt.Printf("%t \n", strings.HasPrefix(str, "Th"))
	// 尾
	fmt.Printf("%t \n", strings.HasPrefix(str, "int"))
	// 包含
	fmt.Printf("%t \n", strings.Contains(str, "exa"))
	// 次数
	fmt.Printf("%d \n", strings.Count(str, "a"))
	// 相等
	fmt.Printf("%d \n", strings.Compare("a", "e"))
	// 位置
	fmt.Printf("%d \n", strings.Index(str, "e"))
	// 位置 , 不存在 -1
	fmt.Printf("%d \n", strings.Index(str, "oo"))
	// 最后位置
	fmt.Printf("%d \n", strings.LastIndex(str, "e"))
	// 替换 , -1 表示所有，n 前几个
	fmt.Printf("%s \n", strings.Replace(str, "e", "-----", -1))
	// 重复
	newStr := strings.Repeat(str, 2)
	fmt.Printf("%s \n", newStr)
	// 大小写
	lower := strings.ToLower(str)
	upper := strings.ToUpper(str)
	title := strings.ToTitle(str)
	fmt.Printf("%s \n %s \n %s \n", lower, upper, title)
	// 头尾空格
	fmt.Println("--------------------------")
	fmt.Printf("%s\n", strings.TrimSpace(str))
	// 头尾 i
	fmt.Printf("%s\n", strings.Trim(str, "i"))
	// 左右
	fmt.Printf("%s\n %s\n", strings.TrimRight(str, "ing "), strings.TrimLeft(str, " Th"))
	// 空格分割，切片
	for i, v := range strings.Fields(str) {
		fmt.Printf("%s --> ", strings.Fields(str)[i])
		fmt.Println(v)
	}
	fmt.Printf("%s \n", strings.Fields(str)[0])
	// 指定格式，切片
	fmt.Printf("%s \n", strings.Split(str, ",")[1])

	str2 := strings.Fields(str)
	/* 字符串拼接*/ str3 := strings.Join(str2, ";")
	fmt.Println(str3)

}
