package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connection() *sql.DB {
	cnx := "user=postgres dbname=loja password=bile host=db sslmode=disable"
	db, err := sql.Open("postgres", cnx)
	if err != nil {
		panic(err.Error())
	}
	return db
}
