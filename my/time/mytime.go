package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Unix())
	fmt.Println(now.Nanosecond())
	// 转换时间
	Unix := now.Unix()
	fmt.Println(time.Unix(Unix, 0))
	// 格式化时间  2006 1  2  3  4   5
	//            年  月  日 时  分  秒
	fmt.Println(now.Format("01/02 03:04:05 2006"))
	fmt.Println(now.Format("01/02 03:04:05.000 2006"))
}
