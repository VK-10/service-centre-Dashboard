package main

import "service-tracker-go/internal/models"

type Handler struct {
	vehicles *models.VehicleModel
}

func NewHandler(dbModel *models.DBModel) *Handler {
	return &Handler{
		vehicles: &dbModel.Vehicle,
	}
}
