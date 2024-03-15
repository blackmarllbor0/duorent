package models

import "time"

type User struct {
	ID            uint64    `sql:"id"`
	NationalityID uint64    `sql:"nationality_id"`
	FullName      string    `sql:"full_name"`
	Email         string    `sql:"email"`
	PhoneNumber   string    `sql:"phone_number"`
	Gender        bool      `sql:"gender"`
	IsDeleted     bool      `sql:"is_deleted"`
	DateOfBirth   time.Time `sql:"date_of_birth"`
	CreateDate    time.Time `sql:"create_date"`
	UpdateDate    time.Time `sql:"update_date"`
}
