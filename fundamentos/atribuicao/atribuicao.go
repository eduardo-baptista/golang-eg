package main

import "fmt"

func main() {
	// atribuição simples
	var b byte = 3
	fmt.Println(b)

	i := 3 // inferência de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2
	fmt.Println(i)

	x, y := 1, 2 // multiplos valores atribuidos
	fmt.Println(x, y)

	x, y = y, x // inverte os valores das variaveis
	fmt.Println(x, y)
}
