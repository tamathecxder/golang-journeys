package web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("keyword") != "" {
		http.ServeFile(w, r, "./src/index.html")
	} else {
		http.ServeFile(w, r, "./src/404.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

//go:embed src/index.html
var indexTemplate string

//go:embed src/404.html
var notFoundTemplate string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("keyword") != "" {
		fmt.Fprintf(w, "%s", indexTemplate)
	} else {
		fmt.Fprintf(w, "%s", notFoundTemplate)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
