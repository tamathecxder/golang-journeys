package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func DummyHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "(DUMMY_RESPONSE)")
}

func TestHttp(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	recorder := httptest.NewRecorder()

	DummyHandler(recorder, req)

	res := recorder.Result()
	body := res.Body

	result, _ := io.ReadAll(body)

	fmt.Println(string(result))
}
