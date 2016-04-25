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

// Returns a http.Handler function capable of serving the content of a
// Generator.
func NewWebHandler(gen Generator) (func(w http.ResponseWriter, r *http.Request)) {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, gen.Next())
	}

}

