package model

import (
	"time"

	"github.com/lib/pq"
)

type Consultant struct {
	ID string `json:"id" gorm:"primaryKey; type:varchar; not null; unique"`

	UserID          string         `json:"user_id" gorm:"type:varchar;index;unique"`
	Bio             string         `json:"bio" gorm:"type:varchar"`
	Services        pq.StringArray `json:"services" gorm:"type:varchar[]"`
	Specializations pq.StringArray `json:"specializations" gorm:"type:varchar[]"`

	User *User `json:"user" gorm:"foreignKey:UserID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AddConsultantParam struct {
	UserID          string         `json:"user_id" validate:"required"`
	Bio             string         `json:"bio"`
	Services        pq.StringArray `json:"services"`
	Specializations pq.StringArray `json:"specializations"`
}

type UpdateConsultantParam struct {
	ID string `json:"id"`

	Bio             string         `json:"bio"`
	Services        pq.StringArray `json:"services"`
	Specializations pq.StringArray `json:"specializations"`
}
