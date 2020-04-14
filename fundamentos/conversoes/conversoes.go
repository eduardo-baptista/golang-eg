package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	// para realizar operações é necessario ter as variaveis no msm tipo
	fmt.Println(x / float64(y))

	nota := 6.9
	// o arredondamento apenas ignora as casas decimais
	notaFinal := int(nota) // output: 6
	fmt.Println(notaFinal)

	// cuidado
	fmt.Println("Test " + string(97)) // não converte em string, retorna os caracteres da tabela unicode

	// int para string
	fmt.Println("Test " + strconv.Itoa(123)) // ou fmt.Sprint(123)

	// string para int
	num, _ := strconv.Atoi("123") // a função retorna um numero e um erro caso nao seja possivel converter
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}

}
