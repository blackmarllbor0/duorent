package dto

import "time"

type CreateUserDTO struct {
	NationalityId uint64    `json:"nationalityId,omitempty" validate:"number"`
	FullName      string    `json:"fullName" validate:"required"`
	Email         string    `json:"email,omitempty" validate:"required,email"`
	PhoneNumber   string    `json:"phoneNumber,omitempty" validate:"omitempty,e164"`
	Pwd           string    `json:"pwd" validate:"required,min=8"`
	Gender        bool      `json:"gender" validate:"required,boolean"`
	DateOfBirth   time.Time `json:"dateOfBirth,omitempty" validate:"required"`
}
