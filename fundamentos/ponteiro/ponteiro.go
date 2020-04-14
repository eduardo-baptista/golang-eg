package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmética de ponteiros

	var p *int = nil

	p = &i // p recebe o endereço de memória de i

	*p++ // desreferenciando (pegando o valor)
	i++

	fmt.Println(p, *p, i, &i)
}
