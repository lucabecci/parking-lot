package internal

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Data struct {
	Db *gorm.DB
}

var database *gorm.DB

func ConnectDB(uri string) error {
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), //logger blocked
	})

	if err != nil {
		fmt.Println("Error to connect you DB")
		return err
	}
	data := Data{Db: db}
	fmt.Println("DB is connected")
	database = data.Db
	return nil
}

func (d *Data) Migrations() error {
	return nil
}
