package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "6.23 / 89 ? GO"
	format                 = "%f / %d ? %s"
)

func main() {
	fmt.Println("input string:")
	fmt.Scanln(&firstName, &lastName) //从键盘输入
	fmt.Printf("Hi %s %s!\n", firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s) //从字符串中输入,固定格式
	fmt.Println("From the string we read : ", f, i, s)

}
