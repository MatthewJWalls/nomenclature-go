package main

import (
	"log"
	"net/http"
	
	"github.com/MatthewJWalls/nomenclature/server"
	"github.com/MatthewJWalls/nomenclature/generator"
)


func main() {

	log.Printf("Starting server")
	
	pre := make([]string, 0)
	pst := make([]string, 0)

	a := generator.NewStandardGenerator(&pre, &pst)
	
	http.HandleFunc("/", server.NewWebHandler(a))
	log.Fatal(http.ListenAndServe(":8080", nil))

}

