package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	// transforma number em string
	xs := fmt.Sprint(x)
	// ambos os valores precisam ser strings
	fmt.Println("O valor de x é " + xs)
	// contatena sem necessidade de converter
	fmt.Println("O valor de x é", x)
	// é possivel formatar os valores numericos
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	// por padrao o Printf mantem a mesma linha caso seja necessario sua quebra basta adicionar \n
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	// usando %v é possivel imprimir varios tipo de variaveis
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
