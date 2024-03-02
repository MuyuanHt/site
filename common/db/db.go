package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// InitMysql 初始化数据库连接
func InitMysql(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

// CloseMysql 断开数据库连接
func CloseMysql(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println(err)
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
