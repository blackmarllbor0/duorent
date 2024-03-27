package service

import (
	"context"
	"crypto/rand"
	"duorent.ru/internal/config"
	"duorent.ru/internal/repository"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserHashService interface {
	CreateUserHash(ctx context.Context, userID uint64, pwd string) error
}

type userHashService struct {
	userHashRepo repository.UserHashRepo
	cfgService   config.ConfigService
}

func NewUserHashService(
	userHashRepo repository.UserHashRepo,
	cfgService config.ConfigService,
) UserHashService {
	return &userHashService{
		userHashRepo: userHashRepo,
		cfgService:   cfgService,
	}
}

func (uhs userHashService) CreateUserHash(ctx context.Context, userID uint64, pwd string) error {
	userSalt := make([]byte, 16)
	_, err := rand.Read(userSalt)
	if err != nil {
		return fmt.Errorf("srvc: user-hash: failed to create user salt: %v", err)
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(pwd+string(userSalt)+uhs.cfgService.GetLocalHash().Salt),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return fmt.Errorf("srvc: user-hash: failed to generate hash: %v", err)
	}

	if _, err := uhs.userHashRepo.CreateUserHash(
		ctx,
		userID,
		string(hash),
		base64.URLEncoding.EncodeToString(userSalt),
	); err != nil {
		return fmt.Errorf("srvc: user-hash: failed to create user-hash: %v", err)
	}

	return nil
}
