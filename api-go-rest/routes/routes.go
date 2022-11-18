package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ioannova/alura-golang/api-go-rest/controllers"
	"github.com/ioannova/alura-golang/api-go-rest/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditarPersonalidade).Methods("Put")
	r.HandleFunc("/api/personalidades", controllers.CriarPersonalidade).Methods("Post")
	log.Fatal(http.ListenAndServe(":8001", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
