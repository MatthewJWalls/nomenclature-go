package server

import (
	"testing"
	"net/http"	
	"io/ioutil"	
	"net/http/httptest"
)

// mock generator

type MockGenerator struct { }

func (this MockGenerator) Next() string {
	return "yup"
}

// test server

func TestNewWebHandler(t *testing.T) {

	gen := MockGenerator{}
	
	ts := httptest.NewServer(
		http.HandlerFunc(NewWebHandler(gen)),
	)

	defer ts.Close()

	result, err := http.Get(ts.URL)

	if err != nil {
		t.Errorf("No response from server")
	}

	body, _ := ioutil.ReadAll(result.Body)

	if string(body) != "{ \"name\" : \"yup\" }" {
		t.Errorf("Failed to serve correct generator value")
	}
	
	result.Body.Close()
	
}

// test pump

func TestPumping(t *testing.T) {

	testQueue := make(chan string, 2)
	
	go pumpGeneratorIntoChannel(MockGenerator{}, testQueue)

	if <- testQueue != "yup" {
		t.Errorf("Queue was not filled with data")
	}
	
}
