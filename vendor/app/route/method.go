package route

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/context"
)

func GET(path string, fn http.HandlerFunc) {
	router.GET(path, handlerFunc(fn))
}

func POST(path string, fn http.HandlerFunc) {
	router.POST(path, handlerFunc(fn))
}

func DELETE(path string, fn http.HandlerFunc) {
	router.DELETE(path, handlerFunc(fn))
}

func PUT(path string, fn http.HandlerFunc) {
	router.PUT(path, handlerFunc(fn))
}

func handlerFunc(h http.HandlerFunc) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		context.Set(r, "params", p)
		h.ServeHTTP(w, r)
	})
}
