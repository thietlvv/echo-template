package redis

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

var client *redis.Client

func newClient() *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	redisDb, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	redisPass := os.Getenv("REDIS_PASS")
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisHost, redisPort),
		Password: redisPass,
		DB:       redisDb,
	})

	return client
}
func GetClient() *redis.Client {
	if client == nil {
		return newClient()
	} else {
		return client
	}
}
