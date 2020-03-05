package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitor = 2
const delay = 5

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
	sites := readSitesFile()

	for i := 0; i < monitor; i++ {
		for _, site := range sites {
			ping(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func ping(site string) {
	fmt.Println("ping", site)
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("[ERROR]", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("status up")
	} else {
		fmt.Println("status DOWN")
	}
}

func readSitesFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")
	//file, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("[ERROR]", err)
	}

	reader := bufio.NewReader(file)

	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)
		sites = append(sites, row)
		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}
