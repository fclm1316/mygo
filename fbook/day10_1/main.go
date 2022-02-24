package main

import "fmt"

// Address 构造结构体
type Address struct {
	//结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
	Province string
	City     string
	Create   string
}

// Address 构造结构体
type User struct {
	Name   string
	Gender string
	Create string
	// 结构体套嵌
	Address Address
}

// newPerson 构造函数，调用时方便
func newPerson(Name string, Gender string, Province string, City string, ACreate string, BCreate string) *User {
	return &User{
		Name:   Name,
		Gender: Gender,
		Create: ACreate,
		Address: Address{
			Province: Province,
			City:     City,
			Create:   BCreate,
		},
	}
}

func (p *User) setGender(newGender string) {
	p.Gender = newGender
}

func main() {
	p := newPerson("小鸡", "雄", "森林市", "三黄镇", "2021-01-26", "2020-02-02")
	fmt.Println(p)
	fmt.Printf("user1=%#v\n", p)
	p.setGender("雌")
	fmt.Println(p)
	p.Gender = "未知"
	fmt.Println(*p)
	c := newPerson("小鸡", "雄", "森林市", "三黄镇", "2021-01-26", "2020-02-02")
	fmt.Println(c)
	c.setGender("雌")
	fmt.Println(*p)

}
