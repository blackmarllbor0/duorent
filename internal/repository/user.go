package repository

import (
	"context"
	"duorent.ru/internal/models"
)

type UserRepo interface {
	GetAllUsers(ctx context.Context, limit uint) (res []models.User, err error)
	CreateUser(ctx context.Context, user models.User) (id uint64, err error)
	GetByEmail(ctx context.Context, email string) (models.User, error)
}
