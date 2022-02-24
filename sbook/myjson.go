package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Address struct {
	Type     string
	City     string
	Countory string
}

type VCard struct {
	Firstname string
	Lastname  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Afjj", "Bel"}
	wa := &Address{"work", "sh", "BJ"}
	vc := VCard{"jan", "jdd", []*Address{pa, wa}, "sz"}
	js, _ := json.Marshal(vc)
	fmt.Printf("%s", js)

	v := []byte(string(js))
	var d VCard
	json.Unmarshal(v, &d)
	fmt.Printf("\n%v\n", d.Addresses)         // 内存地址
	fmt.Printf("\n%v\n", d.Addresses[0].City) //
	fmt.Printf("\n%v\n", d.Firstname)         //
	fmt.Printf("\n%v\n", reflect.TypeOf(d.Addresses[0]))
	for i := 0; i < len(d.Addresses); i++ {
		fmt.Printf("----------> %s", d.Addresses[i].Countory)
	}
	//for i := 0; i <= len(d.Addresses); i++ {
	//    for k1, v1 := d.Addresses[i] {
	//        fmt.Printf("k:%v,v:%v", k1, v1)
	//   }
	//}
	fmt.Println("\n--------------------------------\n")

	// 空接口获取未知格式的json
	var f interface{}
	json.Unmarshal(v, &f)
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Printf("%s is string:%v\n", k, vv)
		case int:
			fmt.Printf("%s is int:%v\n", k, vv)
		case []interface{}:
			fmt.Println("an array")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		}
	}

}
