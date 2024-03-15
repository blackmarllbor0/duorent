package models

import "time"

type Role struct {
	ID         uint64    `sql:"id"`
	Role       string    `sql:"role"`
	IsDeleted  bool      `sql:"is_deleted"`
	CreateDate time.Time `sql:"create_date"`
	UpdateDate time.Time `sql:"update_date"`
}
