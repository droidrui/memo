package route

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"app/route/middleware/logrequest"
	"github.com/gorilla/context"
)

var router *httprouter.Router

func init() {
	router = httprouter.New()
}

func LoadHttp() http.Handler {
	return middleware(router)
}

func middleware(h http.Handler) http.Handler {
	h = logrequest.Handler(h)
	h = context.ClearHandler(h)
	return h
}
