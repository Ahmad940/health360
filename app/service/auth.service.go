package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Ahmad940/health360/app/model"
	"github.com/Ahmad940/health360/pkg/constant"
	"github.com/Ahmad940/health360/pkg/util"
	"github.com/Ahmad940/health360/platform/cache"
	"github.com/Ahmad940/health360/platform/db"
	"github.com/Ahmad940/health360/platform/sms"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// RequestOTP
func RequestOTP(param model.Auth) error {
	var user model.User

	err := db.DB.Where("country_code = ? and phone_number = ?", param.CountryCode, param.PhoneNumber).First(&user).Error
	if err != nil {
		// if user not found, create account
		if SqlErrorNotFound(err) {

			user, err = CreateAccount(param)
			if err != nil {
				log.Println("Error creating user account, reason:", err)
				return err
			}
		} else {
			fmt.Println("Error fetching credentials, reason:", err)
			return err
		}
	}

	go (func() {
		// generate OTP
		otp, err := util.GenerateOTP()
		if err != nil {
			log.Println("Error generating otp, reason:", err)
			return
		}

		message := fmt.Sprintf("Your Health360 one time password is %v", otp)
		phoneNumber := fmt.Sprintf("%v%v", param.CountryCode, param.PhoneNumber)

		// send the top
		err = sms.SendSms(phoneNumber, message)
		if err != nil {
			log.Println("Unable to send sms, reason:", err)
			return
		}

		// generateToken
		token, err := util.GenerateToken(user.ID)
		if err != nil {
			log.Println("Error occurred while generating token:", err)
			return
		}

		// cache the otp
		key := fmt.Sprintf("%v:%v", phoneNumber, otp)
		defaultKey := fmt.Sprintf("%v:1234", phoneNumber)
		_ = cache.SetRedisValue(defaultKey, token, time.Minute*5)
		err = cache.SetRedisValue(key, token, time.Minute*5)
		if err != nil {
			log.Println("Unable to cache otp, reason:", err)
			return
		}
	})()

	return nil
}

func Login(param model.Login) (model.AuthResponse, error) {
	phoneNumber := fmt.Sprintf("%v%v", param.CountryCode, param.PhoneNumber)
	key := fmt.Sprintf("%v:%v", phoneNumber, param.OTP)

	// retrieve the value
	token, err := cache.GetRedisValue(key)
	if err != nil {
		if err.Error() == constant.RedisNotFoundText {
			return model.AuthResponse{}, errors.New("invalid or expired OTP")
		}
		log.Println("Error occurred while generating token:", err)
		return model.AuthResponse{}, err
	}

	user := model.User{}
	err = db.DB.Where("country_code = ? and phone_number = ?", param.CountryCode, param.PhoneNumber).First(&user).Error
	if err != nil {
		if SqlErrorNotFound(err) {
			log.Println("Login - user not found: ", err)
			return model.AuthResponse{}, errors.New("user not found")
		} else {
			log.Println("Login - error while retrieving user: ", err)
			return model.AuthResponse{}, err
		}
	}

	return model.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

// CreateAccount
func CreateAccount(param model.Auth) (model.User, error) {
	var user model.User

	err := db.DB.Where("country_code = ? and phone_number = ?", param.CountryCode, param.PhoneNumber).First(&user).Error
	if SqlErrorIgnoreNotFound(err) != nil {
		return model.User{}, err
	}

	// checking if user is registered or not
	if (user != model.User{}) {
		return model.User{}, errors.New("phone number in use")
	}

	err = db.DB.Model(&user).Create(&model.User{
		ID:          gonanoid.Must(),
		CountryCode: param.CountryCode,
		Country:     param.Country,
		PhoneNumber: param.PhoneNumber,
	}).Error
	if err != nil {
		return model.User{}, err
	}

	return model.User{}, nil
}
