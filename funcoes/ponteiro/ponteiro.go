package main

import "fmt"

func inc1(n int) {
	n++
}

// ponteiro
func inc2(n *int) {
	// desreferencia o ponteiro
	*n++
}

func main() {
	n := 1
	inc1(n) // passagem por valr
	fmt.Println(n)

	inc2(&n) // passagem por referencia
	fmt.Println(n)
}
