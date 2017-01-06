package response

import (
	"net/http"
	"encoding/json"
	"app/response/errcode"
	"time"
	"fmt"
)

type Response struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data interface{} `json:"data,omitempty"`
}

func SendError(w http.ResponseWriter, code int) {
	var res = &Response{
		Code:code,
		Msg: errcode.Err[code],
		Time:fmt.Sprint(time.Now().Unix()),
	}
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "JSON Marshal Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	var res = &Response{
		Code:errcode.Success,
		Msg: errcode.Err[errcode.Success],
		Time:fmt.Sprint(time.Now().Unix()),
		Data:data,
	}
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "JSON Marshal Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
