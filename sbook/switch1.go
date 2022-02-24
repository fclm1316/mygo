package main

import "fmt"

func main() {
	var num1 int = 90

	switch num1 {
	case 90, 91:
		fmt.Println("not 90,91: ", num1)
	case 100:
		fmt.Println("is 100:", num1)
	default:
		fmt.Println("what happend")
	}

	switch {
	case num1 < 100:
		fmt.Println("num1 < 100")
	case num1 == 100:
		fmt.Println("num1 == 100")
	case num1 > 100:
		fmt.Println("num1 > 100")

	}

	/*    switch result := Daboluo() {
	      case result < 0:
	          fmt.Println("aaa")
	      default:
	          fmt.Println("bbb")
	      }
	*/
}
