package controller

import (
	"app/route"
	"net/http"
)

func init() {
	route.GET("/hello", hello)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
