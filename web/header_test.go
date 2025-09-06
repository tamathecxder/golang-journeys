package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func PrintHeader(w http.ResponseWriter, r *http.Request) {
	author := r.Header.Get("X-Author")
	fmt.Fprint(w, author)
}

func TestHeader(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, BASE_URL, nil)

	r.Header.Add("X-Author", "gwej")

	PrintHeader(w, r)

	res := w.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
