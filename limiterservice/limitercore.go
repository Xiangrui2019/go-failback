package limiterservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/xiangrui2019/redis"
	"strconv"
)

// Limiter(上下文对象, redis客户端, 唯一标示为一个clientid, 单次duration下和单个id下可以进入的次数, 时间)

func Limiter(serviceName string, ctx context.Context, client redis.Client, clientid string, limit int64, duration int32) error {
	var sum int64
	var err error

	data, err := client.Get(ctx, clientid)

	if data == nil {
		sum = 1
		sumstring := strconv.Itoa(int(sum))

		err := client.Set(ctx, &redis.Item{
			Key:   clientid,
			Value: []byte(sumstring),
			TTL:   duration,
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
			sum, err = client.IncrBy(ctx, clientid, 1)

			if err != nil {
				panic(err)
			}
		}
	}

	return nil
}
