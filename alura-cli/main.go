package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const MONITORING = 2
const DELAY = 5

func main() {
	for {
		printMenu()

		switch option() {
		case 1:
			monitoring()
		case 2:
			fmt.Println("Exibindo...")
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Opção inválida!")
			os.Exit(-1)
		}
	}
}

func printMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair")
}

func option() int {
	var cmd int
	fmt.Scan(&cmd)
	return cmd
}

func monitoring() {
	sites := []string{
		"https://random-status-code.herokuapp.com",
		"https://g1.globo.com",
		"https://resultadosdigitais.com.br",
	}

	for i := 0; i < MONITORING; i++ {
		for _, site := range sites {
			ping(site)
		}
		time.Sleep(DELAY * time.Second)
		fmt.Println("")
	}
}

func ping(site string) {
	fmt.Println("ping", site)
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("status up")
	} else {
		fmt.Println("status DOWN")
	}
}
