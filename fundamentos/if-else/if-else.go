package main

import "fmt"

func imprime(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com", nota)
	} else {
		fmt.Println("Reprovado com", nota)
	}
}

func main() {
	imprime(7.3)
	imprime(5.3)
}
