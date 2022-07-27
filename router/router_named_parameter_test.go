package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouteTest(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(writer http.ResponseWriter, request *http.Request,
		params httprouter.Params) {
		text := "Product " + params.ByName("id") + " Item " + params.ByName("itemId")
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:4000/products/1/items/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1 Item 2", string(bytes))

}
