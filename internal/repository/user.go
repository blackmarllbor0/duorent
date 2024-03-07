package repository

import "duorent.ru/internal/models"

type UserRepo interface {
	GetAllUsers(limit uint) ([]models.User, error)
}
