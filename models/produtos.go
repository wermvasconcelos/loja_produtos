package models

import (
	"loja_produtos/db"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade, Id  int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &quantidade, &preco)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Quantidade = quantidade
		p.Preco = preco

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(idDoProduto string) {
	db := db.ConectaComBancoDeDados()

	deletarProduto, err := db.Prepare("DELETE FROM produtos WHERE ID = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(idDoProduto)

	defer db.Close()
}

func EditaProduto(id string) Produto {
	db.ConectaComBancoDeDados()

	produtoDoBanco, err := db.ConectaComBancoDeDados().Query("SELECT * FROM produtos WHERE ID = $1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &quantidade, &preco)
		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Quantidade = quantidade
		produtoParaAtualizar.Preco = preco
	}

	defer db.ConectaComBancoDeDados().Close()
	return produtoParaAtualizar
}

func UpdateProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	updateProduto, err := db.Prepare("UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5")

	if err != nil {
		panic(err.Error())
	}

	updateProduto.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
