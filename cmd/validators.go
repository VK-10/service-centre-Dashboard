package main

import (
	"service-tracker-go/internal/models"
	"slices"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.registerValidation("Valid_model_type", createVehicleValidator(models.VehicleModels))
		v.registerValidation("Valid_issue_type", createVehicleIssueValidator(models.VehicleIssues))
	}
}

func CreateVehicleValidator(validModels []string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		return slices.Contains(validModels, fl.Field().String())
	}
}
