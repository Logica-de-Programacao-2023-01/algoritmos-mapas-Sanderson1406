package main

import "fmt"

func divida(pessoas map[string]float64) map[string]float64 {
	divida := make(map[string]float64)
	var soma float64
	for _, valor := range pessoas {
		soma += valor
	}
	divi := soma / float64(len(pessoas))
	for pessoa, valo := range pessoas {
		if valo < soma {
			divida[pessoa] = valo - divi
		} else if valo > soma {
			divida[pessoa] = valo + divi
		}
	}
	return divida
}

func main() {
	pessoas := map[string]float64{
		"Julia":     1445.4,
		"Pedro":     600.4,
		"Sanderson": 975.8,
		"Maria":     1354.4,
		"Miguel":    1290.8,
	}
	r := divida(pessoas)
	fmt.Println(r)
}
