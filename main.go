package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=39151626dim host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Na moda", Preco: 39.99, Quantidade: 20},
		{"Tenis", "Confortavel", 89, 3},
		{"Oculos", "Kurt Cobain", 9, 666},
		{"Presente misterioso", "2 reais ou presente misterioso", 2, 1},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
