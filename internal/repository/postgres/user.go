package postgres

import (
	"context"
	"database/sql"
	"duorent.ru/internal/models"
	"duorent.ru/internal/repository"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type userRepo struct {
	conn repository.SQLConnection
}

func NewUserRepo(conn repository.SQLConnection) repository.UserRepo {
	return &userRepo{conn: conn}
}

func (ur *userRepo) GetAllUsers(ctx context.Context, limit uint) ([]models.User, error) {
	conn, err := ur.conn.GetConnection()
	if err != nil {
		return nil, fmt.Errorf("pg: user: failed to get conn")
	}

	var rows pgx.Rows

	if limit == 0 {
		rows, err = conn.Query(
			ctx,
			"SELECT u.* FROM public.users u WHERE u.is_deleted = $1 ORDER BY u.full_name",
			false,
		)
	} else {
		rows, err = conn.Query(
			ctx,
			" SELECT u.* FROM public.users u WHERE u.is_deleted = $1 ORDER BY u.full_name LIMIT $2",
			false,
			limit,
		)
	}

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, fmt.Errorf("pg: user: no users found: %v", err)
		}

		return nil, fmt.Errorf("pg: user: error with get all users: %v", err)
	}

	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user); err != nil {
			return nil, fmt.Errorf("pg: user: error with scan users: %v", err)
		}

		users = append(users, user)
	}

	return users, err
}

func (ur *userRepo) GetByEmail(ctx context.Context, email string) (models.User, error) {
	conn, err := ur.conn.GetConnection()
	if err != nil {
		return models.User{}, fmt.Errorf("pg: user: failed to get pool: %v", conn)
	}
	defer conn.Release()

	var user models.User
	if err := conn.QueryRow(
		ctx,
		"SELECT * FROM public.users WHERE is_deleted = $1 AND email = $2",
		false,
		email,
	).Scan(&user); err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return models.User{}, fmt.Errorf("pg: user: user by %s not found: %v", email, err)
		}

		return models.User{}, fmt.Errorf("pg: user: failed to get user by email: %v", err)
	}

	return user, nil
}

func (ur *userRepo) CreateUser(ctx context.Context, user models.User) (id uint64, err error) {
	conn, err := ur.conn.GetConnection()
	if err != nil {
		return 0, fmt.Errorf("pg: user: failed to get pool: %v", err)
	}

	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("pg: user: failed to run transaction: %v", err)
	}

	if err := tx.QueryRow(
		ctx,
		`INSERT INTO public.users (nationality_id, full_name, email, phone_number, gender, date_of_birth)
				VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		user.NationalityID,
		user.FullName,
		user.Email,
		user.PhoneNumber,
		user.Gender,
		user.DateOfBirth,
	).Scan(&id); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return 0, fmt.Errorf("pg: user: failed to cancel transaction: %v", err)
		}

		return 0, fmt.Errorf("pg: user: failed to insert user data: %v", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("pg: user: failed to commit transaction: %v", err)
	}

	return uint64(id), nil
}
