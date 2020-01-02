package hippo

import "net/http"

type redirect struct {
	baseResponse
}

func NewRedirect(url string) redirect {
	return redirect{
		NewResponse().
			WithStatusCode(http.StatusSeeOther).
			WithHeader("Location", url),
	}
}
