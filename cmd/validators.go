package main

import (
	"service-tracker-go/internal/models"
	"slices"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("Valid_model_type", CreateSliceValidator(models.VehicleTypes))
		v.RegisterValidation("Valid_issue_type", CreateSliceValidator(models.VehicleIssues))
	}
}

func CreateSliceValidator(allowedValues []string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		return slices.Contains(allowedValues, fl.Field().String())
	}
}
