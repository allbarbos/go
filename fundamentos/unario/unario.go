package main

import "fmt"

func main() {
	x := 1
	y := 2

	x++
	fmt.Println(x)

	y--
	fmt.Println(y)

	/*
		go não possuí:
			- operadores com precedência, ex.: --x ou ++y
			- operadores ternários
	*/
}
