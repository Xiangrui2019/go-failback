package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	app := gin.Default()
	client, _ := redis.Dial("tcp", os.Getenv("REDIS_HOST"))

	app.GET("/", func(context *gin.Context) {
		err := limiter(client, context.ClientIP(), 2, 10)

		if err != nil {
			context.Status(409)
			return
		}

		context.String(200, "limit")
	})

	http := http.Server{
		Addr: "localhost:8080",
		Handler: app,
	}

	err := http.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
