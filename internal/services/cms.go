package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	Username string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type CmsApp struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewCmsApp() *CmsApp {
	app := &CmsApp{}
	connDB(app)
	connRdb(app)
	return app
}

func connDB(app *CmsApp) {
	// 读取配置文件信息
	config_data, err := os.ReadFile("common.config")
	if err != nil {
		panic(err)
	}
	var config Config
	if err := json.Unmarshal(config_data, &config); err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port)
	fmt.Println(dsn)
	mysqlDB, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)

	// mysqlDB = mysqlDB.Debug()
	app.db = mysqlDB
}

func connRdb(app *CmsApp) {
	// redis-cli
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	// return rdb
	app.rdb = rdb
}
