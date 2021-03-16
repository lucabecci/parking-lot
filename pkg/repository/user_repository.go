package repository

import (
	"github.com/lucabecci/parking-lot/pkg"
	"github.com/lucabecci/parking-lot/pkg/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func (ur *UserRepository) Create(email, password string) (models.User, error) {
	user := models.User{
		Email:    email,
		Password: password,
	}

	user.HashPassword()
	result := ur.database.Create(&user)

	if result.RowsAffected < 1 {
		return models.User{}, pkg.ErrToCreate
	}

	result.Scan(&user)

	return user, nil
}

func (ur *UserRepository) GetByEmail(email string) (models.User, error) {
	var usr models.User
	result := ur.database.Where("email = ?", email).Find(&usr)
	if result.RowsAffected < 1 {
		return models.User{}, pkg.ErrEmailNotExists
	}
	return usr, nil
}

func (ur *UserRepository) GetByID(id uint) (models.User, error) {
	var usr models.User
	result := ur.database.Where("id = ?", id).Find(&usr)
	if result.RowsAffected < 1 {
		return models.User{}, pkg.ErrEmailNotExists
	}
	return usr, nil
}
