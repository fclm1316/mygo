package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"
)

var begin = time.Now()
var wg sync.WaitGroup
var IpChan = make(chan string, 5)

func main() {
	go GetIp()
	for data := range IpChan {
		wg.Add(1)
		go ScanPort(data)
	}
	wg.Wait()
	var end = time.Now().Sub(begin)
	fmt.Printf("耗时: %v\n", end)
}

func GetIp() {
	fileOjb, err := os.Open("myojb/scanport/nmap.txt")
	if err != nil {
		panic(err)
	}

	defer fileOjb.Close()

	reader := bufio.NewReader(fileOjb)

	for {
		line, _, err := reader.ReadLine()
		// 文件读完
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("open file err %s", err)
			return
		}
		GetIpPort := string(line) + ":21"
		IpChan <- GetIpPort
	}
	close(IpChan)
}

func ScanPort(ipport string) {
	defer wg.Done()
	connTime := 2 * time.Second
	conn, err := net.DialTimeout("tcp", ipport, connTime)
	if err != nil {
		fmt.Printf("%s is close \n", ipport)
		return
	}
	defer conn.Close()
	fmt.Printf("%s is open\n", ipport)
}
