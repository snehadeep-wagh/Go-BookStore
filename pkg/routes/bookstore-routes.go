package routes

import (
	"github.com/gorilla/mux"
	"github.com/snehadeep-wagh/go-bookstore/pkg/controllers"
)

var GetAllRoutes = func(r *mux.Router) {
	r.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
