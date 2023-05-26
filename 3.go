package main

import "fmt"

func soma(numeros map[string]int) int {
	var soma int
	for _, valor := range numeros {
		soma += valor
	}
	return soma
}

func main() {
	numeros := map[string]int{
		"a": 2,
		"b": 8,
		"c": 5,
		"d": 7,
		"e": 9,
	}
	r := soma(numeros)
	fmt.Println("A soma do mapa é:", r)
}
