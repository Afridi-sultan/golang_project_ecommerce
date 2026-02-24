package middleware

import "net/http"

// middleware setup

func CorsMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, OPTIONS")
		w.Header().Set("content-type", "application/json")

		next.ServeHTTP(w, r)
	})
}
