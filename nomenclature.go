package main

import (
	"os"
	"log"
	"flag"
	"net/http"

	"github.com/MatthewJWalls/nomenclature/util"	
	"github.com/MatthewJWalls/nomenclature/server"
	"github.com/MatthewJWalls/nomenclature/generator"
)

func main() {

	port := flag.String("p", "80", "port to serve names on")
	state := flag.String("s", ".state", "path to state file")
	prefile := flag.String("f1", "pre", "path to prefix words file")
	postfile := flag.String("f2", "post", "path to prefix words file")

	flag.Parse()

	log.Printf("Initialising data")
	
	pre := util.GetFileLines(*prefile)
	pst := util.GetFileLines(*postfile)
	
	log.Printf("Starting server on port %s", *port)
	
	a := generator.NewStandardGenerator(pre, pst, *state)
	
	if _, err := os.Stat(*state); err == nil {
		
		log.Printf("Using state file %q", *state)
		_ = a.Load(*state)
		
	}
	
	http.HandleFunc("/", server.NewWebHandler(a))
	log.Fatal(http.ListenAndServe(":"+*port, nil))

}

