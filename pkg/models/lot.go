package models

import "gorm.io/gorm"

type Lot struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	File     int
	Place    int
	Avalible bool
	UserID   uint
}
