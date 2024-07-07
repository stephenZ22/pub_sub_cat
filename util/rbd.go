package util

import (
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	rdb        *redis.Client
	clientOnce sync.Once
)

// type RedisConn struct {
// 	Addr string
// }

func CreateRedisConn() *redis.Client {

	clientOnce.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})
	})

	return rdb
}
