package routes

import (
	"github/fokaz-c/go-book-manager/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookID}", controllers.DeleteBook).Methods("DELETE")

	router.HandleFunc("/tags/", controllers.CreateTag).Methods("POST")
	router.HandleFunc("/tags/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/tags/{tags}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/tags/{tags}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/tags/{tags}", controllers.DeleteBook).Methods("DELETE")

}
