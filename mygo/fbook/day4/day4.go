package main

import "fmt"

func ifDemo1() {
	source := 65
	if source >= 90 {
		fmt.Printf("大于 90 %d\n", source)
	} else if source > 75 {
		fmt.Printf("大于 75 %d\n", source)
	} else {
		fmt.Printf("大于 40 : %d\n", source)
	}

}

func forDemo1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forDemo2() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("a")
	case 2:
		fmt.Println("b")
	case 3:
		fmt.Println("c")
	case 4:
		fmt.Println("d")
	case 5, 6:
		fmt.Println("f")
	default:
		fmt.Println("H")

	}

}

func switchDemo2() {
	age := 44
	switch {
	case age < 20:
		fmt.Println("a")
	case age > 20 && age < 50:
		fmt.Println("b")
	case age > 50:
		fmt.Println("c")

	}

}

func gotoDemo1() {
	// breakFlage 在 i 循环体内
	var breakFlage bool
	for i := 0; i < 10; i++ {
		fmt.Println("-------")
		fmt.Println(i)
		for j := 0; j < 10; j++ {
			fmt.Println("+++++++")
			fmt.Println(j)
			if j == 5 {
				breakFlage = true
				// break
			}
		}
		if breakFlage {
			break
		}
	}

}

func gotoDemo2() {
	var i int
	var j int
	// i:= 只可以在循环体内使用
	for i = 0; i < 10; i++ {
		for j = 0; j < 10; j++ {
			// fmt.Printf("i = %d , j = %d \n", i,j)
			if i == 5 && j == 5 {
				goto breakTag
			}
			// fmt.Printf("i = %d , j = %d \n", i,j)
		}
		// breakTag:
		// fmt.Printf("i = %d , j = %d ", i, j)
	}
	// return
	// 标签
breakTag:
	fmt.Printf("i = %d , j = %d ", i, j)
	// fmt.Println("over")

}

func breakDemo() {
	// 打标签
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 5 && j == 5 {
				break BREAKDEMO1
			}
			fmt.Printf(" i = %d ,j = %d \n", i, j)
		}
	}

}

func continueDemo() {
forloop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 5 && j == 5 {
				fmt.Println("hahah")
				// 跳出当前的 for j ;继续 for i
				continue forloop
				fmt.Println("kuku")
			} else if i == 8 && j == 6 {
				fmt.Println("heihei")
				break forloop
			}
			fmt.Printf(" i = %d ,j = %d \n", i, j)
		}
	}
}

func main() {
	// ifDemo1()
	// forDemo1()
	// forDemo2()
	// switchDemo1()
	// switchDemo2()
	// gotoDemo1()
	// gotoDemo2()
	// breakDemo()
	continueDemo()

}
