package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jdjaxon/learn_go/bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
} // main
