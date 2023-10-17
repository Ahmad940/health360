package service

import (
	"github.com/Ahmad940/health360/app/model"
	"github.com/Ahmad940/health360/platform/db"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm/clause"
)

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

	err := db.DB.Preload("User").Where("specializations @> ARRAY[?]::varchar[]", []string{"cardiology"}).Find(&consultants).Error

	if err != nil {
		return []model.Consultant{}, err
	}

	return consultants, nil
}

func AddConsultant(param model.ModifyConsultantParam) (model.Consultant, error) {
	var consultant model.Consultant
	err := db.DB.Preload("User").Model(&consultant).Clauses(clause.Returning{}).Create(&model.Consultant{
		ID:     gonanoid.Must(),
		UserID: param.UserID,
	}).Error
	if err != nil {
		return model.Consultant{}, err
	}

	return consultant, nil
}

func UpdateConsultant(param model.ModifyConsultantParam) (model.Consultant, error) {
	var consultant model.Consultant = model.Consultant{
		UserID:          param.ID,
		Services:        param.Services,
		Specializations: param.Specializations,
	}
	// err := db.DB.Model(&user).Clauses(clause.Returning{}).Updates(param).Error
	err := db.DB.Model(&consultant).Preload("User").Clauses(clause.Returning{}).Updates(param).Error
	if err != nil {
		return model.Consultant{}, err
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
