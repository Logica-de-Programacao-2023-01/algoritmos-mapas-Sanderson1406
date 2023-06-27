package main

import "fmt"

func uau(numbers []int) map[[2]int]int {
	frequencyMap := make(map[[2]int]int)
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			pair := [2]int{numbers[i], numbers[j]}
			frequencyMap[pair]++
		}
	}
	return frequencyMap
}

func main() {
	numbers := []int{1, 2, 3, 1, 2, 4, 5, 3, 2, 1, 4}
	pairsFrequency := uau(numbers)
	for pair, frequency := range pairsFrequency {
		fmt.Printf("Pair: %v, Frequency: %d\n", pair, frequency)
	}
}
