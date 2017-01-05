package logrequest

import (
	"net/http"
	"fmt"
	"time"
)

func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
