package hippo

import "net/http"

type Response struct {
	header     http.Header
	statusCode int
	body       string
}

func (r Response) Body() string {
	if r.body == "" {
		return http.StatusText(r.statusCode)
	}

	return r.body
}

func NewResponse() Response {
	return Response{
		statusCode: http.StatusOK,
		header:     http.Header{},
	}
}

func (r Response) WithBody(body string) Response {
	r.body = body
	return r
}

func (r Response) WithStatusCode(code int) Response {
	r.statusCode = code
	return r
}

func (r Response) WithCookie(cookie http.Cookie) Response {
	r.header.Add("Set-Cookie", (&cookie).String())
	return r
}
