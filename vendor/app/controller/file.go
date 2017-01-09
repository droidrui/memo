package controller

import (
	"app/route"
	"net/http"
)

func init() {
	route.POST("/file", upload)
}

func upload(w http.ResponseWriter, r *http.Request) {

}
