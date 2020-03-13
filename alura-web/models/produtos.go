package models

import "github.com/allbarbos/go/alura-web/db"

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

	getAll, err := db.Query("SELECT id, nome, descricao, preco, quantidade FROM public.produtos order by id asc;")
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

func GetProduct(id string) Produto {
	db := db.Connection()
	defer db.Close()

	result, err := db.Query("SELECT id, nome, descricao, preco, quantidade FROM public.produtos where id=$1;", id)
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	for result.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = result.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.ID = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
	}
	return p
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

func DeleteProduct(id string) {
	db := db.Connection()
	defer db.Close()
	query, err := db.Prepare("delete from produtos where id =$1")
	if err != nil {
		panic(err.Error())
	}
	query.Exec(id)
}

func UpdateProduct(id, quantidade int, nome, descricao string, preco float64) {
	db := db.Connection()
	defer db.Close()

	query, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	query.Exec(nome, descricao, preco, quantidade, id)
}
