package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func RenderHTML(w http.ResponseWriter, r *http.Request) {
	tmpl := `<html><body>{{.}}</body></html>`
	key := "Placeholder"

	// t, err := template.New(key).Parse(tmpl)
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.New(key).Parse(tmpl))
	t.ExecuteTemplate(w, key, "Whatever!")
}

func TestTemplate(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	rec := httptest.NewRecorder()

	RenderHTML(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}
