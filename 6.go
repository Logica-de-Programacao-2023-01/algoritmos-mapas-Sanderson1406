package main

import (
	"fmt"
	"strings"
)

func fina(r map[string]int, r2 map[string]int, r3 map[string]int) map[string]int {
	saida := make(map[string]int)
	for chave := range r {
		saida[chave]++
	}
	for chave2 := range r2 {
		saida[chave2]++
	}
	for chave3 := range r3 {
		saida[chave3]++
	}
	return saida
}

func cont(s string) map[string]int {
	carac := strings.Split(s, " ")
	ocorrencias1 := make(map[string]int)
	for _, cara := range carac {
		ocorrencias1[cara]++
	}
	return ocorrencias1
}

func cont2(s2 string) map[string]int {
	carac := strings.Split(s2, " ")
	ocorrencias2 := make(map[string]int)
	for _, cara := range carac {
		ocorrencias2[cara]++
	}
	return ocorrencias2
}
func cont3(s3 string) map[string]int {
	carac := strings.Split(s3, " ")
	ocorrencias3 := make(map[string]int)
	for _, cara := range carac {
		ocorrencias3[cara]++
	}
	return ocorrencias3
}

func main() {
	s := "eu sou um"
	s2 := "eu sou"
	s3 := "eu"
	r := cont(s)
	r2 := cont2(s2)
	r3 := cont3(s3)
	ss := fina(r, r2, r3)
	fmt.Println(ss)
}
