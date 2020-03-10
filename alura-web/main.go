package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaDb() *sql.DB {
	cnx := "user=postgres dbname=alura_loja password=bile host=localhost sslmode=disable"
	db, err := sql.Open("postgres", cnx)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Id, Quantidade  int
	Nome, Descricao string
	Preco           float64
}

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaDb()
	defer db.Close()

	getAll, err := db.Query("SELECT id, nome, descricao, preco, quantidade FROM public.produtos;")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for getAll.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = getAll.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	tmpl.ExecuteTemplate(w, "Index", produtos)
}
