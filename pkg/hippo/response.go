package hippo

import "net/http"

type Response interface {
	StatusCode() int
	Header() map[string][]string
	Body() string
}

type BaseResponse struct {
	body    string
	cookies []http.Cookie
}

func (b BaseResponse) StatusCode() int {
	return http.StatusOK
}

func (b BaseResponse) Header() map[string][]string {
	header := map[string][]string{}
	for _, cookie := range b.cookies {
		header["Set-Cookie"] = append(header["Set-Cookie"], cookie.String())
	}

	return header
}

func (b BaseResponse) Body() string {
	return b.body
}

func (b *BaseResponse) AddCookie(cookie http.Cookie) {
	b.cookies = append(b.cookies, cookie)
}

func (b *BaseResponse) SetBody(body string) {
	b.body = body
}
