package main

import "fmt"

// a funça=ão init é executada antes das demais funções, sem necessidade de ser chamada
// o método não precisa estar no mesmo arquivo da func main
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
