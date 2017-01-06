package server

import (
	"net/http"
	"fmt"
	"time"
	"log"
	"app/route"
)

func Run() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Running HTTP "+httpAddress())
	log.Fatal(http.ListenAndServe(httpAddress(), route.LoadHttp()))
}

func httpAddress() string {
	return ":80"
}
