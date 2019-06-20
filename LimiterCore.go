package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/xiangrui2019/redis"
	"strconv"
)

func limiter(ctx context.Context, client redis.Client, id string, limit int64, duration int32) error {
	var sum int64
	var err error

	data, err := client.Get(ctx, id)

	if data == nil {
		sum = 1
		sumstring := strconv.Itoa(int(sum))

		err := client.Set(ctx, &redis.Item{
			Key: id,
			Value: []byte(sumstring),
			TTL: duration,
		})

		if err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	} else {
		sum, err = strconv.ParseInt(string(data.Value), 10, 0)

		if err != nil {
			return err
		}

		if sum >= limit {
			return errors.New("limit too big.")
		} else {
			sum, err = client.IncrBy(ctx, id, 1)

			if err != nil {
				panic(err)
			}
		}
	}

	return nil
}