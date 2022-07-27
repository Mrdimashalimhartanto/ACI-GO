package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./templates/simple.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML CMS")
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:9090/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	resposne := recorder.Result()
	body, _ := io.ReadAll(resposne.Body)
	fmt.Println(string(body))
}
