package service

import (
	"context"
	"database/sql"
	"duorent.ru/internal/models"
	"duorent.ru/internal/repository"
	"duorent.ru/internal/transport/rest/dto"
	"fmt"
	"strings"
)

type UserService interface {
	GetAllUsers(ctx context.Context, limit uint) ([]models.User, error)
	CreateUser(ctx context.Context, userDTO dto.CreateUserDTO) (uint64, error)
}

type userService struct {
	userRepo        repository.UserRepo
	userHashService UserHashService
}

func NewUserService(userRepo repository.UserRepo, userHashService UserHashService) UserService {
	return &userService{
		userRepo:        userRepo,
		userHashService: userHashService,
	}
}

func (us userService) GetAllUsers(ctx context.Context, limit uint) ([]models.User, error) {
	return us.userRepo.GetAllUsers(ctx, limit)
}

func (us userService) CreateUser(ctx context.Context, userDTO dto.CreateUserDTO) (uint64, error) {
	_, err := us.userRepo.GetByEmail(ctx, userDTO.Email)
	containErrMsgOfUserExisting := strings.ReplaceAll(err.Error(), sql.ErrNoRows.Error(), "")
	if err != nil && containErrMsgOfUserExisting == "" {
		return 0, fmt.Errorf("srvc: user: failed to get user by email: %v", err)
	}

	user := models.User{
		NationalityID: userDTO.NationalityId,
		FullName:      userDTO.FullName,
		Email:         userDTO.Email,
		PhoneNumber:   userDTO.PhoneNumber,
		Gender:        userDTO.Gender,
		DateOfBirth:   userDTO.DateOfBirth,
	}

	createdUserID, err := us.userRepo.CreateUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("srvc: user: failed to create user: %v", err)
	}

	if err := us.userHashService.CreateUserHash(ctx, createdUserID, userDTO.Pwd); err != nil {
		return 0, fmt.Errorf("srvc: user: failed to create user hash: %v", err)
	}

	return createdUserID, nil
}
