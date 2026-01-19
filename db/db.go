package db

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
	_ "github.com/lib/pq"
)

type ServidorBancodeDados struct {
	tipo string
	url  string
}

// conexao database
func ConectarDatabase() *sql.DB {

	var servidor ServidorBancodeDados

	servidor = ServidorBancodeDados{tipo: "sqlite", url: "my_database.db"}
	//servidor = ServidorBancodeDados{"postgres","user=root password=12345678 dbname=goweb host=localhost sslmode=disable"}

	db, err := sql.Open(servidor.tipo, servidor.url)

	if err != nil {
		panic(err.Error())
	}

	return db
}
