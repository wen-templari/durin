package util

import (
	"flag"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	Pool        *redis.Pool
	redisServer = flag.String("redisServer", os.Getenv("REDIS_ADDRESS"), "")
)

func InitPool() {
	Pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", *redisServer) },
	}
}
