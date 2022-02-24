package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func godefer() {

	log.Println("godefer")

}

func testdefer() {
	for i := 0; i <= 100; i++ {
		log.Printf("协程")
		time.Sleep(1e9)
	}
}
func main() {
	log.Println("aaaa")
	go testdefer()
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT)
	func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.Println("close chan")
				os.Exit(0)
			default:
			}
		}
	}()
}
