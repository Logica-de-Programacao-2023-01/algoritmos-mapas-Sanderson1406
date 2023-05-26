package main

import "fmt"

func addmapa(frutas2 map[string]int, frutas map[string]int) map[string]int {
	novaFrutas := make(map[string]int)
	novaFrutas["laranja"] = frutas["laranja"]
	novaFrutas["uva"] = frutas["uva"]
	novaFrutas["banana"] = frutas["banana"]
	novaFrutas["tomate"] = frutas["tomate"]
	novaFrutas["pera"] = frutas2["pera"]
	novaFrutas["abacate"] = frutas2["abacate"]
	novaFrutas["pitanga"] = frutas2["pitanga"]
	return novaFrutas
}

func main() {
	frutas := map[string]int{
		"laranja": 50,
		"uva":     48,
		"banana":  47,
		"tomate":  55,
	}
	frutas2 := map[string]int{
		"pera":    54,
		"abacate": 48,
		"pitanga": 47,
	}
	r := addmapa(frutas2, frutas)
	fmt.Println(r)
}
