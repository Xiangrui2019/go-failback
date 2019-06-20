package main

import (
	"github.com/gin-gonic/gin"
	"github.com/garyburd/redigo/redis"
	"net/http"
)

const IPMAXConnection = 2
const Secound = 1

var client redis.Conn


func main() {
	app := gin.Default()

	app.GET("/", func(context *gin.Context) {
		ip := context.ClientIP()


	})

	http := http.Server{
		Addr: "localhost:8080",
		Handler: app,
	}

	err = http.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
