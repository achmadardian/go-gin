package repositories

import (
	"go-gin/config"
	"go-gin/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: DB,
	}
}

func (r *UserRepo) GetAll() ([]models.User, error) {
	var users []models.User

	err := config.DB.Model(&models.User{}).Select("id", "name").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
