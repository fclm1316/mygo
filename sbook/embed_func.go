package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func main() {
	c := new(Customer)
	c.Name = "batman"
	c.log = new(Log)
	c.log.msg = "wahaha"
	//c =&Customer{"superman",&Log{"kukuku"}}
	c.log.Add("flash")
	fmt.Println(c.log)
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string { //结构体带 String() 方法时，可以直接fmt打印
	return l.msg
}

func (c *Customer) String() string { //结构体带 String() 方法时，可以直接fmt打印
	return c.Name + "\nLog:" + fmt.Sprintln(c.log)
}
