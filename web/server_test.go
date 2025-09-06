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

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Mikum")
	})

	mux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "gwej")
	})

	server := http.Server{
		Addr:    ADDR,
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Method)
		fmt.Fprint(w, r.RequestURI)
	})

	server := http.Server{
		Addr:    ADDR,
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
