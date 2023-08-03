package bootstrap

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"tik-tok-server/global"
)

func InitializeRedis() *redis.Client {

	addr := fmt.Sprintf("%s:%d", global.App.Config.Redis.Host, global.App.Config.Redis.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: global.App.Config.Redis.Password, // no password set
		DB:       global.App.Config.Redis.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.App.Log.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return client
}
