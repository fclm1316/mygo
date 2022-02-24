package main

import (
	"fmt"
	"math"
)

//子类 继承
type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point
	name string
}

//覆写方法，父类，母类，多重继承
func (n *NamedPoint) Abs() float64 {
	return n.Point.Abs() * 100
}

func main() {
	//n := &NamedPoint{Point{3,4},"Python"} //初始化
	//fmt.Println(n.Abs())
	n := new(NamedPoint)
	n.x = 3
	n.y = 4
	n.name = "Python"
	fmt.Println(n.Abs())
}
