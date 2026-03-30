package main

import (
	"net/http"
	"service-tracker-go/internal/models"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Error string
}

type AdminDashboardData struct {
	Vehicles []models.Vehicle
	Statuses []string
	Username string
}

func (h *Handler) HandleLoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", LoginData{})
}

func (h *Handler) HandleLoginPost(c *gin.Context) {
	var form struct {
		Username string `form:"username" binding:"required,min=3,max=50"`
		Password string `form:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusOK, "login.tmpl", LoginData{Error: "Invalid input: " + err.Error()})
		return
	}

	user, err := h.users.AuthenticateUser(form.Username, form.Password)
	if err != nil {
		c.HTML(http.StatusOK, "login.tmpl", LoginData{
			Error: "Invalid credentials",
		})

		return
	}

	SetSessionValue(c, "userID", user.ID)
	SetSessionValue(c, "username", user.Username)

	c.Redirect(http.StatusSeeOther, "/admin")

}

func (h *Handler) HandleLogout(c *gin.Context) {
	if err := ClearSession(c); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, "/login")
}

func (h *Handler) ServeAdminDashboard(c *gin.Context) {
	vehicles, err := h.vehicles.GetAllVehicles()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	username := GetSessionString(c, "username")

	c.HTML(http.StatusOK, "admin.tmpl", AdminDashboardData{
		Vehicles: vehicles,
		Username: username,
		Statuses: models.VehicleStatus,
	})

}

func (h *Handler) HandleOrderPut(c *gin.Context) {
	vehicleId := c.Param("id")
	newStatus := c.Param("status")

	if err := h.vehicles.UpdateVehicleStatus(vehicleId, newStatus); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	h.notificationManager.Notify("vehicle:"+vehicleId, "status_update")

	c.Redirect(http.StatusSeeOther, "/admin")
}

func (h *Handler) HandleVehicleDelete(c *gin.Context) {
	vehicleId := c.Param("id")

	if err := h.vehicles.DeleteVehicle(vehicleId); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, "/admin")
}
