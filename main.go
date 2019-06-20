package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/xiangrui2019/redis"
	"net/http"
)

const IPMAXConnection = 2
const Secound = 1

var redisconn redis.Client
var ctx context.Context



func main() {
	app := gin.Default()
	ctx = context.Background()
	redisconn = redis.New(redis.Options{
		Address: "localhost:6379",
		PoolSize: 10,
	})

	err := redisconn.Ping(ctx)

	if err != nil {
		panic(err)
	}

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
