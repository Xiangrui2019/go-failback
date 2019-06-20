package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func LimiterMiddleware(client redis.Conn, limit int64, duration int64) func(context *gin.Context) {
	return func(context *gin.Context) {
		err := limiter(client, context.ClientIP(), limit, duration)

		if err != nil {
			context.Status(404)
			return
		}

		context.Next()
	}
}