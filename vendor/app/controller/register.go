package controller

import (
	"app/route"
	"net/http"
	"app/response"
	"app/response/errcode"
	"app/model"
	"app/database"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
	"app/constant"
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
	var err error
	user := model.User{}
	err = database.SQL.Get(&user, "SELECT id FROM user WHERE phone=? LIMIT 1", phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			createUser(w, phone, password)
			return
		}
		log.Println(err)
		response.SendError(w, errcode.ServerError)
		return
	}
	response.SendError(w, errcode.UserRegistered)
}

func createUser(w http.ResponseWriter, phone, password string) {
	var err error
	_, err = database.SQL.Exec("INSERT INTO user(phone, password) VALUES(?,?)", phone, password)
	if err != nil {
		log.Println(err)
		response.SendError(w, errcode.ServerError)
		return
	}
	user := model.User{}
	err = database.SQL.Get(&user, "SELECT * FROM user WHERE phone=? LIMIT 1", phone)
	if err != nil {
		log.Println(err)
		response.SendError(w, errcode.ServerError)
		return
	}
	//生成accessToken
	accessJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":  user.ID,
		"phone":user.Phone,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	accessToken, err := accessJWT.SignedString([]byte(constant.AccessTokenKey))
	if err != nil {
		log.Println(err)
		response.SendError(w, errcode.ServerError)
		return
	}
	//生成refreshToken
	refreshJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":  user.ID,
		"phone":user.Phone,
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	refreshToken, err := refreshJWT.SignedString([]byte(constant.RefreshTokenKey))
	if err != nil {
		log.Println(err)
		response.SendError(w, errcode.ServerError)
		return
	}
	loginInfo := model.LoginInfo{
		AccessToken: accessToken,
		RefreshToken:refreshToken,
		User:        user,
	}
	response.SendSuccess(w, &loginInfo)
}
