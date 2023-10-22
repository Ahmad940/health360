package service

import (
	"fmt"
	"github.com/Ahmad940/health360/app/model"
	"github.com/Ahmad940/health360/platform/db"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// getAppointmentsWithRelation returns the provided appointment with its relations(consultant and user) included
func getAppointmentWithRelation(appointment model.Appointment) (model.Appointment, error) {
	// Preload the user record for the appointment.
	err := db.DB.Preload("User").First(&appointment, "user_id = ?", appointment.UserID).Error
	if err != nil {
		return model.Appointment{}, err
	}
	// Preload the user record for the appointment.
	err = db.DB.Preload("Consultant").First(&appointment, "consultant_id = ?", appointment.ConsultantID).Error
	if err != nil {
		return model.Appointment{}, err
	}

	return appointment, nil
}

// getAppointmentsWithRelation returns the provided array of appointments with their relations attached
func getAppointmentsWithRelation(appointments []model.Appointment) ([]model.Appointment, error) {
	var appointmentsWithRelation []model.Appointment

	for _, appointment := range appointments {
		appointmentWithRelation, err := getAppointmentWithRelation(appointment)
		if err != nil {
			return nil, err
		}

		appointmentsWithRelation = append(appointmentsWithRelation, appointmentWithRelation)
	}

	return appointmentsWithRelation, nil
}

// GetUserAppointments returns all appointments for the user with the specified userID
func GetUserAppointments(userID string) ([]model.Appointment, error) {
	var appointments []model.Appointment

	err := db.DB.Find(&appointments, "user_id = ?", userID).Error

	if err != nil {
		return []model.Appointment{}, err
	}

	return getAppointmentsWithRelation(appointments)
}

// GetAppointmentById returns the appointment with the provided ID, if it exists
func GetAppointmentById(appointmentID string, userID string) (model.Appointment, error) {
	appointment := model.Appointment{}

	err := db.DB.Where("id = ? AND user_id = ?", appointmentID, userID).First(&appointment).Error
	if err != nil {
		return model.Appointment{}, err
	}

	return getAppointmentWithRelation(appointment)
}

// CreateAppointment creates an appointment
func CreateAppointment(param model.Appointment) (model.Appointment, error) {
	var appointment model.Appointment
	var user model.User
	var consultant model.Consultant

	// Preload the user record for the appointment.
	err := db.DB.First(&user, "id = ?", param.UserID).Error
	if err != nil {
		if SqlErrorNotFound(err) {
			return model.Appointment{}, fmt.Errorf("user not found")
		}
		return model.Appointment{}, err
	}

	// Preload the consultant record for the appointment.
	err = db.DB.First(&consultant, "id = ?", param.ConsultantID).Error
	if err != nil {
		if SqlErrorNotFound(err) {
			return model.Appointment{}, fmt.Errorf("consultant not found")
		}
		return model.Appointment{}, err
	}

	err = db.DB.Model(&appointment).Create(&model.Appointment{
		ID:           gonanoid.Must(),
		UserID:       param.UserID,
		ConsultantID: param.ConsultantID,
		Time:         param.Time,
	}).Error
	if err != nil {
		return model.Appointment{}, err
	}

	return getAppointmentWithRelation(appointment)
}

// UpdateAppointment updates the details of an appointment
func UpdateAppointment(userId string, appointmentId string) (model.Appointment, error) {
	return model.Appointment{}, nil
}

// DeleteAppointment deletes the appointment with the specified ID
func DeleteAppointment(userId string, appointmentId string) error {
	var appointment model.Appointment

	err := db.DB.First(appointment, "id = ? AND user_id = ?", appointmentId, userId).Error
	if err != nil {
		if SqlErrorNotFound(err) {
			return fmt.Errorf("appointment not found")
		}
		return err
	}

	err = db.DB.Where("id = ? AND user_id = ?", appointmentId, userId).Delete(&appointment).Error
	if err != nil {
		return err
	}

	return nil
}
