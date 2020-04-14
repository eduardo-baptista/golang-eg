package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint21 uint64
	var b byte = 255 // byte é um alias para uint8
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int21 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println("O valor da variavel i2", i2)

	// números reais (float32, float64)
	var x float32 = 49.99 // or var x = float32(49.99)
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// strings
	s1 := "Olá neu nome é Nome"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho de string é", len(s1))

	// string com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Nomeee`
	fmt.Println("O tamanho da string é", len(s2))

	// char???
	// não a tipo char em GO o mesmo se refere a int32 para a referencia na tabela unicode
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}
