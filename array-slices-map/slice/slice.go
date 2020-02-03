package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}
	slice := []int{1, 2, 3} // o slice possui tamanho dinâmico - slice NÃO É UM ARRAY, ele apenas define um pedaço de um array.

	fmt.Println(array, slice)

	arr2 := [5]int{1, 2, 3, 4, 5}
	slc2 := arr2[1:3]
	fmt.Println(arr2, slc2)

	slc3 := arr2[:2]
	fmt.Println(arr2, slc3)

	slc4 := slc2[:1]
	fmt.Println(slc2, slc4)
}
