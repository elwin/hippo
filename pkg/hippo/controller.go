package hippo

import (
	"fmt"
	"net/http"
)

type ServerHandler struct {
	mapping map[string]HandlerFunc
}

func New() ServerHandler {
	return ServerHandler{mapping: map[string]HandlerFunc{}}
}

func (s ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
	}

	request := Request{
		url:      *r.URL,
		header:   r.Header,
		form:     r.Form,
		postForm: r.PostForm,
		cookies:  r.Cookies(),
	}

	response := s.mapping[request.url.Path](request)

	header := w.Header()
	for name, values := range response.header {
		for _, value := range values {
			header.Add(name, value)
		}
	}

	w.WriteHeader(response.statusCode)

	if _, err := w.Write([]byte(response.Body())); err != nil {
		fmt.Println(err)
	}
}

func (s ServerHandler) Get(path string, handler HandlerFunc, middleware ...Middleware) {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	s.mapping[path] = handler
}
