package main

import "fmt"

func main() {
	// array tem uma estrutura homogênea (mesmo tipo) e estática (fixo)
	var notas [3]float64
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Printf("média: %.2f", media)
}
