package hippo

import (
	"net/http"
	"net/url"
)

type Request struct {
	url       url.URL
	header    http.Header
	form      url.Values
	postForm  url.Values
	cookies   CookieJar
	sessionID string
}

type HandlerFunc func(Request) Response

type CookieJar []*http.Cookie

func (j CookieJar) Get(name string) *http.Cookie {
	for _, cookie := range j {
		if cookie.Name == name {
			return cookie
		}
	}

	return nil
}
