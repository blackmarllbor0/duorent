package postgres

import (
	"context"
	"duorent.ru/internal/repository"
	"fmt"
)

type userHashRepo struct {
	conn repository.SQLConnection
}

func NewUserHashRepo(conn repository.SQLConnection) repository.UserHashRepo {
	return &userHashRepo{conn: conn}
}

func (uhr userHashRepo) CreateUserHash(ctx context.Context, userID uint64, hash, salt string) (id uint, err error) {
	conn, err := uhr.conn.GetConnection()
	if err != nil {
		return 0, fmt.Errorf("pg: user-hash: failed to get conn: %v", err)
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return 0, err
	}

	if err := tx.QueryRow(
		ctx,
		`INSERT INTO public.users_hash (user_id, hash, salt) VALUES ($1, $2, $3) RETURNING id`,
		userID,
		hash,
		salt,
	).Scan(&id); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return 0, fmt.Errorf("pg: user-hash: failed to rollback transaction: %v", err)
		}

		return 0, fmt.Errorf("pg: user-hash: failed to write date: %v", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("pg: user-hash: failed to commit transaction: %v", err)
	}

	return id, nil
}
