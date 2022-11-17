package routes

import (
	"net/http"

	"github.com/alura-golang/app-web/controllers"
)

func Run() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/cadastrar", controllers.Cadastrar)
	http.HandleFunc("/inserir", controllers.Inserir)
	http.HandleFunc("/editar", controllers.Editar)
	http.HandleFunc("/deletar", controllers.Deletar)
}
