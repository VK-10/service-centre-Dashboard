package main

import "service-tracker-go/internal/models"

type Handler struct {
	vehicles            *models.VehicleModel
	users               *models.UserModel
	notificationManager *NotificationManager
}

func NewHandler(dbModel *models.DBModel) *Handler {
	return &Handler{
		vehicles:            &dbModel.Vehicle,
		users:               &dbModel.User,
		notificationManager: NewNotificationmanager(),
	}
}
