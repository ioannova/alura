package main

import (
	"net/http"

	"github.com/alura-golang/app-web/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.Run()
	http.ListenAndServe(":8001", nil)

}
