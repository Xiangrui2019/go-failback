package main

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

const IPMAXConnection = 2
const Secound = 1

func limiter(id string, client redis.Conn) (int64, error) {
	var sum int64
	var err error

	data, _ := redis.String(client.Do("GET", id))

	if data == "" {
		sum, err = redis.Int64(client.Do("INCR", id))
		if err != nil {
			return 0, err
		}
		_, err = client.Do("EXPIRE", id, 1)
		if err != nil {
			return 0, err
		}

	} else {
		sum, err = strconv.ParseInt(data, 10, 0)
		if err != nil {
			return 0, err
		}
		if sum > 2 {
			return 0, errors.New("limit too big.")
		} else {
			sum, err = redis.Int64(client.Do("INCR", id))
			if err != nil {
				panic(err)
			}
		}
	}

	return sum, nil
}

func main() {
	app := gin.Default()
	client, _ := redis.Dial("tcp", os.Getenv("REDIS_HOST"))

	app.GET("/", func(context *gin.Context) {
		ip := context.ClientIP()

		_, err := limiter(ip, client)

		if err != nil {
			context.Status(404)
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
