package main

import (
	"context"
	"errors"
	"github.com/xiangrui2019/redis"
	"strconv"
)

func limiter(ctx context.Context, client redis.Client, id string, limit int64, duration int64) error {
	var sum int64
	var err error

	data, err := client.Get(ctx, id)

	if err != nil {
		sum = 1
		client.Set(ctx, redis.Item{})
	}

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