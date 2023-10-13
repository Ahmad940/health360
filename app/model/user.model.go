package model

import (
	"database/sql"
	"time"
)

type UserRole string

const (
	AdminRoleAdmin      = UserRole("admin")
	ConsultantRoleAdmin = UserRole("consultant")
	UserRoleAdmin       = UserRole("user")
)

type User struct {
	ID string `json:"id" gorm:"primaryKey; type:varchar; not null; unique"`

	FullName    string `json:"full_name" gorm:"type:varchar; unique"`
	Country     string `json:"country" gorm:"type:varchar; index" validate:"required"`
	CountryCode string `json:"country_code" gorm:"type:varchar; not null" validate:"required"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar; not null" validate:"required"`

	Role UserRole `json:"role" gorm:"type:varchar; check:role IN ('admin', 'user'); not null; default:user"`

	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index"`
}

type Auth struct {
	ID          string `json:"-"`
	Country     string `json:"country" validate:"required"`
	CountryCode string `json:"country_code" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type Login struct {
	CountryCode string `json:"country_code" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	OTP         string `json:"otp" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
