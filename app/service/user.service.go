package service

import (
	"github.com/Ahmad940/health360/app/model"
	"github.com/Ahmad940/health360/pkg/util"
	"github.com/Ahmad940/health360/platform/db"
)

// UpdateUserPassword update user password
func UpdateUserPassword(username string, password string) error {
	var user model.User

	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return err
	}

	return db.DB.Where("username = ?", username).First(&user).Update("password", hashedPassword).Error
}

// DeleteUser delete user
func DeleteUser(username string) error {
	var user model.User
	return db.DB.Unscoped().Where("username = ?", username).First(&user).Delete(&user).Error
}
