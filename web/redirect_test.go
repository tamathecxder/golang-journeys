package web

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectFrom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from RedirectFrom")
}

func RedirectTo(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/redirect-from", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)

	err := http.ListenAndServe(HOST, mux)

	if err != nil {
		panic(err)
	}
}
