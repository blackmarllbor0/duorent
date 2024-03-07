package repository

import "database/sql"

type SQLConnection interface {
	GetConnection() (*sql.DB, error)
	ReleaseConnection()
}
