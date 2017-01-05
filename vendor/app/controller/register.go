package controller

import (
	"app/route"
	"net/http"
	"app/response"
	"app/response/errcode"
)

func init() {
	route.POST("/register", register)
}

func register(w http.ResponseWriter, r *http.Request) {
	phone := r.FormValue("phone")
	password := r.FormValue("password")
	if phone == "" || password == "" {
		response.SendError(w, errcode.ParamInvalid)
		return
	}

}
