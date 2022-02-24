package main

import "fmt"

// 继承，改写父类
type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodo struct {
	Base
}

func (Voodo) Magic() {
	fmt.Println("Voodo magic")
}

func main() {
	v := new(Voodo)
	v.Magic()
	v.MoreMagic()
}
