package controllers

import (
	"html/template"
	"net/http"

	"github.com/allbarbos/go/alura-web/models"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.GetAllProducts()
	tmpl.ExecuteTemplate(w, "Index", produtos)
}
