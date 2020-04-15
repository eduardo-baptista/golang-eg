package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

// composiçao não erança
type ferrari struct {
	carro       // campos anonimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	// é possivel acessar os campos de carro diretamente por ferrari
	// porem a estrutura se matem no formato de composição
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s esta com o tubo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f)
}
