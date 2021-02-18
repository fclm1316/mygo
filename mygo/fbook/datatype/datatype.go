package main

import (
	"fmt"
	"math"
	"strings"
)

//定义返回两个值
func foo() (string, int) {
	return "alex", 9000
}

func changstring() {
	//修改字符串，必须先转换成 []rune []byte , 单个修改 字符 ,重新分配内存
	s := "big"
	byte1 := []byte(s)
	byte1[0] = 'p'
	fmt.Println(string(byte1))
	d := "中文"
	rune1 := []rune(d)
	rune1[0] = '英'
	fmt.Println(string(rune1))
}

func sqrtDemo() {
	//强制转换字符串
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {

	var name string
	var age int
	age = 2
	name = "猿"
	fmt.Println(name)
	fmt.Println(age)
	var (
		a string
		b int
		c bool
		d string
	)
	a = "猴子"
	b = 100
	c = true
	d = "100"
	fmt.Println(a, b, c, d)

	var x string = "上蹿下跳"
	fmt.Println(x)
	fmt.Printf("%s 占位打印 %d\n", x, b)

	var y = 200
	var u = "中文"
	fmt.Println(y)
	fmt.Println(u)

	//简短变量声明（只能在函数内部）
	tt := 100
	fmt.Println(tt)

	//调用foo函数
	//用于接收不需要的值,匿名变量
	aa, _ := foo()
	fmt.Println(aa)

	var aa1 int = 10
	var bb1 int = 077
	var cc1 int = 0xff
	fmt.Println(a, b, c)
	fmt.Printf("%b\n", aa1)
	fmt.Printf("%o\n", bb1)
	fmt.Printf("%x\n", cc1)
	//内存地址
	fmt.Printf("%p\n", &c)
	fmt.Println(math.MaxFloat64)

	fmt.Println("c:\\go")

	var s1 string = "啊啊啊"
	var s2 string = `
	啊啊啊  “” 、\

	`
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(len(s2))

	fmt.Println(s1 + s2)
	s3 := fmt.Sprintf("%s-----%s", s1, s2)
	fmt.Println(s3)

	s4 := "aasd,asdddd,asssss"
	//分割
	ret1 := strings.Split(s4, ",")
	fmt.Println(ret1)
	//包含
	ret2 := strings.Contains(s4, "asss")
	fmt.Println(ret2)
	//前缀后缀
	ret3 := strings.HasPrefix(s4, "aa")
	ret4 := strings.HasSuffix(s4, "sss")
	fmt.Println(ret3)
	fmt.Println(ret4)
	//求字串位置
	fmt.Println(strings.Index(s4, "s"))
	fmt.Println(strings.LastIndex(s4, "s"))
	//join
	a1 := []string{"Python", "Java", "JavaScript", "Golang"}
	fmt.Println(strings.Join(a1, " + "))
	changstring()
	sqrtDemo()

}
