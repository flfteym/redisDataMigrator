package utils

import (
	"time"

	"github.com/go-redis/redis/v8"
)

var cluster *redis.ClusterClient

func redisInit(url []string) {
	cluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        url,
		DialTimeout:  100 * time.Microsecond,
		ReadTimeout:  100 * time.Microsecond,
		WriteTimeout: 100 * time.Microsecond,
	})
}
