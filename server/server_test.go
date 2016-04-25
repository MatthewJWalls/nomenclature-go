package server

import (
	"testing"
	_ "net/http"
	_ "net/http/httptest"
)

// mock generator

type MockGenerator struct { }

func (this MockGenerator) Next() string {
	return "yup"
}

func TestNewWebHandler(t *testing.T) {
	// todo
}
