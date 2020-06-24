package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type Component struct {
	DB    *gorm.DB
	Cache *redis.Client
}

var Com *Component

func InitComponent() error {
	if Com == nil {
		Com = new(Component)
	}
	dbCli, err := loadDBConfig()
	if err != nil {
		return err
	}
	Com.DB = dbCli
	redisCli, err := loadRedisConfig()
	if err != nil {
		return err
	}
	Com.Cache = redisCli
	return nil
}
