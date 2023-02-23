package initialize

import (
	"backend-go/internal/app/config"
	"context"
	"github.com/redis/go-redis/v9"
	"log"

	"go.uber.org/zap"
)

func NewRedis(c *config.Redis) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password, // no password set
		DB:       c.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("redis connect ping failed, err:", zap.Error(err))
	} else {
		log.Println("redis connect ping response:", zap.String("pong", pong))
	}
	return client
}
