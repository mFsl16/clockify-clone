package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Mysql gorm.DB
}

func NewDatabase() *Database {

	mysqlDsn := "user:secret@tcp(127.0.0.1:3306)/learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db: " + err.Error())
	}

	return &Database{
		Mysql: *db,
	}
}
