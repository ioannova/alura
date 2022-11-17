package models

import (
	"github.com/alura-golang/app-web/database"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaProdutos() []Produto {
	db := database.GetDb()
	defer db.Close()
	produtosDb, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for produtosDb.Next() {
		var id int
		var quantidade int
		var nome string
		var descricao string
		var preco float64
		err = produtosDb.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p)
	}
	return produtos
}

func CriarProduto(nome string, descricao string, preco float64, quantidade int) {
	db := database.GetDb()
	defer db.Close()
	inserirDados, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	inserirDados.Exec(nome, descricao, preco, quantidade)
}

func DeletarProduto(id string) {
	db := database.GetDb()
	defer db.Close()
	deletarProduto, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)
}

func VisualizarProduto(id string) Produto {
	db := database.GetDb()
	defer db.Close()

	produtoDb, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}
	for produtoDb.Next() {
		var id int
		var quantidade int
		var nome string
		var descricao string
		var preco float64

		err = produtoDb.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}

	return produto
}

func AtualizarProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := database.GetDb()
	defer db.Close()

	produtoDb, err := db.Prepare("UPDATE produtos SET nome=$2, descricao=$3, preco=$4, quantidade=$5 WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}
	produtoDb.Exec(id, nome, descricao, preco, quantidade)
}
