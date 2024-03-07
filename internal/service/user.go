package service

import (
	"duorent.ru/internal/models"
	"duorent.ru/internal/repository"
)

type UserService interface {
	GetAllUsers(limit uint) ([]models.User, error)
}

type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &userService{userRepo: userRepo}
}

func (us userService) GetAllUsers(limit uint) ([]models.User, error) {
	return us.userRepo.GetAllUsers(limit)
}
