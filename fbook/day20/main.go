package main

import "fmt"

// 两种方式在接收值的时候判断该通道是否被关闭
// channel 练习
func main() {
	//定义阻塞试通道,无缓冲,必须要有接收者,同步通道
	//var ch1 chan int
	//ch := make(chan int, 1) 有缓冲通道
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~10的数发送到ch1中
	go func() { //匿名函数
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			// 第一种
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	// 第二种，使用较多
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}
