package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/glebarez/go-sqlite"
	_ "github.com/lib/pq"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float32
	Quantidade int
}

var my_server string = ":8000"

// variavel que armazena os templates html
var my_templates = template.Must(template.ParseGlob("templates/*.html"))

// conexao database
func conectarDatabase() *sql.DB {
	// conexao:= map[string]string{
	// 	"tipo": "sqlite",
	// 	"url": "my_database.db",
	// }
	conexao := map[string]string{
		"tipo": "postgres",
		"url":  "user=root password=12345678 dbname=goweb host=localhost sslmode=disable",
	}

	db, err := sql.Open(conexao["tipo"], conexao["url"])

	if err != nil {
		panic(err.Error())
	}

	return db
}

// main
func main() {
	http.HandleFunc("/", indexPage)

	fmt.Println("Ouvindo servidor HTTP: ", my_server)
	http.ListenAndServe(my_server, nil)

}

// pagina inder
func indexPage(w http.ResponseWriter, r *http.Request) {
	// conectar ao banco de dados
	fmt.Println("Conectando ao banco de dados")
	db := conectarDatabase()
	defer db.Close()

	// realizar consulta
	query, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	produtos := []Produto{}

	// ler cada linha do banco de dados
	for query.Next() {
		row := Produto{}
		if err = query.Scan(&row.Id, &row.Nome, &row.Descricao, &row.Quantidade, &row.Preco); err != nil {
			panic(err.Error())
		}

		fmt.Println("resultado: ", row)
		//adicionar a uma lista
		produtos = append(produtos, row)
	}

	// executar o template
	my_templates.ExecuteTemplate(w, "index", produtos)
}
