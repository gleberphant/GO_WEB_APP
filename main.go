package main

import (
	"fmt"
	"go_web_app/routes"
	"net/http"
)

// subir o servidor
func main() {
	my_server := ":8000"

	fmt.Println("Ouvindo servidor HTTP: ", my_server)
	routes.CarregarRotas()
	http.ListenAndServe(my_server, nil)
}
