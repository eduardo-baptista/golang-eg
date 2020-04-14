package main

import "fmt"

func main() {
	// dictionary like
	// maps devem ser inicializados
	// var aprovados map[int]string
	aprovados := make(map[int]string)

	aprovados[1234567891] = "Eu"
	aprovados[6161561561] = "Você"
	aprovados[1231231232] = "Ele"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	delete(aprovados, 1231231232)
	fmt.Println(aprovados)

}
