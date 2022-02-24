package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, outputError := os.OpenFile("d:/aaa.txt", os.O_WRONLY|os.O_CREATE, 0666)
	// 打开一个文件，只读或者创建
	if outputError != nil {
		// 有错误，退出
		fmt.Println("An error: ", outputError)
		return
	}
	defer outputFile.Close() // 关闭

	outputWrite := bufio.NewWriter(outputFile) // 缓冲器
	outputString := "hahahaha\n"
	for i := 0; i <= 10; i++ {
		outputWrite.WriteString(outputString) // 写入文件
	}
	outputWrite.Flush() //刷新
}
