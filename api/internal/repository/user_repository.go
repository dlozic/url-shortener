package repository

import (
	"api/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(url *model.User) error
	FindByEmail(email string) (*model.User, error)
}

type gormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}

func (repo *gormUserRepository) Create(user *model.User) error {
	return repo.db.Create(user).Error
}

func (repo *gormUserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := repo.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}
