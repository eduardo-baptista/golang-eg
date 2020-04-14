package main

import (
	"fmt"
	"time"
)

func main() {
	// variavel externa
	i := 1
	for i <= 10 { /// não tem while em Go
		fmt.Println(i)
		i++
	}

	// inicializaçao no laço
	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	fmt.Println("\nMisturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	// laço infinito
	for {
		fmt.Println("Para sempre ...")
		time.Sleep(time.Second)
	}

}
