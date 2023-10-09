package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Ahmad940/health360/app/model"
	"github.com/Ahmad940/health360/pkg/util"
	"github.com/Ahmad940/health360/platform/db"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var invalidCred = "Invalid email or password"

// Login
func Login(param model.Auth) (string, error) {
	var user model.User

	err := db.DB.Where("username LIKE ?", param.UserName).First(&user).Error
	if err != nil {
		if SqlErrorNotFound(err) {
			return "", errors.New(invalidCred)
		}
		return "", err
	}
	return "", nil
}

// CreateAccount
func CreateAccount(param model.Auth) error {
	var user model.User

	err := db.DB.Where("username = ?", param.UserName).First(&user).Error
	if SqlErrorIgnoreNotFound(err) != nil {
		fmt.Println("Yo")
		return err
	}

	// checking if user is registered or not
	if (user != model.User{}) {
		return errors.New("Username in use")
	}

	err = db.DB.Create(&model.User{
		ID: gonanoid.Must(),
	}).Error

	return nil
}
