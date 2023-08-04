package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"tik-tok-server/config"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
	DB          *gorm.DB
	Redis       *redis.Client
}

var App = new(Application)
