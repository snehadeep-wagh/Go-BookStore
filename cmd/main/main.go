package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/snehadeep-wagh/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.GetAllRoutes(r)
	http.Handle("/", r)

	fmt.Println("Strting the server at port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))

}
