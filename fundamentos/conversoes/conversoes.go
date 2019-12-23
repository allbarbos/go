package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))
	fmt.Println("--------------")

	nota := 6.9
	final := int(nota)
	fmt.Println(final)
	fmt.Println("--------------")

	// cuidado - esta função busca a referencia na tabela unicode ao invés de converter
	fmt.Println("teste " + string(97))
	fmt.Println("--------------")

	// int para str
	fmt.Println("teste " + strconv.Itoa(123))

	// str para int
	num, _ := strconv.Atoi("321")
	fmt.Println(num)

	// str para bool
	bol, _ := strconv.ParseBool("true")
	fmt.Println(bol)
}
