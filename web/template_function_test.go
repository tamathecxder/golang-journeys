package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

type PageFunc struct {
	Header string
}

func (pf PageFunc) SetHeader(header string) string {
	return fmt.Sprintf("Curernt Page header is: %s. Replaced to: %s", pf.Header, header)
}

func PrintOutHtmlTemplateFunction(w http.ResponseWriter, r *http.Request) {
	// tmpl := `{{len "Agus"}}` //
	tmpl := `{{.SetHeader "TempFunc_Module"}}`
	key := "FUNCTION"

	t := template.Must(template.New(key).Parse(tmpl))
	t.ExecuteTemplate(w, key, PageFunc{
		Header: "DEFAULT_HEADER",
	})
}

func TestTemplateFunction(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	rec := httptest.NewRecorder()

	PrintOutHtmlTemplateFunction(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

func PrintOutHtmlTemplateGlobalFunction(w http.ResponseWriter, r *http.Request) {
	key := "GLOBALFUNCTION"
	t := template.New(key)

	t.Funcs(template.FuncMap{
		"upper": func(name string) string {
			return strings.ToUpper(name)
		},
	})

	t = template.Must(t.Parse(`{{ upper .Header}}`))

	t.ExecuteTemplate(w, key, PageFunc{
		Header: "def jam",
	})
}

func TestTemplateGlobalFunction(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	rec := httptest.NewRecorder()

	PrintOutHtmlTemplateGlobalFunction(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

func PrintOutHtmlTemplateGlobalFunctionPipeline(w http.ResponseWriter, r *http.Request) {
	key := "GLOBALFUNCTIONPipeline"
	t := template.New(key)

	t.Funcs(template.FuncMap{
		"SaySalam": func(name string) string {
			return fmt.Sprintf("Salam Alaykum: %s", name)
		},
		"upper": func(name string) string {
			return strings.ToUpper(name)
		},
	})

	t = template.Must(t.Parse(`{{ SaySalam .Header | upper }}`))

	t.ExecuteTemplate(w, key, PageFunc{
		Header: "def jam",
	})
}

func TestTemplateGlobalFunctionPipeline(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	rec := httptest.NewRecorder()

	PrintOutHtmlTemplateGlobalFunctionPipeline(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}
