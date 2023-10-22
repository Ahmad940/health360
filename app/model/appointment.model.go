package model

import (
	"time"
)

type Appointment struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id" validate:"required"`
	ConsultantID string    `json:"consultant_id" validate:"required"`
	Time         time.Time `json:"time" validate:"required"`

	User       *User       `json:"user" gorm:"foreignKey:UserID"`
	Consultant *Consultant `json:"consultant" gorm:"foreignKey:ConsultantID"`
}
