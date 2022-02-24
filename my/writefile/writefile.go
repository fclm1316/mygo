package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	//WriteFile1()
	//BufIoWrite()
	IoutilWrite()

}

func WriteFile1() {
	//fileOjb, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	fileOjb, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file err : %v", err)
		return
	}
	defer fileOjb.Close()
	// write
	n, _ := fileOjb.Write([]byte("bbbbb\n"))
	fmt.Println(n)
	// wirtestring
	fileOjb.WriteString("aaaaa")
}

func BufIoWrite() {
	file, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file err : %v", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		//字符串拼接
		str1 := fmt.Sprintf("%d : 我是谁\n", i)
		writer.WriteString(str1)
	}
	writer.Flush() //缓存中的内容写入

}

func IoutilWrite() {
	str1 := "我在哪里\n"
	// buf 写文件,字符串拼接
	var str2 bytes.Buffer
	for i := 0; i < 10; i++ {
		//转换成字符串
		str2.WriteString(strconv.Itoa(i))
		str2.WriteString(" : ")
		str2.WriteString(str1)
	}
	ioutil.WriteFile("./xx.txt", []byte(str2.String()), 0644)
	fmt.Println(str2.String())
}
