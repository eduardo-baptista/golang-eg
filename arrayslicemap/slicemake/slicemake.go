package main

import "fmt"

func main() {
	s := make([]int, 10) // cria um novo Slice, o array interno matem o mesmo tamanho do slice
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // cria um Slice de tamanho 10, o array interno terá tamanho 2o

	/*
		len = retorna tamanho atual do slice
		cap = retorna a capacidade do slice (tamanho do array interno)
	*/
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // limite do array interno
	fmt.Println(s)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 12) // ao atingir a capacidade do array é criado um novo array com o dobro da capacidade
	fmt.Println(s)
	fmt.Println(s, len(s), cap(s))

}
