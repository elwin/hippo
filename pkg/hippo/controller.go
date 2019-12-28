package hippo

import (
	"fmt"
	"net/http"
)

type ServerHandler struct {
	Handler HandlerFunc
}

func (s ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
	}

	request := BaseRequest{
		url:      *r.URL,
		header:   r.Header,
		form:     r.Form,
		postForm: r.PostForm,
	}


	response := s.Handler(request)

	header := w.Header()
	for name, values := range response.Header() {
		for _, value := range values {
			header.Add(name, value)
		}
	}

	w.WriteHeader(response.StatusCode())

	if _, err := w.Write([]byte(response.Body())); err != nil {
		fmt.Println(err)
	}
}
