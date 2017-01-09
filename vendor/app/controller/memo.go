package controller

import (
	"app/route"
	"net/http"
	"app/route/middleware"
	"github.com/gorilla/context"
	"fmt"
	"app/model"
	"app/database"
	"app/response"
	"app/response/errcode"
	"log"
)

func init() {
	route.POST("/memo", middleware.VerifyToken(createMemo))
	route.DELETE("/memo", middleware.VerifyToken(deleteMemo))
	route.PUT("/memo", middleware.VerifyToken(updateMemo))
	route.GET("/memo", middleware.VerifyToken(getMemo))
}

func createMemo(w http.ResponseWriter, r *http.Request) {
	uid := context.Get(r, "uid")
	phone := context.Get(r, "phone")
	fmt.Println("uid=", uid, "phone=", phone)
	title := r.FormValue("title")
	content := r.FormValue("content")
	if title == "" || content == "" {
		response.SendError(w, errcode.ParamInvalid)
		return
	}
	fmt.Println("title=", title, "content=", content)
	result, err := database.SQL.Exec("INSERT INTO memo(title, content, uid) VALUES(?,?,?)", title, content, uid)
	if err != nil {
		log.Println(err)
		response.SendError(w, errcode.ServerError)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		response.SendError(w, errcode.ServerError)
		return
	}
	memo := model.Memo{}
	err = database.SQL.Get(&memo, "SELECT * FROM memo WHERE id=?", fmt.Sprint(id))
	if err != nil {
		log.Println(err)
		response.SendError(w, errcode.ServerError)
		return
	}
	response.SendSuccess(w, &memo)
}

func deleteMemo(w http.ResponseWriter, r *http.Request) {
	uid := context.Get(r, "uid")
	phone := context.Get(r, "phone")
	fmt.Println("uid=", uid, "phone=", phone)
}

func updateMemo(w http.ResponseWriter, r *http.Request) {
	uid := context.Get(r, "uid")
	phone := context.Get(r, "phone")
	fmt.Println("uid=", uid, "phone=", phone)
}

func getMemo(w http.ResponseWriter, r *http.Request) {
	var uid interface{} = r.FormValue("uid")
	if uid == "" {
		uid = context.Get(r, "uid")
		phone := context.Get(r, "phone")
		fmt.Println("uid=", uid, "phone=", phone)
	}
	fmt.Println("uid=", uid)
	var err error
	var list []model.Memo
	err = database.SQL.Select(&list, "SELECT * FROM memo WHERE uid=? ORDER BY createdAt DESC", uid)
	if err != nil {
		log.Println(err)
		response.SendError(w, errcode.ServerError)
		return
	}
	if len(list) == 0 {
		response.SendSuccess(w, "[]")
		return
	}
	response.SendSuccess(w, &list)
}
