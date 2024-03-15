package models

import "time"

type Nationality struct {
	ID         uint64    `sql:"id"`
	Title      string    `sql:"title"`
	IsDeleted  bool      `sql:"is_deleted"`
	CreateDate time.Time `sql:"create_date"`
	UpdateDate time.Time `sql:"update_date"`
}
