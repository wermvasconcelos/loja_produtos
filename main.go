package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Na moda", Preco: 39.99, Quantidade: 20},
		{"Tenis", "Confortavel", 89, 3},
		{"Oculos", "Kurt Cobain", 9, 666},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
