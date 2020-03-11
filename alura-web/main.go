package main

import (
	"net/http"

	"github.com/allbarbos/go/alura-web/routes"
)

func main() {
	routes.Load()
	http.ListenAndServe(":8000", nil)
}
