package main

import (
	"log"
	"net/http"

	"github.com/alihassan1177/ecom-backend/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	ProductRoutes.RegisterProductRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
