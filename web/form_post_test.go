package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func PrintPayload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		panic(err)
	}

	name := r.PostForm.Get("name")
	email := r.PostForm.Get("email")

	fmt.Fprintf(w, "%s", fmt.Sprintf("Current user name is %s and the email is %s", name, email))
}

func TestFormPost(t *testing.T) {
	payload := strings.NewReader("name=Ganjar&email=pranowo@gov.id")

	wrc := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, BASE_URL, payload)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	PrintPayload(wrc, r)

	res := wrc.Result()

	body, _ := io.ReadAll(res.Body)

	fmt.Println("Body:", string(body))
}
