package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int // slice com array de origin de tamanho zero

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)

	// copia os elementos de um slice para o outro porém respeita o tamanho maximo do destinatario (slice2)
	// não mantém o msm array de origem
	copy(slice2, slice1)
	fmt.Println(slice2)
}
