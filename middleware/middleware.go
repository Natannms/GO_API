package middleware

import "net/http"

func HeaderMiddleare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(r http.ResponseWriter, rq *http.Request) {
		r.Header().Set("Content-Type", "application/json")
	})
}
