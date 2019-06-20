package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/xiangrui2019/go-failback/limiterservice"
	"github.com/xiangrui2019/redis"
	"net/http"
	"os"
)

func main() {
	app := gin.Default()

	client := redis.New(redis.Options{
		Address: os.Getenv("REDIS_HOST"),
		PoolSize: 10,
	})
	ctx := context.Background()

	err := client.Ping(ctx)

	if err != nil {
		panic(err)
	}

	app.Use(limiterservice.LimiterMiddleware(ctx, client, 2, 10))

	app.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"code": 0,
			"message": "OK.",
		})
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