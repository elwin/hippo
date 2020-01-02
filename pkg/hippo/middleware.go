package hippo

import (
	"fmt"
	"net/http"
	"time"
)

const sessionID = "sessionID"

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

func (s ServerHandler) SessionMiddleware(next HandlerFunc) HandlerFunc {
	return func(request Request) Response {
		sessionCookie := request.cookies.Get(sessionID)
		if sessionCookie != nil {
			if session, ok := s.sessions.Get(sessionCookie.Value); ok {
				request.Session = session
				return next(request)
			}
		}

		session, err := s.sessions.New()
		if err != nil {
			return NewResponse().WithStatusCode(http.StatusInternalServerError)
		}

		fmt.Printf("New Session with id %s\n", session.ID())

		request.Session = session
		response := next(request)
		response.WithCookie(http.Cookie{Name: sessionID, Value: session.ID()})
		return response
	}
}
