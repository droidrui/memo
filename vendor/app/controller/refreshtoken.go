package controller

import (
	"app/route"
	"net/http"
	"fmt"
	"app/response"
	"app/response/errcode"
	"github.com/dgrijalva/jwt-go"
	"app/constant"
	"log"
	"time"
	"app/model"
)

func init() {
	route.POST("/refreshToken", refreshToken)
}

func refreshToken(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("x-access-token")
	if accessToken == "" {
		response.SendError(w, errcode.AccessTokenInvalid)
		return
	}
	token1, _ := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(constant.AccessTokenKey), nil
	})
	if claims, ok := token1.Claims.(jwt.MapClaims); ok {
		uid1 := claims["uid"]
		phone1 := claims["phone"]

		refreshToken := r.FormValue("refreshToken")
		if refreshToken == "" {
			response.SendError(w, errcode.ParamInvalid)
			return
		}
		token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(constant.RefreshTokenKey), nil
		})
		if err != nil {
			log.Println(err)
			response.SendError(w, errcode.RefreshTokenInvalid)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			uid := claims["uid"]
			phone := claims["phone"]
			exp := int64(claims["exp"].(float64))
			if exp <= time.Now().Unix() {
				log.Println(uid, phone, "accessToken timeout:", exp)
				response.SendError(w, errcode.RefreshTokenTimeout)
				return
			}
			if uid1 != uid || phone1 != phone {
				log.Println("uid不相等")
				response.SendError(w, errcode.AccessTokenInvalid)
				return
			}
			//生成新的accessToken
			accessJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"uid":  uid,
				"phone":phone,
				"exp":  time.Now().Add(time.Hour * 24).Unix(),
			})
			accessToken, err := accessJWT.SignedString([]byte(constant.AccessTokenKey))
			if err != nil {
				log.Println(err)
				response.SendError(w, errcode.ServerError)
				return
			}
			//生成新的refreshToken
			refreshJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"uid":  uid,
				"phone":phone,
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
			}
			response.SendSuccess(w, &loginInfo)
		} else {
			response.SendError(w, errcode.RefreshTokenInvalid)
		}
	}

}
