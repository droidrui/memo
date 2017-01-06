package route

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/gorilla/context"
	"app/route/middleware"
)

var router *httprouter.Router

func init() {
	router = httprouter.New()
}

func LoadHttp() http.Handler {
	return handle(router)
}

func handle(h http.Handler) http.Handler {
	h = middleware.LogRequest(h)
	h = context.ClearHandler(h)
	return h
}
