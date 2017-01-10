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
	"time"
)

func init() {
	route.POST("/file", middleware.VerifyToken(upload))
}

func upload(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	mediaType := r.Header.Get("Content-Type")
	fmt.Println("mediaType=", mediaType)
	var err error
	file, err := os.Create(fmt.Sprintf("E:/file/%v.jpg", time.Now().Unix()))
	if err != nil {
		log.Println(err)
		response.SendError(w, errcode.ServerError)
		return
	}
	defer file.Close()
	nbytes, err := io.Copy(file, r.Body)
	if err != nil {
		log.Println(err)
		response.SendError(w, errcode.ServerError)
		return
	}
	fmt.Println("nbytes=", nbytes)
	response.SendSuccess(w, nil)
}
