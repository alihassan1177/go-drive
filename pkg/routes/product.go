package ProductRoutes;

import (
  "github.com/gorilla/mux"
  "github.com/alihassan1177/ecom-backend/pkg/controllers"
)

func RegisterProductRoutes(router *mux.Router){
  router.HandleFunc("/products/", ProductController.Index).Methods("GET")
  router.HandleFunc("/products/create", ProductController.Create).Methods("POST")
  router.HandleFunc("/products/delete/{id}", ProductController.Delete).Methods("DELETE")
  router.HandleFunc("/products/update/{id}", ProductController.Update).Methods("POST")
  router.HandleFunc("/form", ProductController.CheckForm).Methods("POST")
  router.HandleFunc("/form", ProductController.CheckForm).Methods("GET")
}
