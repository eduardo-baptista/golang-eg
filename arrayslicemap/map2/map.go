package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.00,
	}

	// adicionando
	funcsESalarios["Rafael Filho"] = 1350.00

	// ao tentar deletar um registro que nao consta a operação é ignorada sem gerar erros
	delete(funcsESalarios, "Nao conheco")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
