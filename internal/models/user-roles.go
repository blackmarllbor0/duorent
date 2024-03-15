package models

import "time"

type UserRoles struct {
	ID         uint64    `sql:"id"`
	UserID     uint64    `sql:"user_id"`
	RoleId     uint64    `sql:"role_id"`
	IsDeleted  bool      `sql:"is_deleted"`
	CreateDate time.Time `sql:"create_date"`
	UpdateDate time.Time `sql:"update_date"`
}
