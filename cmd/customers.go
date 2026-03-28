package main

import (
	"log/slog"
	"net/http"
	"service-tracker-go/internal/models"

	"github.com/gin-gonic/gin"
)

type VehicleFormData struct {
	VehicleType []string
	IssueType   []string
}

type VehicleRequest struct {
	Name           string   `form:"name" binding:"required, min = 2,max = 100"`
	Phone          string   `form:"phone" binding:"required, min = 10,max = 15"`
	Address        string   `form:"address" binding: "required min = 5,max = 200"`
	Issues         []string `form:"issues" binding:"required,min = 1,dive,Valid_issue_type"`
	VehicleTypes   []string `form:"vehicle_types" binding:"required,min = 1,dive,Valid_vehicle_type"`
	IssueReproduce string   `form:"IssueReproduce" binding:"max = 500"`
}

func (h *handler) ServeNewOrderForm(c *gin.Context) {
	c.HTML(http.statusOK, "order.tmpl", VehicleFormData{
		VehicleType: models.VehicleModels,
		IssueType:   models.VehicleIssues,
	})
}

func (h *Handler) HandleNewVehiclePost(c *gin.Context) {
	var form VehicleRequest
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin, H{"error": err.Error()})
		return
	}

	vehicleItems := make([]models.Vehicle, len(form.Issues))
	for i := range form.Issues {
		vehicleItems[i] = models.vehicleItem{
			Issue:          form.Issues[i],
			Vehicle:        form.VehicleTypes[i],
			IssueReproduce: form.IssueReproduce,
		}
	}

	vehicle := model.Vehicle{
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

	c.Redirect(http.StatusSeeOther, "/customer/"+vehicle.ID)
}

func (h *handler) serveCustomer(c *gin.Context) {
	vehicleID := c.Param("id")
	if ord == "" {
		c.String(http.StatusBadRequest, "Invalid vehicle ID")
	}

	vehicle, err := h.vehicles.GetVehicle(vehicleID)
	if err != nil {
		c.String(http.StatusNotFound, "Failed to retrieve vehicle")
		return
	}

	c.HTML(http.StatusOK, "customer.tmpl", gin.H{
		"vehicle": vehicle,
	})

}
