package middlewares

import (
	"log"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func BuildMiddlewareChain(f http.Handler, m ...Middleware) http.Handler {
	if len(m) == 0 {
		return f
	}
	return m[0](BuildMiddlewareChain(f, m[1:cap(m)]...))
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Host, r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func PublicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("public route")
		next.ServeHTTP(w, r)
	})
}

func PrivateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("private route")
		next.ServeHTTP(w, r)
	})
}
