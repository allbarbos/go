package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/allbarbos/go/alura-web/models"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.GetAllProducts()
	tmpl.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("erro ao converter o preço", err)
		}

		quantidadeConv, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("erro ao converter a quantidade", err)
		}

		models.CreateProduct(nome, descricao, precoConv, quantidadeConv)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	produto := models.GetProduct(id)
	tmpl.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idStr := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		precoStr := r.FormValue("preco")
		quantidadeStr := r.FormValue("quantidade")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println("erro ao converter o id", err)
		}
		preco, err := strconv.ParseFloat(precoStr, 64)
		if err != nil {
			log.Println("erro ao converter o preco", err)
		}
		qtd, err := strconv.Atoi(quantidadeStr)
		if err != nil {
			log.Println("erro ao converter a qtd", err)
		}
		models.UpdateProduct(id, qtd, nome, descricao, preco)
	}
	http.Redirect(w, r, "/", 301)
}
