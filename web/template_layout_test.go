package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func PrintOutHtmlTemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./src/header.gohtml", "./src/layout.gohtml", "./src/footer.gohtml"))
	t.ExecuteTemplate(w, "layout", map[string]any{
		"Title": "Testing",
		"Name":  "Agus",
	})
}

func TestTemplateLayout(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	rec := httptest.NewRecorder()

	PrintOutHtmlTemplateLayout(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}
