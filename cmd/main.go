package main

import (
	"log"
	"net/http"
)
import "github.com/elwin/hippo/pkg/hippo"

func main() {

	mux := hippo.New()
	mux.Get("/", index).WithMiddleware(hippo.TimeMiddleware, hippo.LogMiddleware, hippo.SessionMiddleware)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

func index(request hippo.Request) hippo.Response {
	return hippo.NewResponse().WithBody("Hello World")
}
