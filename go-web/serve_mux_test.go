package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello Broker")
	})

	mux.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "welcome to login cms epay")
	})

	mux.HandleFunc("/cmsepay/register/formregister", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Please Register to get information about cms")
	})

	mux.HandleFunc("/member/cms/homepage", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "welcome to dashboard CMS Epay")
	})

	mux.HandleFunc("/member/transactiondetail", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Detail Transaction Member")
	})

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
