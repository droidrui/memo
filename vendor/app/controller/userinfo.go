package controller

import (
	"app/route"
	"net/http"
	"app/database"
	"app/route/middleware"
	"github.com/gorilla/context"
	"log"
	"app/response"
	"app/response/errcode"
)

func init() {
	route.PUT("/userinfo", middleware.VerifyToken(update))
}

func update(w http.ResponseWriter, r *http.Request) {
	uid := context.Get(r, "uid")
	nickname := r.FormValue("nickname")
	gender := r.FormValue("gender")
	headUrl := r.FormValue("headUrl")
	if nickname != "" {
		_, err := database.SQL.Exec("UPDATE user SET nickname=? WHERE id=?", nickname, uid)
		if err != nil {
			log.Println(err)
			response.SendError(w, errcode.ServerError)
			return
		}
		response.SendSuccess(w, nil)
		return
	}
	if gender != "" {
		_, err := database.SQL.Exec("UPDATE user SET gender=? WHERE id=?", gender, uid)
		if err != nil {
			log.Println(err)
			response.SendError(w, errcode.ServerError)
			return
		}
		response.SendSuccess(w, nil)
		return
	}
	if headUrl != "" {
		_, err := database.SQL.Exec("UPDATE user SET headUrl=? WHERE id=?", headUrl, uid)
		if err != nil {
			log.Println(err)
			response.SendError(w, errcode.ServerError)
			return
		}
		response.SendSuccess(w, nil)
	}
}
