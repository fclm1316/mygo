package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := ""
	fmt.Println(len(os.Args))
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ") // 0 是本身
	} else {
		fmt.Println(os.Args[0])
		return
	}
	fmt.Printf("good morning %s", who)
}
