package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Data struct {
	Db *gorm.DB
}

var database *gorm.DB

func NewDatabase(uri string) (Data, error) {
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal("Error to connect the db")
	}
	data := Data{
		Db: db,
	}
	database = data.Db
	return data, nil
}
