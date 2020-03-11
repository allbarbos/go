package routes

import (
	"net/http"

	"github.com/allbarbos/go/alura-web/controllers"
)

// Load carrega todas as rotas da aplicação
func Load() {
	http.HandleFunc("/", controllers.Index)
}
