package lot

import "gorm.io/gorm"

type Lot struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Section  string
	Place    string
	Avalible bool
	UserID   uint
}
