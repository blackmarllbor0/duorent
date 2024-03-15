package postgres

import (
	"database/sql"
	"duorent.ru/internal/models"
	"duorent.ru/internal/repository"
	"fmt"
)

type userRepo struct {
	conn repository.SQLConnection
}

func NewUserRepo(conn repository.SQLConnection) repository.UserRepo {
	return &userRepo{conn: conn}
}

func (ur *userRepo) GetAllUsers(limit uint) ([]models.User, error) {
	pool, err := ur.conn.GetConnection()
	if err != nil {
		return nil, err
	}

	var rows *sql.Rows

	if limit == 0 {
		rows, err = pool.Query("select u.* from public.users u order by u.full_name")
	} else {
		rows, err = pool.Query(" select u.* from public.users u order by u.full_name limit $1", limit)
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

	ur.conn.ReleaseConnection()

	return users, err
}
