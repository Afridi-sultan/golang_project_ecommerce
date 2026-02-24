// test middleware
package middleware

import (
	"fmt"
	"net/http"
)
func Hudai(next http.Handler)http.Handler{
	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("This is hudai funciton or test function.")
		next.ServeHTTP(w,r)
	})
}