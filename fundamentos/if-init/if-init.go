package main

import (
	"fmt"
	"math/rand"
	"time"
)

func aleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := aleatorio(); i > 5 { // tb suportado por swtich
		fmt.Println("Ganhou!")
	} else {
		fmt.Println("Perdeu")
	}
}
