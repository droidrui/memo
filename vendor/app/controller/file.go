package controller

import (
	"app/route"
	"net/http"
	"io"
	"app/response"
	"app/response/errcode"
	"log"
	"fmt"
	"app/route/middleware"
	"os"
)

func init() {
	route.POST("/file", middleware.VerifyToken(upload))
}

func upload(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var err error
	_, err = os.Create("test.jpg")
	if err != nil {
		log.Println(err)
		response.SendError(w, errcode.ServerError)
		return
	}
	f, err := os.OpenFile("test.jpg", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		response.SendError(w, errcode.ServerError)
		return
	}
	defer f.Close()
	nbytes, err := io.Copy(f, r.Body)
	if err != nil {
		log.Println(err)
		response.SendError(w, errcode.ServerError)
		return
	}
	fmt.Println("nbytes=", nbytes)
	response.SendSuccess(w, nil)
}
