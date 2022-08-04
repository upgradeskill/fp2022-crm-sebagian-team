package config

import (
	"fmt"

	"crm-sebagian-team/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConn() *utils.Conn {
	return &utils.Conn{
		GORM: initGorm(),
	}
}

func initGorm() *gorm.DB {
	username := "root"
	password := ""
	dbHost := "127.0.0.1"
	dbPort := 3306
	dbName := "bootcamp-majoo"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
