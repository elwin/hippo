package hippo

import (
	"fmt"
	"net/http"
)

const (
	verbGET = "get"
)

type ServerHandler struct {
	mapping map[pattern]*handler
}

type handler struct {
	f  HandlerFunc
	ms []Middleware
}

// TODO: Chain middlewares only once
func (h handler) handle(request Request) Response {
	f := h.f
	for i := len(h.ms) - 1; i >= 0; i-- {
		f = h.ms[i](f)
	}

	return f(request)
}

func New() *ServerHandler {
	return &ServerHandler{mapping: map[pattern]*handler{}}
}

func (s *ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	p := pattern{
		verb: verbGET,
		path: r.URL.Path,
		s:    s,
	}
	response := s.mapping[p].handle(request)

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

func (s *ServerHandler) Get(path string, h HandlerFunc) pattern {
	p := pattern{
		verb: verbGET,
		path: path,
		s:    s,
	}

	s.mapping[p] = &handler{f: h}

	return p
}

type pattern struct {
	verb, path string
	s          *ServerHandler
}

func (p pattern) WithMiddleware(m Middleware, ms ...Middleware) pattern {
	handler, ok := p.s.mapping[p]
	if !ok {
		panic("handler is not registered")
	}

	handler.ms = append(handler.ms, m)
	handler.ms = append(handler.ms, ms...)

	return p
}
