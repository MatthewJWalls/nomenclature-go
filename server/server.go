package server

import (
	"net/http"
	"fmt"
)

// We introduce a Generator interface so that we can support multiple
// types of generators in the future.
type Generator interface {
	Next() string
}

// Given a generator and a channel, fills the channel with
// the results of generator.Next(). This is our basic thread
// safety mechanism.
func pumpGeneratorIntoChannel(gen Generator, nameQueue chan string) {
	for {
		nameQueue <- gen.Next()
	}
}

// Returns a http.Handler function capable of serving the content of a
// Generator.
func NewWebHandler(gen Generator) (func(w http.ResponseWriter, r *http.Request)) {
	
	nameQueue := make(chan string, 5)
	go pumpGeneratorIntoChannel(gen, nameQueue)
	
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{ \"name\" : \"%s\" }",  <- nameQueue )
	}

}

