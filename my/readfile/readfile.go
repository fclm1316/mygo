package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile() {
	fileOjb, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed ,err:%v", err)
		return
	}
	defer fileOjb.Close()
	//var tmp = make([]byte,128)
	var tmp [128]byte
	for {
		n, err := fileOjb.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完")
			return
		}
		if err == nil {
			fmt.Println("open filed ,err:$v", err)
			return

		}

		fmt.Printf("读取%d个字节", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

func readFromFileByIO() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open filed ,err:$v", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件完成")
			break
		}
		if err != nil {
			fmt.Print("open filed ,err:$s", err)
			return
		}
		fmt.Println(line)
	}

}

func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("open filed ,err:$v", err)
		return
	}
	fmt.Println(string(ret))

}

type Jsonone struct {
	Name string
	Age  int
	From Jsontwo
}

type Jsontwo struct {
	City   string
	Number int
	Like   [4]string
}

func main() {
	//jsonFile, err := os.Open("./test.json")
	//if err != nil {
	//    fmt.Println("open filed ,err:$v", err)
	//    return
	//}
	//defer jsonFile.Close()
	//byteValue, _ := ioutil.ReadAll(jsonFile)
	//var Myjson Jsonone
	//json.Unmarshal([]byte(byteValue),&Myjson)
	//fmt.Println(Myjson.From.Like[1])
	jsonFile, err := ioutil.ReadFile("./test.json")
	if err != nil {
		fmt.Println("open filed ,err:$v", err)
		return
	}
	var Myjson Jsonone
	_ = json.Unmarshal([]byte(jsonFile), &Myjson)
	fmt.Println(Myjson.From.Like[0])
}

func useScan() {
	var s string
	fmt.Print("请输入内容:")
	// 无法处理空格输入
	fmt.Scanln(&s)
	fmt.Printf("输入的内容是 %s", s)
}

func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容:")
	// 处理空格输入
	s, _ = reader.ReadString('\n')
	fmt.Printf("输入的内容是 %s", s)

}
