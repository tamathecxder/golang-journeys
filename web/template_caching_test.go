package web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

//go:embed src/*.gohtml
var src embed.FS

var templates = template.Must(template.ParseFS(src, "src/*.gohtml"))

func PrintOutHtmlTemplateCaching(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.gohtml", map[string]any{
		"Title":   "Test Title",
		"Heading": "Test Heading",
	})
}

func TestTemplateCaching(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	rec := httptest.NewRecorder()

	PrintOutHtmlTemplateCaching(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}
