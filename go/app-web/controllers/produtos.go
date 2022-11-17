package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/alura/go/app-web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func Cadastrar(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Cadastrar", nil)
}

func Inserir(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter preco para float: ", err)
		}

		quantidadeConv, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro ao converter quantidade para int: ", err)
		}
		models.CriarProduto(nome, descricao, precoConv, quantidadeConv)
	}
	http.Redirect(w, r, "/", 301)
}

func Editar(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	if r.Method == "GET" {
		produto := models.VisualizarProduto(idProduto)
		temp.ExecuteTemplate(w, "Editar", produto)
	} else if r.Method == "POST" {
		_id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		_preco := r.FormValue("preco")
		_quantidade := r.FormValue("quantidade")

		id, err := strconv.Atoi(_id)
		if err != nil {
			log.Println("Erro ao converter id para inteiro: ", err)
		}
		preco, err := strconv.ParseFloat(_preco, 64)
		if err != nil {
			log.Println("Erro ao converter preco para float64: ", err)
		}
		quantidade, err := strconv.Atoi(_quantidade)
		if err != nil {
			log.Println("Erro ao converter quantidade para inteiro: ", err)
		}
		models.AtualizarProduto(id, nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", 301)
}

func Deletar(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	// id, err := strconv.Atoi(idProduto)
	// if err != nil {
	// 	panic(err.Error())
	// }

	models.DeletarProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}
