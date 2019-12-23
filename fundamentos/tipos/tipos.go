package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(34500))

	// positivos - uint8 uint16 uint32 uint64
	var b byte = 255 // byte é um alias para uint8
	fmt.Println("Byte é", reflect.TypeOf(b))

	fmt.Println("Max inteiro", math.MaxInt64)

	var c rune = 'a' // representa um mapeamento da tabela unicode (int32)
	fmt.Println("Rune é", reflect.TypeOf(c))

	// números reais - float32, float64
	var d float32 = 13.12
	fmt.Println("O tipo de D é:", reflect.TypeOf(d))
	fmt.Println("O tipo do literal 67.45 é:", reflect.TypeOf(67.45))
	// for default sempre é float64, caso precise do 32 é necessário declarar explicitamente

	// boolean
	bo := true
	fmt.Println("O tipo de BO é:", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Bile, meu jaguar"
	fmt.Println(s1 + "!")
	fmt.Println("Tamanho da frase:", len(s1))

	s2 := `Quem
	pediu
	mato grosso
	e
	mathias
	?`
	fmt.Println(s2)
	fmt.Println("Tamanho da frase:", len(s2))
}
