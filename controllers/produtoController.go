package controllers

import (
	"fmt"
	"go_web_app/models"
	"net/http"
	"text/template"
)

// var my_templates = template.Must(template.ParseGlob("templates/*.html"))
var my_templates, _ = template.ParseFiles("templates/index.html", "templates/new.html")

// pagina index
func IndexProduto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TEMPLATE INDEX")
	lista := models.ListarProdutos()
	my_templates.ExecuteTemplate(w, "Index", lista)
	fmt.Println("--------------------------------")
}

func NewProduto(w http.ResponseWriter, r *http.Request) {

	fmt.Println("TEMPLATE NEW PRODUTO ")
	my_templates.ExecuteTemplate(w, "New", nil)
	fmt.Println("--------------------------------")
}
