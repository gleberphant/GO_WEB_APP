package routes

import (
	"go_web_app/controllers"
	"net/http"
)

// função apenas para carregar as rotas
func CarregarRotas() {
	http.HandleFunc("/", controllers.IndexProduto)

	http.HandleFunc("/new", controllers.NewProduto)
}
