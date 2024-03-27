package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type SQLConnection interface {
	GetConnection() (*pgxpool.Conn, error)
}
