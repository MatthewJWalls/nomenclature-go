package main

import (
	"log"
	"net/http"
	
	"github.com/MatthewJWalls/nomenclature/server"
	"github.com/MatthewJWalls/nomenclature/generator"
)


func main() {

	log.Printf("Starting server")
	
	pre := []string{"Big", "Bad", "Horrible"}
	pst := []string{"Wolf", "Chicken", "Hernia"}

	a := generator.NewStandardGenerator(pre, pst)
	
	http.HandleFunc("/", server.NewWebHandler(a))
	log.Fatal(http.ListenAndServe(":8080", nil))

}

