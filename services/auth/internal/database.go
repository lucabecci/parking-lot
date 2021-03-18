package internal

import (
	"fmt"
	"log"

	"github.com/lucabecci/parking-lot/pkg/models"
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
	err = data.Migrations()
	if err != nil {
		log.Panic("Error to create the models.")
	}
	return data, nil
}

func (d *Data) Migrations() error {
	userTable := d.Db.Migrator().HasTable(&models.User{})
	lotTable := d.Db.Migrator().HasTable(&models.Lot{})
	if userTable == false || lotTable == false {
		err := d.Db.AutoMigrate(&models.User{})
		if err != nil {
			return err
		}
		err = d.Db.AutoMigrate(&models.Lot{})
		if err != nil {
			return err
		}
		fmt.Println("Models Created")
		return nil
	}
	return nil
}

func (d *Data) Close() error {
	d.Close()
	return nil
}
