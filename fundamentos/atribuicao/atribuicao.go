package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // atribuição por inferencia de tipo
	i += 3
	i -= 3
	i /= 3
	i *= 3
	i %= 3
	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	// troca valor das variaveis
	x, y = y, x
	fmt.Println(x, y)
}
