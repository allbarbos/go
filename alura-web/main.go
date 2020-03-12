package main

import (
	"fmt"
	"net/http"

	"github.com/allbarbos/go/alura-web/routes"
)

func main() {
	routes.Load()
	fmt.Println("listen port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err.Error())
	}
}
