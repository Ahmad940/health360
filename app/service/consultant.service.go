package service

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/Ahmad940/health360/app/model"
	"github.com/Ahmad940/health360/pkg/util"
	"github.com/Ahmad940/health360/platform/db"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm/clause"
)

func GetAllCategories() []string {
	return []string{"cardiology", "hair issues", "general checkup", "optician", "dermatologist"}
}

func GetAllConsultants() ([]model.Consultant, error) {
	var consultants []model.Consultant

	err := db.DB.Preload("User").Find(&consultants).Error
	if err != nil {
		return []model.Consultant{}, err
	}

	return consultants, nil
}

func GetConsultantsBySpecialization(specialization string) ([]model.Consultant, error) {
	var consultants []model.Consultant

	err := db.DB.Preload("User").Where("specializations @> ARRAY[?]::varchar[]", []string{specialization}).Find(&consultants).Error

	if err != nil {
		return []model.Consultant{}, err
	}

	return consultants, nil
}

func AddConsultant(param model.AddConsultantParam) (model.Consultant, error) {
	var consultant model.Consultant
	// Preload the user record for the consultant.
	err := db.DB.First(&consultant, "user_id = ?", param.UserID).Error
	if SqlErrorIgnoreNotFound(err) != nil {
		return model.Consultant{}, nil
	}

	if (!reflect.DeepEqual(consultant, model.Consultant{})) {
		return model.Consultant{}, errors.New("consultant already added")
	}

	// check if specialization valid
	err = CheckSpecializationValidity(param.Specializations)
	if err != nil {
		return model.Consultant{}, err
	}

	err = db.DB.Model(&consultant).Create(&model.Consultant{
		ID:              gonanoid.Must(),
		UserID:          param.UserID,
		Services:        param.Services,
		Bio:             param.Bio,
		Specializations: param.Specializations,
	}).Error
	if err != nil {
		return model.Consultant{}, err
	}

	// Preload the user record for the consultant.
	err = db.DB.Preload("User").First(&consultant, "user_id = ?", param.UserID).Error
	if err != nil {
		return model.Consultant{}, nil
	}

	return consultant, nil
}

func UpdateConsultant(param model.UpdateConsultantParam) (model.Consultant, error) {
	var consultant model.Consultant = model.Consultant{
		ID: param.ID,
	}

	// check if specialization valid
	err := CheckSpecializationValidity(param.Specializations)
	if err != nil {
		return model.Consultant{}, err
	}

	// err := db.DB.Model(&user).Clauses(clause.Returning{}).Updates(param).Error
	err = db.DB.Model(&consultant).Clauses(clause.Returning{}).Updates(model.Consultant{
		Bio:             param.Bio,
		Services:        param.Services,
		Specializations: param.Specializations,
	}).Error
	if err != nil {
		return model.Consultant{}, err
	}

	// Preload the user record for the consultant.
	err = db.DB.Preload("User").First(&consultant, "id = ?", param.ID).Error
	if err != nil {
		return model.Consultant{}, nil
	}

	return consultant, nil
}

func RemoveConsultant(id string) error {
	var consultant model.Consultant

	err := db.DB.Where("id = ?", id).Delete(&consultant).Error
	if err != nil {
		return err
	}

	return nil
}

func CheckSpecializationValidity(specializations []string) error {
	for _, specialization := range specializations {
		if isPresent := util.IsItemPresentInArray(specialization, GetAllCategories()); !isPresent {
			return fmt.Errorf("%v is not found in categories", specialization)
		}
	}

	return nil
}
