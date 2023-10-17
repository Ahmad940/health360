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

	User User `json:"user" gorm:"foreignKey:UserID;reference:ID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AddConsultantParam struct {
	UserID          string   `json:"user_id" validate:"required"`
	Bio             string   `json:"full_name"`
	Services        []string `json:"services"`
	Specializations []string `json:"specializations"`
}

type UpdateConsultantParam struct {
	ID              string   `json:"id" validate:"required"`
	Bio             string   `json:"full_name"`
	Services        []string `json:"services"`
	Specializations []string `json:"specializations"`
}
