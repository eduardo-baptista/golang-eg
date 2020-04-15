package main

import "fmt"

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	useID int
	itens []item
}

func (p pedido) valoTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}

	return total
}

func main() {
	pedido := pedido{
		useID: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 1, 23.19},
			item{11, 100, 3.13},
		},
	}

	fmt.Printf("Valor total do pedido é R$%.2f", pedido.valoTotal())
}
