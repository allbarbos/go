package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connection() *sql.DB {
	cnx := "user=postgres dbname=alura_loja password=bile host=localhost sslmode=disable"
	db, err := sql.Open("postgres", cnx)
	if err != nil {
		panic(err.Error())
	}
	return db
}
