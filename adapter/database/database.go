package database

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
}

type GormDatabase struct {
	Db *gorm.DB
}

var lock = &sync.Mutex{}
var gormDatabaseInstance *GormDatabase

func startConnection() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3307)/pokedex-service?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database: ", err.Error())
		panic("Database error")
	}
	return db
}

func GetInstance() *GormDatabase {
	if gormDatabaseInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if gormDatabaseInstance == nil {
			db := startConnection()
			gormDatabaseInstance = &GormDatabase{Db: db}
		}
	}

	return gormDatabaseInstance
}
