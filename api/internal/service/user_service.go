package service

import (
	"api/internal/model"
	"api/internal/repository"
)

type UserService interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{
		userRepo: repo,
	}
}

func (s *userService) Create(user *model.User) error {
	return s.userRepo.Create(user)
}

func (s *userService) FindByEmail(email string) (*model.User, error) {
	return s.userRepo.FindByEmail(email)
}
