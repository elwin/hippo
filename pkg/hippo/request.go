package hippo

import (
	"net/http"
	"net/url"
)

type Request interface {
	URL() url.URL
	Header() http.Header
	Form() url.Values
	PostForm() url.Values
}

type BaseRequest struct {
	url      url.URL
	header   http.Header
	form     url.Values
	postForm url.Values
}

func (b BaseRequest) URL() url.URL {
	return b.url
}

func (b BaseRequest) Header() http.Header {
	return b.header
}

func (b BaseRequest) Form() url.Values {
	return b.form
}

func (b BaseRequest) PostForm() url.Values {
	return b.postForm
}
