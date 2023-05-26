package main

import (
	"fmt"
	"strings"
)

func frase(s string) map[string]map[rune]int {
	saida := make(map[string]map[rune]int)
	palavras := strings.Fields(s)
	for _, palavra := range palavras {
		letras := make(map[rune]int)
		for _, letra := range palavra {
			letras[letra]++
		}
		saida[palavra] = letras
	}
	return saida
}

func main() {
	s := "zeou a vida mas ta tudo bem a rainha morreu, morrou mas passa bem"
	r := frase(s)
	for palavra, letras := range r {
		fmt.Println(palavra, letras)
	}
}
