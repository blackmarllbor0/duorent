package postgres

import (
	"database/sql"
	"duorent.ru/internal/models"
	"duorent.ru/internal/repository"
	"fmt"
)

type userRepo struct {
	pool *sql.DB
}

func NewUserRepo(pool *sql.DB) repository.UserRepo {
	return &userRepo{pool: pool}
}

func (ur *userRepo) GetAllUsers(limit uint) ([]models.User, error) {
	var (
		err  error
		rows *sql.Rows
	)

	if limit == 0 {
		rows, err = ur.pool.Query("select u.* from public.users u order by u.full_name")
	} else {
		rows, err = ur.pool.Query(" select u.* from public.users u order by u.full_name limit $1", limit)
	}

	if err != nil {
		return nil, fmt.Errorf("pg: error with get all users: %v", err)
	}

	defer func() {
		err = rows.Close()
	}()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user); err != nil {
			return nil, fmt.Errorf("pg: error with scan users: %v", err)
		}

		users = append(users, user)
	}

	return users, err
}
