package web

import (
	"fmt"
	"net/http"
	"testing"
)

const ADDR = "localhost:8080"

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: ADDR,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello, Server!")
	}

	server := http.Server{
		Addr:    ADDR,
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
