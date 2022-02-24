package main

import "fmt"

type specialString string

var whatIsThis specialString = "hello"

func TypeSwitch() {
	testfunc := func(any interface{}) { //匿名函数
		switch v := any.(type) { //type-switch 用来检测类型
		case bool:
			fmt.Printf("any %v is a bool type", v)
		case int:
			fmt.Printf("any %v is a int type", v)
		case string:
			fmt.Printf("any %v is a string type", v)
		case specialString:
			fmt.Printf("any %v is a specialString type", v)
		}
	}
	testfunc(whatIsThis)
}

func main() {
	TypeSwitch()
}
