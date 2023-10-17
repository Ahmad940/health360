package model

import (
	"time"
)

type Consultant struct {
	ID string `json:"id" gorm:"primaryKey; type:varchar; not null; unique"`

	UserID          string   `json:"user_id" gorm:"type:varchar;index"`
	Bio             string   `json:"full_name" gorm:"type:varchar"`
	Services        []string `json:"services" gorm:"type:varchar[]"`
	Specializations []string `json:"specializations" gorm:"type:varchar[]"`

	User User `json:"user"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ModifyConsultantParam struct {
	ID              string   `json:"id"`
	UserID          string   `json:"user_id" validate:"required"`
	Bio             string   `json:"full_name"`
	Services        []string `json:"services"`
	Specializations []string `json:"specializations"`
}
