package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const BaseURL string = "http://localhost:8080"

func HandleSearch(w http.ResponseWriter, r *http.Request) string {
	keyword := r.URL.Query().Get("keyword")

	if keyword != "" {
		return fmt.Sprintf("Keyword found: %s", keyword)
	}

	return ""
}

func HandleMultipleSearch(w http.ResponseWriter, r *http.Request) string {
	queryParam := r.URL.Query()

	var keywords []string = queryParam["keyword"]

	if len(keywords) > 0 {
		return fmt.Sprintf("Keyword found: %s", strings.Join(keywords, ", "))
	}

	return ""
}

func TestQueryParam(t *testing.T) {
	param := "?keyword=something"
	recWriter := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf(BaseURL+"/%s", param), nil)
	search := HandleSearch(recWriter, req)

	fmt.Fprint(recWriter, search)

	response := recWriter.Body
	body, _ := io.ReadAll(response)

	fmt.Println(string(body))
}

func TestMultipleQueryParamValues(t *testing.T) {
	param := "?keyword=something&keyword=blablabla"
	recWriter := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, BaseURL+param, nil)

	search := HandleMultipleSearch(recWriter, req)
	fmt.Fprint(recWriter, search)

	response := recWriter.Body
	body, _ := io.ReadAll(response)

	fmt.Println(string(body))
}
