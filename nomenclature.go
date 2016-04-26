package main

import (
	"log"
	"flag"
	"strings"
	"net/http"
	"io/ioutil"
	
	"github.com/MatthewJWalls/nomenclature/server"
	"github.com/MatthewJWalls/nomenclature/generator"
)

// This is the naive approach for loading the word lists.
// In the future we may want to deal with potentially large
// lists which will need a bit of a rethink here so we're
// only reading in the bytes we want.

func readWordList(path string) ([]string) {

	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("Failed to read word list %q", path)
	}

	raw := strings.Split(string(data), "\n")
	out := []string{}

	for i := 0; i < len(raw); i++ {
		if strings.TrimSpace(raw[i]) != "" {
			out = append(out, strings.TrimSpace(raw[i]))
		}
	}

	return out

}

func main() {

	log.Printf("Initialising data")

	prefile := flag.String("f1", "pre", "path to prefix words file")
	postfile := flag.String("f2", "post", "path to prefix words file")

	flag.Parse()
	
	pre := readWordList(*prefile)
	pst := readWordList(*postfile)
	
	log.Printf("Starting server")
	
	a := generator.NewStandardGenerator(pre, pst)
	
	http.HandleFunc("/", server.NewWebHandler(a))
	log.Fatal(http.ListenAndServe(":8080", nil))

}

