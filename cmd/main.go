package main

import (
	"fmt"
	"log"
	"net/http"
)
import "github.com/elwin/hippo/pkg/hippo"

func main() {

	mux := hippo.New()
	mux.Get("/", index).WithMiddleware(hippo.TimeMiddleware, hippo.LogMiddleware, mux.SessionMiddleware)
	mux.Get("/redirect", redirect).WithMiddleware(hippo.TimeMiddleware, hippo.LogMiddleware, mux.SessionMiddleware)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

func index(request hippo.Request) hippo.Response {
	msg, _ := request.Session.Get("message")

	return hippo.NewResponse().WithBody(fmt.Sprintf("Hello World: %s", msg))
}

func redirect(request hippo.Request) hippo.Response {
	msg, ok := request.Query("message")
	if ok {
		request.Session.Set("message", msg)
	}

	return hippo.NewRedirect("/")
}
