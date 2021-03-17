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

func ConnectDB(uri string) (Data, error) {
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), //logger blocked
	})

	if err != nil {
		fmt.Println("Error to connect you DB")
		return Data{}, err
	}
	data := Data{Db: db}
	fmt.Println("DB is connected")
	return data, nil
}

func (d *Data) Migrations() error {
	return nil
}
