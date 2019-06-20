package main

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func limiter(client redis.Conn, id string, limit int64, duration int64) error {
	var sum int64
	var err error

	data, _ := redis.String(client.Do("GET", id))

	if data == "" {
		sum, err = redis.Int64(client.Do("INCR", id))

		if err != nil {
			return err
		}

		_, err = client.Do("EXPIRE", id, duration)

		if err != nil {
			return err
		}

	} else {
		sum, err = strconv.ParseInt(data, 10, 0)
		if err != nil {
			return err
		}
		if sum >= limit {
			return errors.New("limit too big.")
		} else {
			sum, err = redis.Int64(client.Do("INCR", id))
			if err != nil {
				panic(err)
			}
		}
	}

	return nil
}

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
