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

func PrintOutHtml(w http.ResponseWriter, r *http.Request) {
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

	PrintOutHtml(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

func PrintOutHtmlFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./src/index.gohtml"))
	t.ExecuteTemplate(w, "index.gohtml", "123")
}

func TestTemplateFile(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	rec := httptest.NewRecorder()

	PrintOutHtmlFile(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

func PrintOutHtmlGlob(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./src/*.gohtml"))
	t.ExecuteTemplate(w, "index.gohtml", "123")
}

func TestTemplateGlob(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	rec := httptest.NewRecorder()

	PrintOutHtmlGlob(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

//go:embed src/*.gohtml
var goHTMLDir embed.FS

func PrintOutHtmlEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(goHTMLDir, "src/*.gohtml"))
	t.ExecuteTemplate(w, "index.gohtml", "123")
}

func TestTemplateEmbed(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	rec := httptest.NewRecorder()

	PrintOutHtmlEmbed(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}
