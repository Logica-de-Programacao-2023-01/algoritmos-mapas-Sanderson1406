package main

import (
	"fmt"
	"strings"
)

func occo(s string) map[string]int {
	letra := strings.Split(s, " ")
	ocorrencias := make(map[string]int)
	for _, palavra := range letra {
		ocorrencias[palavra]++
	}
	return ocorrencias
}

func main() {
	s := "Eu fui com lucas, mas eu tambem fui com pedro, mas é mentira"
	r := occo(s)
	fmt.Println(r)

}
