package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

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

type Account struct {
	Id       int64     `gorm:"column:id;primary_key"`
	UserID   string    `gorm:"column:user_id"`
	Password string    `gorm:"column:password"`
	Nickname string    `gorm:"column:nickname"`
	Ct       time.Time `gorm:"column:created_at"`
	Ut       time.Time `gorm:"column:updated_at"`
}

func (a Account) TableName() string {
	return "cms_account.account"
}

// func main() {
// 	db := connDB()
// 	// var accounts []Account
// 	var account Account
// 	// if err := db.Find(&accounts).Error; err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	if err := db.Where("id=?", 2).First(&account).Error; err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(account)

// }

func connDB() *gorm.DB {
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

	mysqlDB = mysqlDB.Debug()
	return mysqlDB
}
