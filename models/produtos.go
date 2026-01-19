package models

import (
	"fmt"
	"go_web_app/db"
)

// modelo de produto
type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float32
	Quantidade int
}

// -----
func ListarProdutos() []Produto {

	fmt.Println("listar produtos")
	// conectar ao banco de dados
	my_db := db.ConectarDatabase()
	defer my_db.Close()

	// realizar consulta
	query, err := my_db.Query("SELECT * FROM produtos")

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
	return produtos
}
