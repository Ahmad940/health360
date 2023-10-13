package model

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
