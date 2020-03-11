package main

import (
	"net/http"

	"github.com/allbarbos/go/loja-web/routes"
)

func main() {
	routes.Load()
	http.ListenAndServe(":8000", nil)
}
