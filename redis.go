package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

var _redisClient *redis.Client

func connectRedis() *redis.Client {
	if _redisClient != nil {
		return _redisClient
	}

	REDIS_HOST := env("REDIS_HOST", "localhost")
	REDIS_PORT := env("REDIS_PORT", "6379")
	REDIS_PASS := env("REDIS_PASS", "")
	REDIS_USER := env("REDIS_USER", "default")
	REDIS_ADDR := fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT)

	_redisClient = redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDR,
		Username: REDIS_USER,
		Password: REDIS_PASS,
	})

	return _redisClient
}
