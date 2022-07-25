package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello HTTPRouter")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:4000",
	}

	server.ListenAndServe()
}
