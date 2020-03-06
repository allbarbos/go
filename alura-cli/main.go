package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
			printLog()
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
		logging(site, true)
		fmt.Println("status up")
	} else {
		logging(site, false)
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

func logging(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("2006-01-02 15:04:05") + " " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLog() {
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("[ERROR]", err)
	}
	fmt.Println(string(file))
}
