package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5, 6, 7} // ao utilizar ... o compilador conta a quantidade do array de acordo com o que foi declarado

	for indice, elemento := range numeros {
		fmt.Printf("indice %d - valor %d\n", indice, elemento)
	}

	// _ descarta o indice
	for _, elemento := range numeros {
		fmt.Printf("valor %d\n", elemento)
	}
}
