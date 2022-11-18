package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Curso struct {
	gorm.Model
	Nome      string
	Area      string
	Professor string
}

func main() {
	db, err := gorm.Open(sqlite.Open("curso.db"), &gorm.Config{})
	if err != nil {
		panic("Erro ao acessar banco de dados: ")
	}
	db.AutoMigrate(&Curso{})

	db.Create(&Curso{Nome: "Go: Fundamentos de uma aplicação web", Area: "Programação", Professor: "Guilherme Lima"})

	var curso Curso
	db.First(&curso, 1)
	fmt.Println(&curso)

	db.Model(&curso).Update("Nome", "Go: Fundamentos de um amigação Web com GORM")
	db.First(&curso, 1)
	fmt.Println(&curso)
	db.Delete(&curso, 1)
}
