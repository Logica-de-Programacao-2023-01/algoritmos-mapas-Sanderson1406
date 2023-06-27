package main

import "fmt"

func fibonatti(n int) map[int]int {
	fibMap := make(map[int]int)
	fibMap[0] = 0
	fibMap[1] = 1
	for i := 2; i <= n; i++ {
		fibMap[i] = fibMap[i-1] + fibMap[i-2]
	}
	return fibMap
}

func main() {
	n := 10
	fibSequence := fibonatti(n)
	for index, number := range fibSequence {
		fmt.Printf("Index: %d, Number: %d\n", index, number)
	}
}
