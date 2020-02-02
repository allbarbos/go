package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10 { // não tem while em go
		fmt.Printf("%d ", i)
		i++
	}

	fmt.Println()

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("par ")
		} else {
			fmt.Print("impar ")
		}
	}

	for { //laço infinito
		fmt.Println("pra sempre...")
		time.Sleep(time.Second)
	}
}
