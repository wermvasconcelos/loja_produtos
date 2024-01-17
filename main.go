package main

import (
	"loja_produtos/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":3000", nil)
}
