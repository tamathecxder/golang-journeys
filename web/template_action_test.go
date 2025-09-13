package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type Hobbies []string

type Experience struct {
	CompanyName string
	Role        string
	IsResigned  bool
}

type WorkExperience struct {
	Detail Experience
}

type Data struct {
	Name, Email    string
	Age            int
	Hobbies        Hobbies
	WorkExperience WorkExperience
}

func PrintOutHtmlAction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./src/action.gohtml"))
	t.ExecuteTemplate(w, "action.gohtml", Data{
		Name:    "Agus",
		Email:   "agus@gmail.com",
		Age:     20,
		Hobbies: Hobbies{"hobby1", "hobby2", "hobby3"},
		WorkExperience: WorkExperience{
			Detail: Experience{
				CompanyName: "PT Kembung Jaya",
				Role:        "Software Engineer",
				IsResigned:  false,
			},
		},
	})
}

func TestTemplateAction(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	rec := httptest.NewRecorder()

	PrintOutHtmlAction(rec, r)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}
