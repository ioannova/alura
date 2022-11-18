package main

import (
	"fmt"

	database "github.com/ioannova/alura-golang/api-go-rest/dataase"
	"github.com/ioannova/alura-golang/api-go-rest/models"
	"github.com/ioannova/alura-golang/api-go-rest/routes"
)

func main() {
	models.Personaldiades = []models.Personalidade{
		{Id: 1, Nome: "Nelson Mandela", Historia: "Foi um grande homem"},
		{Id: 2, Nome: "Nelson Mandela", Historia: "Foi um grande homem 2"},
	}
	database.ConDb()
	fmt.Println("Iniciando o server http://0.0.0.0:8001")
	routes.HandleRequest()
}
