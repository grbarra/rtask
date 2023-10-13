package repository

import (
	"fmt"

	"github.com/go-redis/redis"
)

type RedisConfig struct {
	Host string
	Port string
}

func NewRedisDB(cfg RedisConfig) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	client := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	ping, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(ping)

	return client, nil
}
