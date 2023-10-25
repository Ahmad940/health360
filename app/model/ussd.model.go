package model

type USSDBody struct {
	SessionID   string
	ServiceCode string
	PhoneNumber string
	Text        string
}
