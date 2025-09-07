package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func HandleSetCookie(w http.ResponseWriter, r *http.Request) {
	var v string = "gwej"

	if r.URL.Query().Has("Name") {
		v = r.URL.Query().Get("name")
	}

	cookie := new(http.Cookie)
	cookie.Name = "X-Powered-By"
	cookie.Value = v
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Cookie stored successfully")
}

func HandleGetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-Powered-By")

	if err != nil {
		fmt.Fprint(w, "Cookie not found")
	} else {
		v := cookie.Value
		fmt.Fprintf(w, "Cookie Value: %s", v)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HandleGetCookie)
	mux.HandleFunc("/set-cookie", HandleSetCookie)

	host := strings.Replace(BASE_URL, "http://", "", -1)
	err := http.ListenAndServe(strings.TrimSpace(host), mux)

	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	wrc := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, BASE_URL+"?name=whatever", nil)

	HandleSetCookie(wrc, r)

	cookies := wrc.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s: %s", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	wrc := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, BASE_URL+"?name=whatever", nil)

	cookie := new(http.Cookie)
	cookie.Name = "X-Powered-By"
	cookie.Value = "VALUE"
	r.AddCookie(cookie)

	HandleGetCookie(wrc, r)

	res := wrc.Result().Body
	body, _ := io.ReadAll(res)

	fmt.Println("Cookie Body:", string(body))
}
