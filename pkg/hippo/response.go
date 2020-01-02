package hippo

import "net/http"

type Response interface {
	Body() string
	Header() http.Header
	StatusCode() int

	WithBody(string) baseResponse
	WithStatusCode(int) baseResponse
	WithHeader(key, value string) baseResponse
	WithCookie(http.Cookie) baseResponse
}

type baseResponse struct {
	header     http.Header
	statusCode int
	body       string
}

func NewResponse() baseResponse {
	return baseResponse{
		statusCode: http.StatusOK,
		header:     http.Header{},
	}
}

func (r baseResponse) Body() string {
	if r.body == "" {
		return http.StatusText(r.statusCode)
	}

	return r.body
}

func (r baseResponse) Header() http.Header {
	return r.header
}

func (r baseResponse) StatusCode() int {
	return r.statusCode
}

func (r baseResponse) WithBody(body string) baseResponse {
	r.body = body
	return r
}

func (r baseResponse) WithStatusCode(code int) baseResponse {
	r.statusCode = code
	return r
}

func (r baseResponse) WithCookie(cookie http.Cookie) baseResponse {
	return r.WithHeader("Set-Cookie", (&cookie).String())
}

func (r baseResponse) WithHeader(key, value string) baseResponse {
	r.header.Add(key, value)
	return r
}
