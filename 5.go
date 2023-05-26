package main

import (
	"fmt"
	"strings"
)

func conta(s string) map[string]int {
	carac := strings.Split(s, "")
	ocorrencias := make(map[string]int)
	for _, cara := range carac {
		ocorrencias[cara]++
	}
	return ocorrencias
}

func main() {
	s := "eu sou um"
	r := conta(s)
	fmt.Println(r)
}
