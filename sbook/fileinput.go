package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFile, inputError := os.Open("d:/aaa.txt") // 打开文件句柄

	if inputError != nil { // 打开有错，返回结束
		fmt.Printf("An error occurred on opening the inputfile")
		return
	}
	defer inputFile.Close()                   // 最后关闭
	inputReader := bufio.NewReader(inputFile) //缓冲区
	for {                                     // 无线循环
		inputString, readerError := inputReader.ReadString('\n') // 读取
		if readerError != nil {
			return
		}
		fmt.Printf("The input was : %s ", inputString)
	}
}
