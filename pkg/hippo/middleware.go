package hippo

import (
	"fmt"
	"net/http"
	"time"

	"github.com/elwin/hippo/pkg/crypto"
)

type Middleware func(HandlerFunc) HandlerFunc

func LogMiddleware(next HandlerFunc) HandlerFunc {
	return func(request Request) Response {
		fmt.Printf("%s Received Request: %s\n", time.Now(), request.url.Path)
		return next(request)
	}
}

func TimeMiddleware(next HandlerFunc) HandlerFunc {
	return func(request Request) Response {
		start := time.Now()
		response := next(request)
		fmt.Println(time.Since(start))
		return response
	}
}

func SessionMiddleware(next HandlerFunc) HandlerFunc {
	return func(request Request) Response {
		sessionCookie := request.cookies.Get("sessionID")
		if sessionCookie != nil && sessionCookie.Value != "" {
			request.sessionID = sessionCookie.Value
			return next(request)
		}

		sessionID, err := crypto.GenerateRandomString(32)
		if err != nil {
			return NewResponse().WithStatusCode(http.StatusInternalServerError)
		}

		request.sessionID = sessionID
		response := next(request)
		response.WithCookie(http.Cookie{Name: "sessionID", Value: sessionID})
		return response
	}
}
