package main

import (
	"log/slog"
	"net/http"
	"service-tracker-go/internal/models"

	"github.com/gin-gonic/gin"
)

type CustomerData struct {
	Title    string
	Vehicle  models.Vehicle
	Statuses []string
}

type VehicleFormData struct {
	VehicleTypes  []string
	VehicleIssues []string
}

type VehicleRequest struct {
	Name           string   `form:"name" binding:"required,min=2,max=100"`
	Phone          string   `form:"phone" binding:"required,min=10,max=15"`
	Address        string   `form:"address" binding:"required,min=5,max=200"`
	Issues         []string `form:"issues" binding:"required,min=1,dive,Valid_issue_type"`
	VehicleTypes   []string `form:"vehicle_types" binding:"required,min=1,dive,Valid_model_type"`
	IssueReproduce string   `form:"IssueReproduce" binding:"max=500"`
}

func (h *Handler) ServeNewVehicleForm(c *gin.Context) {
	c.HTML(http.StatusOK, "order.tmpl", VehicleFormData{
		VehicleTypes:  models.VehicleTypes,
		VehicleIssues: models.VehicleIssues,
	})
}

func (h *Handler) HandleNewVehiclePost(c *gin.Context) {
	var form VehicleRequest
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicleItems := make([]models.VehicleItem, len(form.Issues))
	for i := range vehicleItems {
		vehicleItems[i] = models.VehicleItem{
			Issue:          form.Issues[i],
			Vehicle:        form.VehicleTypes[i],
			IssueReproduce: form.IssueReproduce,
		}
	}

	vehicle := models.Vehicle{
		CustomerName: form.Name,
		Phone:        form.Phone,
		Address:      form.Address,
		Status:       models.VehicleStatus[0],
		Items:        vehicleItems,
	}

	if err := h.vehicles.CreateVehicle(&vehicle); err != nil {
		slog.Error("Failed to create request", "error", err)
		c.String(http.StatusInternalServerError, "Failed to create request")
	}

	slog.Info("Service created", "vehicleID", vehicle.ID, "customerName", vehicle.CustomerName)

	c.Redirect(http.StatusSeeOther, "/customers/"+vehicle.ID)
}

func (h *Handler) serveCustomer(c *gin.Context) {
	vehicleID := c.Param("id")
	if vehicleID == "" {
		c.String(http.StatusBadRequest, "Invalid vehicle ID")
	}

	vehicle, err := h.vehicles.GetVehicle(vehicleID)
	if err != nil {
		c.String(http.StatusNotFound, "Failed to retrieve vehicle")
		return
	}

	c.HTML(http.StatusOK, "customer.tmpl", CustomerData{
		Vehicle:  *vehicle,
		Title:    "Vehicle - Service Status" + vehicleID,
		Statuses: models.VehicleStatus,
	})

}
