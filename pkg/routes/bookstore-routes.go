package routes

import (
	"github/fokaz-c/go-book-manager/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookID}", controllers.DeleteBook).Methods("DELETE")
}
