package routes

import (
	"github.com/gorilla/mux"
	"crud-api-with-mongodb/controllers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

	return router
}
