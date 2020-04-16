package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := 0
	for {
		n++
		i := rand.Intn(4200)
		fmt.Println(i)
		if i%42 == 0 {
			break
		}
	}
	fmt.Printf("Saída após %d iterações.\n", n)

	// #### Break em bloco externo
	var i int
externo:
	for {
		for i = 0; i < 10; i++ {
			if i == 8 {
				break externo
			}
		}
	}
	fmt.Println(i) // 8
}
