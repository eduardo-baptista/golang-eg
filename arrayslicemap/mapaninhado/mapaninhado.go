package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"E": {
			"Eduardo":  201231.22,
			"Edimundo": 1999.11,
		},
		"V": {
			"Victoria": 5000.00,
		},
		"P": {
			"Pedro Junior": 1200.00,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
