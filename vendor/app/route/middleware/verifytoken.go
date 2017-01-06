package middleware

import (
	"net/http"
	"app/response"
	"app/response/errcode"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"app/constant"
	"time"
	"log"
	"github.com/gorilla/context"
)

func VerifyToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("x-access-token")
		if accessToken == "" {
			response.SendError(w, errcode.AccessTokenInvalid)
			return
		}
		token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(constant.AccessTokenKey), nil
		})
		if err != nil {
			response.SendError(w, errcode.AccessTokenInvalid)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			uid := claims["uid"]
			phone := claims["phone"]
			exp := int64(claims["exp"].(float64))
			if exp <= time.Now().Unix() {
				log.Println(uid, phone, "accessToken timeout:", exp)
				response.SendError(w, errcode.AccessTokenTimeout)
				return
			}
			context.Set(r, "uid", uid)
			context.Set(r, "phone", phone)
			next.ServeHTTP(w, r)
		} else {
			response.SendError(w, errcode.AccessTokenInvalid)
		}
	})
}
