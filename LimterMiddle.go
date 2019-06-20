package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/xiangrui2019/redis"
)

func LimiterMiddleware(ctx context.Context, client redis.Client, limit int64, duration int32) func(context *gin.Context) {
	return func(context *gin.Context) {
		err := limiter(ctx, client, context.ClientIP(), limit, duration)

		if err != nil {
			context.Status(400)
			return
		}

		context.Next()
	}
}