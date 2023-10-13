package model

import (
	"database/sql"
	"time"
)

type UserRole string

const (
	AdminRoleAdmin      = string("admin")
	ConsultantRoleAdmin = string("consultant")
	UserRoleAdmin       = string("user")
)

type User struct {
	ID string `json:"id" gorm:"primaryKey; type:varchar; not null; unique"`

	FullName    string `json:"full_name" gorm:"type:varchar; unique"`
	Country     string `json:"country" gorm:"type:varchar; index" validate:"required"`
	CountryCode string `json:"country_code" gorm:"type:varchar; not null" validate:"required"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar; not null" validate:"required"`

	Role string `json:"role" gorm:"type:varchar; check:role IN ('admin', 'user'); not null; default:user"`

	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
}

type UpdateUser struct {
	ID string `json:"id" validate:"required"`

	FullName    string `json:"full_name"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateUserAdmin struct {
	ID          string `json:"id" validate:"required"`
	FullName    string `json:"full_name"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role" gorm:"type:varchar; check:role IN ('admin', 'user'); not null; default:user"`
}
