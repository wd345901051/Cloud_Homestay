package models

import (
	"fmt"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = Init()
var RDB = InitRedisDB()

func Init() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/cloud_homestay?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("LINKDB ERROR:", err)
	}
	return DB
}

func InitRedisDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "192.168.188.100:6379",
		Password: "root",
		DB:       0,
	})
}
