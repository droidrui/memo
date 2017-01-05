package controller

import (
	"app/route"
	"net/http"
)

func init() {
	route.POST("/register", register)
}

func register(w http.ResponseWriter, r *http.Request) {

}
