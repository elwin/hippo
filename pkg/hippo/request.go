package hippo

import (
	"net/http"
	"net/url"
)

type Request struct {
	url      url.URL
	header   http.Header
	form     url.Values
	postForm url.Values
	cookies  CookieJar
	Session  Session
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

func (r Request) Query(key string) (string, bool) {
	value := r.url.Query().Get(key)
	if value == "" {
		return "", false
	}

	return value, true
}
