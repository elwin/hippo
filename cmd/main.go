package main

import (
	"github.com/elwin/framework/pkg/hippo"
	"log"
	"net/http"
)

func main() {

	handler := &hippo.ServerHandler{
		Handler: sampleHandler(),
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func sampleHandler() hippo.HandlerFunc {
	return func(request hippo.Request) hippo.Response {
		response := &hippo.BaseResponse{}
		response.AddCookie(http.Cookie{Name: "a", Value: "xxx"})
		response.AddCookie(http.Cookie{Name: "b", Value: "yyy"})
		val := request.Form().Get("foo")
		response.SetBody(val)

		return response
	}
}
