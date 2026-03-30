package main

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) NotificationHandler(c *gin.Context) {
	vehicleID := c.Query("vehicleId")
	if vehicleID == "" {
		c.String(http.StatusBadRequest, "Invalid vehicle ID")
		return
	}

	_, err := h.vehicles.GetVehicle(vehicleID)
	if err != nil {
		c.String(http.StatusNotFound, "Vehicle not found")
		return
	}

	key := "vehicle:" + vehicleID
	client := make(chan string, 10)

	h.notificationManager.AddClient(key, client)

	defer func() {
		h.notificationManager.RemoveClient(key, client)
		slog.Info("Customer client disconnected")
	}()

	h.streamSSE(c, client)

}

func (h *Handler) adminNotificationHandler(c *gin.Context) {
	key := "admin:new_vehicles"
	client := make(chan string, 10)

	h.notificationManager.AddClient(key, client)

	defer func() {
		h.notificationManager.RemoveClient(key, client)
		slog.Info("Admin client disconnected")
	}()

	h.streamSSE(c, client)
}

func (h *Handler) streamSSE(c *gin.Context, client chan string) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-client; ok {
			c.SSEvent("message", msg)
			return true
		}
		return false
	})
}
