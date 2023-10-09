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

	FullName    string `json:"full_name" gorm:"type:varchar; not null; unique" validate:"required"`
	CountryCode string `json:"country_code" gorm:"type:varchar; not null" validate:"required"`
	Phone       string `json:"phone" gorm:"type:int; not null" validate:"required"`

	Role UserRole `json:"role" gorm:"type:user_role; not null; default:user"`

	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index"`
}

type Auth struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}
