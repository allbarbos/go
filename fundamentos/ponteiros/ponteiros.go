package main

import "fmt"

func main() {
	i := 1

	// criando um ponteiro sem referência de endereço de valor
	var p *int = nil

	// pega o endereço da variável e atribui ao ponteiro
	p = &i
	*p++
	i++

	fmt.Println(p, *p, i, &i)

	/*
		ao utilizar * antes de uma variável de ponteiro é possível acessar o valor dela - *p
		sem * acessamos a raferência de memória

		não é permitido realizar operações aritiméticas em cima de ponteiros, ex.: p++
	*/

	fmt.Println(p, *p)
}
