package web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	dir := http.Dir("./src")
	fileServer := http.FileServer(dir)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

//go:embed src
var source embed.FS

func TestFileServerEmbed(t *testing.T) {
	dir, err := fs.Sub(source, "src")

	if err != nil {
		panic(err)
	}

	fileServer := http.FileServer(http.FS(dir))
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	serveError := server.ListenAndServe()

	if serveError != nil {
		panic(serveError)
	}
}
