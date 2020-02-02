package main

import (
	"fmt"
	"time"
)

func nota(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		fallthrough
		// caso entre no case, fallthrough obriga a executar os cases abaixo, sem ele é retornado direto
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float64, float32:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "tipo não conhecido"
	}
}

func main() {
	// ## switch 01
	fmt.Println("------------- switch 01")
	fmt.Println(nota(9.8))
	fmt.Println(nota(6.9))
	fmt.Println(nota(2.1))
	fmt.Println(nota(11))

	// ## switch 02
	fmt.Println("\n------------- switch 02")
	t := time.Now()
	switch { // switch true - quando é declarado sem condição, o primeiro case TRUE que acontecer ele entra
	case t.Hour() < 12:
		fmt.Println("bom dia")
	case t.Hour() < 18:
		fmt.Println("boa tarde")
	default:
		fmt.Println("boa noite")
	}

	// ## switch 03
	fmt.Println("\n------------- switch 03")
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(13))
	fmt.Println(tipo("bile"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
