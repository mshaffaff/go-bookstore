package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mshaffaff/go-bookstore/pkg/routes"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server at port : 9010\n")
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
