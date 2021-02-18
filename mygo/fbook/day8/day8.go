package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func mapX() {
	// map 类型的变量默认初始值为nil 需要使用make()函数来分配内存
	// make(map[KeyType]ValueType,[cap])
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"])
	fmt.Printf("type of a %T \n", scoreMap)
	//没有值时 make() ,一开始定义时不需要
	scoreMap2 := map[string]string{
		"username": "流量",
		"password": "组织",
	}
	fmt.Println(scoreMap2["username"])
	for k, v := range scoreMap2 {
		fmt.Println(k, v)
	}
	delete(scoreMap2, "")
	//是否在字典中
	value, ok := scoreMap2["username"]
	if ok {
		fmt.Printf("ok %s", value)
	} else {
		fmt.Printf("faile")

	}
}

func sortMap() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	//定义map
	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		// %02d 不够两位，前补0
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		// map 中的顺序和添加顺序无关
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	//定义切片，申请内存
	//var keys = make([]string, 0, 200)
	var keys []string
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

func silceMap() {
	//除非直接初始化，不然make
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

func sliceMap2() {
	// 创建字符串map 里面是切片
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	key := "中国"
	// 判断 value 是否在 map里面
	value, ok := sliceMap[key]
	if !ok {
		//创建切片
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

//func main() {
//mapX()
//sortMap()
//silceMap()
//sliceMap2()
//}

func main() {
	words := "how do you do"
	splits := strings.Split(words, " ")
	result := make(map[string]int, 8)
	for _, v := range splits {
		fmt.Println(v)
		result[v] = result[v] + 1
	}
	fmt.Println(result)
}
