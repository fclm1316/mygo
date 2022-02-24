package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{
		"alpha": 34, "bravo": 20, "charlie": 23,
		"delta": 87, "echo": 22, "foxtror": 78}
)

func main() {

	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	fmt.Printf("sortd:\n")
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("key : %s , value : %d\n", k, barVal[k])
	}
}
