package repository

import "context"

type UserHashRepo interface {
	CreateUserHash(ctx context.Context, userID uint64, hash, salt string) (id uint, err error)
}
