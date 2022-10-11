package model

import (
	"fmt"
	"gofaka/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPasswd,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	//fmt.Println(dsn)
	//dsn = "kitnoob:ZAd3F8FC7YKCSiZ7@tcp(127.0.0.1:3306)/gofaka?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("test")
	if err != nil {
		log.Printf("Failed to connected database %s", err)
	}

	//自动d
	err := db.AutoMigrate(&User{}, &Category{}, &Item{}, &Order{})
	if err != nil {
		return
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复t用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

}

func GetDb() *gorm.DB {
	return db
}
