package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func loadDBConfig() (*gorm.DB, error) {
	dbCli, err := gorm.Open("mysql", "root:112358@(127.0.0.1:3306)/kitchen?charset=utf8&parseTime=True")
	if err != nil {
		return nil, err
	}
	return dbCli, nil
}

func loadRedisConfig() (*redis.Client, error) {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := redisCli.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return redisCli, nil
}
