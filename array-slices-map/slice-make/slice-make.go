package main

import "fmt"

func main() {
	// make cria um slice que aponta para um array interno de 10 posições
	slice := make([]int, 10)
	slice[9] = 13
	fmt.Println(slice)

	slice = make([]int, 10, 20)
	fmt.Printf("slice = %d - tamanho slice = %d - capacidade array = %d", slice, len(slice), cap(slice))

	slice = append(slice, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("\nslice = %d - tamanho slice = %d - capacidade array = %d", slice, len(slice), cap(slice))

	// ao atribuir um valor acima da capacidade o slice cria um novo array e continua referenciando
	slice = append(slice, 1)
	fmt.Printf("\nslice = %d - tamanho slice = %d - capacidade array = %d", slice, len(slice), cap(slice))
}
