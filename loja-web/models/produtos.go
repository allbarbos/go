package models

import "github.com/allbarbos/go/loja-web/db"

// Produto representa a entidade
type Produto struct {
	ID, Quantidade  int
	Nome, Descricao string
	Preco           float64
}

// GetAllProducts retorna uma lista de produtos
func GetAllProducts() []Produto {
	db := db.Connection()
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

		p.ID = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	return produtos
}

func CreateProduct(nome, descricao string, preco float64, quantidade int) {
	db := db.Connection()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO public.produtos(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4);")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(nome, descricao, preco, quantidade)
}
