package main

import (
	"runtime"
	"app/database"
	"app/controller"
	"app/server"
	"app/route"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	database.Connect()
	controller.Load()
	server.Run(route.LoadHttp())
}
