package main

import (
	"fmt"
	"time"
)

// 两个时间相差的纳秒数
var week1 time.Duration
var week2 time.Duration

func main() {

	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%04d", t.Day(), t.Month(), t.Year())
	fmt.Println(t.UTC())
	// 使用Add 正负数来时间加减，必须纳秒格式
	week1 = 60 * 60 * 24 * 7 * 1e9
	weekFormNow := t.Add(week1)
	week2 = -60 * 60 * 24 * 7 * 1e9
	weekPastNow := t.Add(week2)
	fmt.Println(weekFormNow)
	fmt.Println(weekPastNow)
	fmt.Println(t.AddDate(1, 1, -1))

	m1, _ := time.ParseDuration("1h")
	m2, _ := time.ParseDuration("-1m")
	m3, _ := time.ParseDuration("-1s")
	// "ns", "us" (or "µs"), "ms", "s", "m", "h".
	time1 := t.Add(m1)
	time2 := t.Add(m2)
	time3 := t.Add(m3)
	fmt.Println(time1)
	fmt.Println(time2)
	fmt.Println(time3)
	// Sub 用来计算两个时间差
	t2 := time.Now()
	subM := t2.Sub(t)
	fmt.Println(subM.Milliseconds(), "毫秒")
	// 时间差
	fmt.Println(time.Since(t))
}
