package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterTest(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello Word")
	})

	request := httptest.NewRequest("GET", "http://localhost:4000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Hello Word", string(bytes))

}

//TODO: NAMED PARAMETER

func TestRouteNamedTest(t *testing.T) {
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

//TODO: CATCH PARAMETER

func TestCatchAllParameterTest(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := "Image : " + params.ByName("image")
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:4000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Image : /small/profile.png", string(bytes))
}

func TestRouterParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := "Product " + params.ByName("id")
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:4000/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1", string(bytes))
}

func TestRouterLogsParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/logs/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := "Logs " + params.ByName("id")
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:4000/logs/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Logs 1", string(bytes))
}
