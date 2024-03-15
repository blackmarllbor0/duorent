package models

import "time"

type City struct {
	ID            uint64    `sql:"id"`
	CountryId     uint64    `sql:"country_id"`
	Title         string    `sql:"title"`
	LinkToFlagImg string    `sql:"link_to_flag_img"`
	IsDeleted     bool      `sql:"is_deleted"`
	CreateDate    time.Time `sql:"create_date"`
	UpdateDate    time.Time `sql:"update_date"`
}
