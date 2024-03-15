package models

import "time"

type UserHash struct {
	ID         uint64    `sql:"id"`
	UserID     uint64    `sql:"user_id"`
	Hash       string    `sql:"hash"`
	Salt       string    `sql:"salt"`
	UpdateDate time.Time `sql:"update_date"`
}
