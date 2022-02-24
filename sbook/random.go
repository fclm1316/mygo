package main

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

func init() {
	// 初始化随机种子,不然每次运行获得一样的值,一般以时间为种子
	rand.Seed(int64(time.Now().Nanosecond()))
}

func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s %d ", file, line)
	}
	for i := 0; i < 2; i++ {
		a := rand.Int()
		fmt.Println(a)
	}

	for i := 0; i < 2; i++ {
		b := rand.Intn(20)
		fmt.Println(b)
	}
	fmt.Println()
	//timens := int64(time.Now().Nanosecond())
	timens := int64(time.Now().Nanosecond())
	// 反射获得类型
	fmt.Println(reflect.TypeOf(timens))
	where()
	fmt.Println(timens)
}
