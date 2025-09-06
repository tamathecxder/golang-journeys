package web

import (
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
