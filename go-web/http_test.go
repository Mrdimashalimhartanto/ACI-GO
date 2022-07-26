package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloCMSHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hallo Broker")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9000/hellomember", nil)

	recorder := httptest.NewRecorder()

	HelloCMSHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))

	//bodyString := string(body)

	//fmt.Println(bodyString)
}
